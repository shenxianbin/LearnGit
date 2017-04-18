package gamemap

/*地图验证思路
1.统计各格数据（计数）
2.金币点验证（障碍物）
*/

import (
	. "Gameserver/cache"
	"Gameserver/global"
	. "Gameserver/logic"
	"common"
	"common/protocol"
	"common/scheme"
	"errors"
	"fmt"
	. "galaxy"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	cache_map_key_t = "Role:%v:Map"
)

type MapSys struct {
	MapCache
	owner     IRole
	cache_key string
	heros     map[int64]bool
}

func (this *MapSys) Init(owner IRole) {
	this.owner = owner
	this.MapCache.Maps = make(map[int32]*MapGrid)
	this.MapCache.MapPointList = make(map[int32]int32)
	this.MapCache.MapPointActive = make([]int32, 0)
	this.cache_key = fmt.Sprintf(cache_map_key_t, this.owner.GetUid())
	this.heros = make(map[int64]bool)
}

func (this *MapSys) Load() error {
	resp, err := GxService().Redis().Cmd("GET", this.cache_key)
	if err != nil {
		return err
	}

	if buf, _ := resp.Bytes(); buf != nil {
		err := proto.Unmarshal(buf, &this.MapCache)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *MapSys) Check() {
	for _, v := range this.MapCache.Maps {
		switch v.GetObjType() {
		case common.MAP_OBJ_MAGIC_HERO:
			this.heros[v.GetId()] = true
		}
	}
}

func (this *MapSys) Save() {
	buf, err := proto.Marshal(&this.MapCache)
	if err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SET", this.cache_key, buf); err != nil {
		LogFatal(err)
		return
	}
}

func (this *MapSys) parseXYToIndex(x, y int32) int32 {
	return x<<16 | y
}

func (this *MapSys) parseIndexToXY(index int32) (x, y int32) {
	x = index >> 16
	y = int32(int16(index))
	return
}

func (this *MapSys) FillMapInfo() *protocol.MapInfo {
	msg := new(protocol.MapInfo)
	msg.MapInfos = make([]*protocol.MapGridInfo, len(this.MapCache.Maps))
	index := 0
	for i, v := range this.MapCache.Maps {
		map_point := new(protocol.MapGridInfo)
		x, y := this.parseIndexToXY(i)
		map_point.SetBaseType(v.GetBaseType())
		map_point.SetObjType(v.GetObjType())
		map_point.SetId(v.GetId())
		map_point.SetX(x)
		map_point.SetY(y)
		msg.MapInfos[index] = map_point
		index++
	}

	msg.MapPointList = make([]int32, len(this.MapCache.GetMapPointList()))
	map_point_index := 0
	for k, _ := range this.MapCache.GetMapPointList() {
		msg.MapPointList[map_point_index] = k
		map_point_index++
	}
	msg.MapPointActive = this.MapCache.MapPointActive

	return msg
}

func (this *MapSys) MapSetPoint(x int32, y int32, base_type int32, obj_type int32, id int64) {
	if x > this.MapCache.GetXSize() || y > this.MapCache.GetYSize() {
		return
	}

	index := this.parseXYToIndex(x, y)
	if v, has := this.MapCache.Maps[index]; has {
		v.SetBaseType(base_type)
		v.SetObjType(obj_type)
		v.SetId(id)
	} else {
		point := &MapGrid{
			BaseType: proto.Int32(base_type),
			ObjType:  proto.Int32(obj_type),
			Id:       proto.Int64(id),
		}
		this.MapCache.Maps[index] = point
	}
}

