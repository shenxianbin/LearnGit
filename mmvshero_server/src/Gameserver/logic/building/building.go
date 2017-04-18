package building

import (
	"Gameserver/global"
	. "Gameserver/logic"
	"common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"errors"
	"fmt"
	. "galaxy"
	"math/rand"

	"github.com/golang/protobuf/proto"
)

const (
	cache_building_autokey_t        = "Role:%v:BuildingAutoKey"
	cache_buildinglist_key_template = "Role:%v:Building"
	cache_buildingobj_key_template  = "Role:%v:Building:%v"
)

func GenBuildingListKey(role_uid int64) string {
	return fmt.Sprintf(cache_buildinglist_key_template, role_uid)
}

func genBuildingCacheKey(role_uid int64, building_uid int64) string {
	return fmt.Sprintf(cache_buildingobj_key_template, role_uid, building_uid)
}

func genBuildingAutoKey(role_uid int64) string {
	return fmt.Sprintf(cache_building_autokey_t, role_uid)
}

type Building struct {
	BuildingCache
	scheme_base_data *scheme.Building
	scheme_lvup_data *scheme.BuildingLvUp
}

func NewBuilding(scheme_id int32, lv int32, role_uid int64) (*Building, error) {
	scheme_base_data, has := scheme.Buildingmap[scheme_id]
	if !has {
		return nil, errors.New("NewBuilding Base Scheme Error")
	}

	scheme_lvup_data := scheme.BuildingLvUpGet(scheme_id, lv)
	if scheme_lvup_data == nil {
		return nil, errors.New("NewBuilding Lvup Scheme Error")
	}

	obj := new(Building)
	resp, err := GxService().Redis().Cmd("INCR", genBuildingAutoKey(role_uid))
	if err != nil {
		return nil, err
	}

	uid, _ := resp.Int64()
	obj.BuildingCache.SetUid(uid)
	obj.BuildingCache.SetSchemeId(scheme_id)
	obj.BuildingCache.SetLv(lv)

	obj.scheme_base_data = scheme_base_data
	obj.scheme_lvup_data = scheme_lvup_data

	return obj, nil
}

func LoadBuilding(buf []byte) (*Building, error) {
	obj := new(Building)
	err := proto.Unmarshal(buf, &obj.BuildingCache)
	if err != nil {
		return nil, err
	}

	scheme_base_data, has := scheme.Buildingmap[obj.GetSchemeId()]
	if !has {
		return nil, errors.New("NewBuilding Base Scheme Error")
	}

	scheme_lvup_data := scheme.BuildingLvUpGet(obj.GetSchemeId(), obj.GetLv())
	if scheme_lvup_data == nil {
		return nil, errors.New("Load Building Scheme Error")
	}

	obj.scheme_base_data = scheme_base_data
	obj.scheme_lvup_data = scheme_lvup_data
	return obj, nil
}

func (this *Building) GetLvUpSoul() int32 {
	return this.scheme_lvup_data.LvUpSoul
}

func (this *Building) GetLvUpRoleLv() int32 {
	return this.scheme_lvup_data.LvUpRoleLv
}

func (this *Building) FillBuildInfo() *protocol.BuildingInfo {
	msg := new(protocol.BuildingInfo)
	msg.SetUid(this.BuildingCache.GetUid())
	msg.SetSchemeId(this.BuildingCache.GetSchemeId())
	msg.SetLv(this.BuildingCache.GetLv())
	return msg
}

type BuildingSys struct {
	owner          IRole
	building_list  map[int64]*Building
	cache_list_key string
}

func (this *BuildingSys) Init(owner IRole) {
	this.owner = owner
	this.building_list = make(map[int64]*Building)
	this.cache_list_key = fmt.Sprintf(cache_buildinglist_key_template, this.owner.GetUid())
}

func (this *BuildingSys) Load() error {
	resp, err := GxService().Redis().Cmd("SMEMBERS", this.cache_list_key)
	if err != nil {
		return err
	}

	cacheKeys, err := resp.List()
	if err != nil {
		return err
	}

	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", this.cache_list_key, key)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			building, err := LoadBuilding(buf)
			if err != nil {
				continue
			}
			this.building_list[building.GetUid()] = building
		}
	}

	return nil
}