func (this *MapSys) map_initial(index_x int32, index_y int32, point_id int32, base_type int32, obj_type int32, scheme_id int32) error {
	switch obj_type {
	case common.MAP_OBJ_NONE, common.MAP_OBJ_ENTRANCE, common.MAP_OBJ_MAGIC_KING:
		this.MapSetPoint(index_x, index_y, base_type, obj_type, 0)
	case common.MAP_OBJ_MAGIC_SOLDIER:
		if this.owner.SoldierCreateFree(scheme_id, 1) {
			soldier := this.owner.SoldierDispatch(scheme_id)
			if soldier == nil {
				return errors.New("SoldierDispatch Error")
			}
			id := int64(soldier.GetAutoId())
			this.MapSetPoint(index_x, index_y, base_type, obj_type, id)
		}
	case common.MAP_OBJ_MAGIC_HERO:
		id := this.owner.HeroObtain(scheme_id/10, 1, scheme_id%10, false)
		if id == common.UID_FAILED {
			return errors.New("HeroObtainByMap Error")
		}
		this.MapSetPoint(index_x, index_y, base_type, obj_type, id)
		this.heros[id] = true
	case common.MAP_OBJ_MAGIC_BUILDING:
		if _, has := this.MapCache.MapPointList[point_id]; has {
			id := this.owner.BuildingObtain(scheme_id, 1, false)
			if id == common.UID_FAILED {
				return errors.New("BuildingObtainByMap Error")
			}
			this.MapSetPoint(index_x, index_y, base_type, obj_type, id)
		} else {
			this.MapSetPoint(index_x, index_y, base_type, obj_type, 0)
		}
	case common.MAP_OBJ_DECORATION:
		if ret := this.owner.DecorationObtainByMap(scheme_id, index_x, index_y); ret != common.RetCode_Success {
			return errors.New("DecorationObtainByMap Error")
		}
	case common.MAP_OBJ_OBSTACLE:
		if ret := this.MapFreshObstacle(scheme_id, index_x, index_y); ret != common.RetCode_Success {
			return errors.New("ObstacleObtainByMap Error")
		}
	default:
		return errors.New("Unknow MapObjType")
	}
	return nil
}

func (this *MapSys) MapFreshObstacle(scheme_id int32, pos_x int32, pos_y int32) common.RetCode {
	obstacle, has := scheme.Obstaclemap[scheme_id]
	if !has {
		return common.RetCode_SchemeData_Error
	}

	size := strings.Split(obstacle.Size, ";")
	size_x, _ := strconv.Atoi(size[0])
	size_y, _ := strconv.Atoi(size[1])
	if size_x == 0 || size_y == 0 {
		return common.RetCode_SchemeData_Error
	}

	for i := pos_x; i < pos_x+int32(size_x); i++ {
		for j := pos_y; j < pos_y+int32(size_y); j++ {
			this.MapSetPoint(i, j, common.MAP_NOWAY, common.MAP_OBJ_OBSTACLE, int64(obstacle.Id))
		}
	}

	return common.RetCode_Success
}

func (this *MapSys) MapRemoveObstacle(scheme_id int32, pos_x int32, pos_y int32) common.RetCode {
	obstacle, has := scheme.Obstaclemap[scheme_id]
	if !has {
		return common.RetCode_SchemeData_Error
	}

	size := strings.Split(obstacle.Size, ";")
	size_x, _ := strconv.Atoi(size[0])
	size_y, _ := strconv.Atoi(size[1])
	if size_x == 0 || size_y == 0 {
		return common.RetCode_SchemeData_Error
	}

	for i := pos_x; i < pos_x+int32(size_x); i++ {
		for j := pos_y; j < pos_y+int32(size_y); j++ {
			index := this.parseXYToIndex(pos_x, pos_y)
			v, has := this.MapCache.Maps[index]
			if !has {
				return common.RetCode_SchemeData_Error
			}
			if v.GetObjType() != common.MAP_OBJ_OBSTACLE || v.GetId() != int64(scheme_id) {
				return common.RetCode_Fail
			}
		}
	}

	switch obstacle.RType {
	case common.RTYPE_GOLD:
		need_gold := obstacle.Num
		if !this.owner.IsEnoughGold(need_gold) {
			return common.RetCode_Unable
		}
		this.owner.CostGold(need_gold, true, true)
	}

	for i := pos_x; i < pos_x+int32(size_x); i++ {
		for j := pos_y; j < pos_y+int32(size_y); j++ {
			this.owner.MapSetPoint(i, j, common.MAP_NOWAY, common.MAP_OBJ_NONE, 0)
		}
	}

	this.Save()

	//添加成就
	this.owner.AchievementAddNum(17, 1, false)

	return common.RetCode_Success
}

func (this *MapSys) MapInitial(x_size int32, y_size int32) error {
	scheme_map := scheme.GetMap(x_size, y_size)
	if scheme_map == nil {
		return errors.New("MapInitial scheme error")
	}
	this.MapCache.SetXSize(x_size)
	this.MapCache.SetYSize(y_size)

	for index_x, v_x := range scheme_map {
		for index_y, v_y := range v_x {
			point_id := v_y.PointId
			base_type := v_y.IsWay
			obj_type := v_y.ObjType
			scheme_id := v_y.SchemeId

			//用于过滤占多格的障碍物和装饰物
			if _, has := this.MapCache.Maps[this.parseXYToIndex(int32(index_x), int32(index_y))]; has {
				continue
			}
			if err := this.map_initial(int32(index_x), int32(index_y), point_id, base_type, obj_type, scheme_id); err != nil {
				return err
			}
		}
	}
	LogDebug("Role[", this.owner.GetUid(), "] MapInitial Success Size", len(this.MapCache.Maps))

	this.Save()
	return nil
}

func (this *MapSys) MapExpand(x_size int32, y_size int32) error {
	if this.MapCache.GetXSize() > x_size || this.MapCache.GetYSize() > y_size {
		return errors.New("ExpandMap size error")
	}

	if this.MapCache.GetXSize() == x_size && this.MapCache.GetYSize() == y_size {
		return nil
	}

	scheme_map := scheme.GetMap(x_size, y_size)
	if scheme_map == nil {
		return errors.New("ExpandMap scheme error")
	}

	old_map := make(map[int32]*MapGrid)
	for index, v := range this.MapCache.Maps {
		old_map[index] = v
	}

	fix_x := (x_size - this.MapCache.GetXSize()) / 2
	fix_y := y_size - this.MapCache.GetYSize()
	this.MapCache.Maps = make(map[int32]*MapGrid)
	for index, v := range old_map {
		x, y := this.parseIndexToXY(index)
		new_index := this.parseXYToIndex(x+fix_x, y+fix_y)
		this.MapCache.Maps[new_index] = v
	}

	this.MapCache.SetXSize(x_size)
	this.MapCache.SetYSize(y_size)
	for index_x, v_x := range scheme_map {
		for index_y, v_y := range v_x {
			point_id := v_y.PointId
			base_type := v_y.IsWay
			obj_type := v_y.ObjType
			scheme_id := v_y.SchemeId

			//用于过滤占多格的障碍物和装饰物
			if _, has := this.MapCache.Maps[this.parseXYToIndex(int32(index_x), int32(index_y))]; has {
				continue
			}
			if err := this.map_initial(int32(index_x), int32(index_y), point_id, base_type, obj_type, scheme_id); err != nil {
				return err
			}
		}
	}

	this.Save()
	this.send_update_info()

	return nil
}

func (this *MapSys) MapFindHero(uid int64) bool {
	_, has := this.heros[uid]
	return has
}

//验证地图合法性
func (this *MapSys) authValid(x int32, y int32, base_type int32, obj_type int32, id int64) bool {
	index := this.parseXYToIndex(x, y)
	v, has := this.MapCache.Maps[index]
	if !has {
		return false
	}

	switch obj_type {
	case common.MAP_OBJ_MAGIC_SOLDIER:
		//魔物时 需转换auto_id | scheme_id
		scheme_id := int32(id)
		if this.owner.SoldierNum(scheme_id) <= 0 {
			return false
		}
	case common.MAP_OBJ_MAGIC_HERO:
		if !this.owner.HeroFind(id) {
			return false
		}
	case common.MAP_OBJ_MAGIC_BUILDING:
		if !this.owner.BuildingFind(id) {
			return false
		}
	}

	if v.GetObjType() == common.MAP_OBJ_OBSTACLE && obj_type != common.MAP_OBJ_OBSTACLE {
		return false
	}

	if v.GetObjType() == common.MAP_OBJ_OBSTACLE && v.GetId() != id {
		return false
	}

	return true
}