func (this *BuildingSys) Save(building *Building) {
	buf, err := proto.Marshal(&building.BuildingCache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := genBuildingCacheKey(this.owner.GetUid(), building.BuildingCache.GetUid())
	if _, err := GxService().Redis().Cmd("SET", key, buf); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SADD", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
}

func (this *BuildingSys) FillBuildingListInfo() *protocol.BuildingListInfo {
	msg := new(protocol.BuildingListInfo)
	msg.BuildList = make([]*protocol.BuildingInfo, len(this.building_list))
	index := 0
	for _, v := range this.building_list {
		msg.BuildList[index] = v.FillBuildInfo()
		index++
	}
	return msg
}

func (this *BuildingSys) BuildingSize() int32 {
	return int32(len(this.building_list))
}

func (this *BuildingSys) BuildingFind(building_uid int64) bool {
	_, has := this.building_list[building_uid]
	return has
}

func (this *BuildingSys) BuildingRandomType(ban_id []int32) int32 {
	ban := make(map[int32]bool)
	if ban_id != nil {
		for _, v := range ban_id {
			ban[v] = true
		}
	}

	list := make([]int32, 0)
	for _, v := range this.building_list {
		if _, has := ban[v.GetSchemeId()]; !has {
			list = append(list, v.GetSchemeId())
		}
	}

	if len(list) > 0 {
		r := rand.Intn(len(list))
		return list[r]
	}

	return 0
}

func (this *BuildingSys) BuildingObtain(scheme_id int32, lv int32, is_notify bool) (int64, common.RetCode) {
	if _, has := scheme.Buildingmap[scheme_id]; !has {
		LogError("BuildingObtain BuildingMap NULL Scheme_id : ", scheme_id)
		return 0, common.RetCode_SchemeData_Error
	}

	if lvUp := scheme.BuildingLvUpGet(scheme_id, lv); lvUp == nil {
		LogError("BuildingObtain BuildingLvUp NULL Scheme_id : ", scheme_id, " Lv : ", lv)
		return 0, common.RetCode_SchemeData_Error
	}

	building, err := NewBuilding(scheme_id, lv, this.owner.GetUid())
	if err != nil {
		LogFatal(err)
		return 0, common.RetCode_SchemeData_Error
	}

	this.building_list[building.GetUid()] = building
	this.Save(building)

	if is_notify {
		this.send_notify(building)
	}

	this.static_building(building)

	return building.GetUid(), common.RetCode_Success
}

func (this *BuildingSys) BuildingLvUp(building_uid int64, is_notify bool) common.RetCode {
	building, has := this.building_list[building_uid]
	LogDebug("enter BuildingLvUp:", building_uid)
	if !has {
		LogError("BuildingStartLvUp uid Error : ", building_uid)
		return common.RetCode_BuildingUidError
	}

	if building.GetLv() >= scheme.Commonmap[define.BuildingLvMax].Value {
		LogDebug("building.GetLv() >= scheme.Commonmap[define.BuildingLvMax].Value:", building.GetLv(), scheme.Commonmap[define.BuildingLvMax].Value)
		return common.RetCode_BuildingLvMax
	}

	//判断等级
	LogDebug("GetLvUpRoleLv ,GetLv", building.GetLvUpRoleLv(), this.owner.GetLv())
	if building.GetLvUpRoleLv() > this.owner.GetLv() {
		return common.RetCode_BuildingLvLimit
	}

	scheme_buildingLvup := scheme.BuildingLvUpGet(building.GetSchemeId(), building.GetLv()+1)
	if scheme_buildingLvup == nil {
		LogError("BuildingStartLvUp BuildingLvUpGet Error Id : ", building.GetSchemeId(), " Lv : ", building.GetLv()+1)
		return common.RetCode_SchemeData_Error
	}

	//魔魂
	if !this.owner.IsEnoughSoul(building.GetLvUpSoul()) {
		LogError("!this.owner.IsEnoughSoul(building.GetLvUpSoul()):", !this.owner.IsEnoughSoul(building.GetLvUpSoul()), building.GetLvUpSoul())
		return common.RetCode_RoleNotEnoughSoul
	}

	this.owner.CostSoul(building.GetLvUpSoul(), true, true)
	building.SetLv(building.GetLv() + 1)
	building.scheme_lvup_data = scheme_buildingLvup
	this.Save(building)

	if is_notify {
		this.send_notify(building)
	}

	this.static_building(building)

	//添加成就
	this.owner.AchievementAddNum(5, building.GetLv(), true)

	return common.RetCode_Success
}

func (this *BuildingSys) BuildingEditLv(building_uid int64, lv int32, is_notify bool) common.RetCode {
	building, has := this.building_list[building_uid]
	if !has {
		return common.RetCode_BuildingUidError
	}

	scheme_buildingLvup := scheme.BuildingLvUpGet(building.GetSchemeId(), lv)
	if scheme_buildingLvup == nil {
		return common.RetCode_SchemeData_Error
	}

	building.BuildingCache.SetLv(lv)
	building.scheme_lvup_data = scheme_buildingLvup

	this.Save(building)

	if is_notify {
		this.send_notify(building)
	}

	this.static_building(building)

	//添加成就
	this.owner.AchievementAddNum(5, building.GetLv(), true)

	return common.RetCode_Success
}

func (this *BuildingSys) send_notify(building *Building) {
	msg := &protocol.MsgBuildingInfoNotify{}
	msg.SetBuilding(&protocol.BuildingInfo{
		Uid:      proto.Int64(building.GetUid()),
		SchemeId: proto.Int32(building.GetSchemeId()),
		Lv:       proto.Int32(building.GetLv()),
	})

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
		return
	}

	global.SendMsg(int32(protocol.MsgCode_BuildingInfoNotify), this.owner.GetSid(), buf)
}

func (this *BuildingSys) static_building(building *Building) {
	msg := &static.MsgStaticBuilding{}
	msg.SetRoleUid(this.owner.GetUid())
	msg.SetUid(building.GetUid())
	msg.SetSchemeId(building.GetSchemeId())
	msg.SetLv(building.GetLv())

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_Building), buf)
}