//刷新地图
func (this *MapSys) MapReFresh(map_info []*protocol.MapGridInfo, map_point_active []int32) common.RetCode {
	if map_info == nil {
		LogError("MapReFresh role_uid : ", this.owner.GetUid(), " mapinfo nil")
		return common.RetCode_Fail
	}

	if int32(len(map_info)) != this.GetXSize()*this.GetYSize() {
		LogError("MapReFresh role_uid : ", this.owner.GetUid(), " mapsize error")
		return common.RetCode_Fail
	}

	if int32(len(map_point_active)) > this.owner.GetFortressLimit() {
		LogError("MapReFresh role_uid : ", this.owner.GetUid(), " fortress limit error: ", len(map_point_active))
		return common.RetCode_Fail
	}

	var way int32
	var entrance bool
	decoration := make(map[int32]int32)
	soldier := make(map[int32]int32)
	hero := make(map[int64]bool)
	building := make(map[int64]bool)
	var king bool
	var population int32

	//统计筛选
	for _, v := range map_info {
		this.authValid(v.GetX(), v.GetY(), v.GetBaseType(), v.GetObjType(), v.GetId())
		if v.GetBaseType() == common.MAP_WAY {
			way++
		}

		switch v.GetObjType() {
		case common.MAP_OBJ_ENTRANCE:
			if entrance {
				LogError("MapReFresh role_uid : ", this.owner.GetUid(), " entrance repeat")
				return common.RetCode_Fail
			}
			entrance = true
		case common.MAP_OBJ_DECORATION:
			d, has := decoration[int32(v.GetId())]
			if has {
				decoration[int32(v.GetId())] = d + 1
			} else {
				decoration[int32(v.GetId())] = 1
			}
		case common.MAP_OBJ_MAGIC_SOLDIER:
			//魔物时 需转换auto_id | scheme_id
			id := int32(v.GetId())
			if _, has := soldier[id]; has {
				soldier[id] += 1
			} else {
				soldier[id] = 1
			}

		case common.MAP_OBJ_MAGIC_HERO:
			_, has := hero[v.GetId()]
			if has {
				LogError("MapReFresh role_uid : ", this.owner.GetUid(), " hero (", v.GetId(), ") error")
				return common.RetCode_Fail
			}
			hero[v.GetId()] = true
		case common.MAP_OBJ_MAGIC_BUILDING:
			if v.GetId() != 0 {
				_, has := building[v.GetId()]
				if has {
					LogError("MapReFresh role_uid : ", this.owner.GetUid(), " building (", v.GetId(), ") error")
					return common.RetCode_Fail
				}
				building[v.GetId()] = true
			}
		case common.MAP_OBJ_MAGIC_KING:
			if king {
				LogError("MapReFresh role_uid : ", this.owner.GetUid(), " king nil")
				return common.RetCode_Fail
			}
			king = true
		}
	}

	//可挖掘数
	if way > this.owner.GetDigLimit() {
		LogError("MapReFresh role_uid : ", this.owner.GetUid(), " diglimit error current : ", way, " limit : ", this.owner.GetDigLimit())
		return common.RetCode_Fail
	}

	//装饰物
	decoration_size := make(map[int32]int32)
	for k, _ := range decoration {
		_, has := decoration_size[k]
		if has {
			continue
		}
		d_sc, has := scheme.Decorationmap[k]
		if !has {
			LogError("MapReFresh role_uid : ", this.owner.GetUid(), " decoration (", k, ") error")
			return common.RetCode_Fail
		}
		size_str := strings.Split(d_sc.Size, ";")
		if len(size_str) != 2 {
			LogError("MapReFresh role_uid : ", this.owner.GetUid(), " decoration scheme error")
			return common.RetCode_SchemeData_Error
		}
		size_x, _ := strconv.Atoi(size_str[0])
		size_y, _ := strconv.Atoi(size_str[1])
		decoration_size[k] = int32(size_x * size_y)
	}

	for k, v := range decoration {
		size, _ := decoration_size[k]
		num := v / size
		if num > this.owner.DecorationSize(k) {
			LogError("MapReFresh role_uid : ", this.owner.GetUid(), " decoration (", k, ") num error ! mapinfo : ", num, " owner : ", this.owner.DecorationSize(k))
			return common.RetCode_Fail
		}
	}

	//魔使
	if int32(len(hero)) > this.owner.HeroSize() {
		LogError("MapReFresh role_uid : ", this.owner.GetUid(), " hero size error cur(", len(hero), ") need(", this.owner.HeroSize(), ")")
		return common.RetCode_Fail
	}

	//完成成就
	this.owner.AchievementAddNum(3, int32(len(hero)), true)

	for k, _ := range hero {
		population += this.owner.HeroPopulation(k)
	}

	//魔物
	for id, num := range soldier {
		if this.owner.SoldierNum(id) < num {
			LogError("MapReFresh role_uid : ", this.owner.GetUid(), " soldier(", id, ") all_num(", this.owner.SoldierNum(id), ") map_num(", num, ") size error")
			return common.RetCode_Fail
		} else {
			population += scheme.Soldiermap[id].Population * num
		}
	}

	//人口
	if population > this.owner.GetPopLimit() {
		return common.RetCode_Fail
	}

	//建筑
	if int32(len(building)) > this.owner.BuildingSize() {
		LogError("MapReFresh role_uid : ", this.owner.GetUid(), " building size error")
		return common.RetCode_Fail
	}

	//保存地图
	this.owner.SoldierWithdraw()
	this.heros = make(map[int64]bool)
	for _, v := range map_info {
		switch v.GetObjType() {
		case common.MAP_OBJ_MAGIC_SOLDIER:
			//魔物时 需转换auto_id | scheme_id
			auto_id := v.GetId() >> 32
			scheme_id := int32(v.GetId())
			LogDebug("Soldier ID:", v.GetId(), " auto_id:", auto_id, " scheme_id:", scheme_id)
			if this.owner.SoldierDispatchByAutoId(scheme_id, auto_id) == nil {
				LogError("SaveMap Dispatch SoldierByAutoId error ! auto_id[", auto_id, "] scheme_id[", scheme_id, "]")
			}
			this.MapSetPoint(v.GetX(), v.GetY(), v.GetBaseType(), v.GetObjType(), auto_id)
		case common.MAP_OBJ_MAGIC_BUILDING:
			this.MapSetPoint(v.GetX(), v.GetY(), v.GetBaseType(), v.GetObjType(), v.GetId())
		case common.MAP_OBJ_MAGIC_HERO:
			this.heros[v.GetId()] = true
			this.MapSetPoint(v.GetX(), v.GetY(), v.GetBaseType(), v.GetObjType(), v.GetId())
		default:
			this.MapSetPoint(v.GetX(), v.GetY(), v.GetBaseType(), v.GetObjType(), v.GetId())
		}
	}
	this.MapCache.SetMapPointActive(map_point_active)
	this.Save()

	return common.RetCode_Success
}

//func (this *MapSys) MapRemoveDeathSoldier(auto_id int64) {
//	for k, v := range this.Maps {
//		if v.GetObjType() == common.MAP_OBJ_MAGIC_SOLDIER && v.GetId() == auto_id {
//			this.Maps[k].SetBaseType(common.MAP_WAY)
//			this.Maps[k].SetObjType(common.MALL_TYPE_NONE)
//			this.Maps[k].SetId(0)
//		}
//	}
//	this.Save()
//}

func (this *MapSys) MapInfo(role_uid int64) (common.RetCode, *protocol.MapFightInfo) {
	target_role := this.owner.OfflineRoleFight(role_uid)
	if target_role != nil {
		return common.RetCode_Success, target_role.FillMapFightInfo()
	}
	return common.RetCode_Fail, nil
}

func (this *MapSys) MapInitPointList() {
	pointMap := make(map[int32]int32)
	for _, scheme := range scheme.PointMap {
		if !scheme.IsLock {
			pointMap[scheme.PointId] = scheme.PointId
		}
	}
	this.MapCache.SetMapPointList(pointMap)
}

func (this *MapSys) MapUnLockPoint(pointId int32) common.RetCode {
	LogDebug(pointId)

	if _, has := this.MapCache.MapPointList[pointId]; has {
		return common.RetCode_Fail
	}

	data, has := scheme.PointMap[pointId]
	if !has {
		LogError("UnLockMapPoint Error Id : ", pointId)
		return common.RetCode_SchemeData_Error
	}

	if data.NeedKingLv > this.owner.GetKingLv() {
		return common.RetCode_Unable
	}

	switch data.RType {
	case common.RTYPE_BLOOD:
		if !this.owner.IsEnoughBlood(data.Value) {
			return common.RetCode_Unable
		}
	case common.RTYPE_SOUL:
		if !this.owner.IsEnoughSoul(data.Value) {
			return common.RetCode_Unable
		}
	case common.RTYPE_GOLD:
		if !this.owner.IsEnoughGold(data.Value) {
			return common.RetCode_Unable
		}
	default:
		LogError("UnLockMapPoint RTypeError Id : ", pointId, " RType : ", data.RType)
		return common.RetCode_SchemeData_Error
	}

	this.MapCache.MapPointList[pointId] = pointId
	switch data.RType {
	case common.RTYPE_BLOOD:
		this.owner.CostBlood(data.Value, true, false)
	case common.RTYPE_SOUL:
		this.owner.CostSoul(data.Value, true, false)
	case common.RTYPE_GOLD:
		this.owner.CostGold(data.Value, true, false)
	}

	LogDebug(scheme.PointMapBuilding)
	if build_list, has := scheme.PointMapBuilding[pointId]; has {
		LogDebug(build_list)
		for _, v := range build_list {
			x, y, err := scheme.ParseXY(v.X, v.Y, this.MapCache.GetXSize(), this.MapCache.GetYSize())
			if err != nil {
				LogError("MapUnLock BuildingObtain ParseXY Error SchemeId : ", v.BuildingId, " X : ", v.X, " Y : ", v.Y)
				continue
			}
			if buildingId := this.owner.BuildingObtain(v.BuildingId, 1, true); buildingId != common.UID_FAILED {
				this.MapSetPoint(x, y, v.IsWay, common.MAP_OBJ_MAGIC_BUILDING, buildingId)
				this.send_update_point_info(x, y, v.IsWay, common.MAP_OBJ_MAGIC_BUILDING, buildingId)
				LogDebug("MapUnLock BuildingObtain Success SchemeId : ", v.BuildingId, " X : ", x, " Y : ", y)
			} else {
				LogError("MapUnLock BuildingObtain Error SchemeId : ", v.BuildingId, " X : ", x, " Y : ", y)
			}
		}
	}

	this.Save()
	return common.RetCode_Success
}

func (this *MapSys) send_update_info() {
	msg := new(protocol.MsgMapUpdateNotify)
	msg.MapInfos = make([]*protocol.MapGridInfo, len(this.MapCache.Maps))
	index := 0
	for i, v := range this.MapCache.Maps {
		map_point := new(protocol.MapGridInfo)
		x, y := this.parseIndexToXY(i)
		map_point.SetBaseType(v.GetBaseType())
		map_point.SetObjType(v.GetObjType())
		map_point.SetId(v.GetId())
		map_point.SetX(x)
		map_point.SetY(y)
		msg.MapInfos[index] = map_point
		index++
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	global.SendMsg(int32(protocol.MsgCode_MapUpdateNotify), this.owner.GetSid(), buf)
}

func (this *MapSys) send_update_point_info(x int32, y int32, base_type int32, obj_type int32, id int64) {
	msg := new(protocol.MsgMapUpdatePointNotify)
	msg.MapInfo = new(protocol.MapGridInfo)
	msg.MapInfo.SetX(x)
	msg.MapInfo.SetY(y)
	msg.MapInfo.SetBaseType(base_type)
	msg.MapInfo.SetObjType(obj_type)
	msg.MapInfo.SetId(id)

	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	LogDebug(msg)
	global.SendMsg(int32(protocol.MsgCode_MapUpdatePointNotify), this.owner.GetSid(), buf)
}
