package soldier

import (
	. "Gameserver/cache"
	"Gameserver/global"
	. "Gameserver/logic"
	"common"
	d "common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"fmt"
	"galaxy"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

type Soldier struct {
	schemeId          int32
	num               int32
	level             int32
	stage             int32
	skillLevel        map[int32]int32
	exp               int32
	timestamp         int64
	evoSpeedTimeStamp int64
	active            int32
	autoId            int64
	CacheKey          string
}

type UserSoldiers struct {
	user          IRole
	Soldiers      map[int32]*Soldier //军营中的魔物，key soldierId
	SoldiersInMap map[int32]*Soldier //在地图中的魔物 key auto id, 一个魔物一个id, Soldier.num =1
	Events        map[int32]*SoldierEvents
	SummaryInMap  map[int32]int32 //地图数据摘要
}

const (
	soldierInCampKey      = "Role:%v:SoldierInCamp:%v"
	soldierInCampKeys     = "RoleSoldiers:%v"
	soldierInMapKey       = "Role:%v:Soldier:%v:InMap"
	soldierInMapKeys      = "Role:%v:SoldiersInMap"
	soldierInMapKeyAutoId = "Role:%v:SoldierInMapAutoId"
)

func GenSoldierKey(roleId int64, soldierId int32) string {
	return fmt.Sprintf(soldierInCampKey, roleId, soldierId)
}

//获取一个用户的所有魔物key
func GenSoldierListKey(roleId int64) string {
	return fmt.Sprintf(soldierInCampKeys, roleId)
}

func GenSoldierInMapKey(roleId int64, autoId int64) string {
	return fmt.Sprintf(soldierInMapKey, roleId, autoId)
}

func GenSoldierInMapListKey(roleId int64) string {
	return fmt.Sprintf(soldierInMapKeys, roleId)
}

func genSoldierInMapKeyAutoId(roleId int64) string {
	return fmt.Sprintf(soldierInMapKeyAutoId, roleId)
}

//获得升级表
func (this *UserSoldiers) getLvUpScheme(soldierId, level int32) *scheme.SoldierLvUp {
	return scheme.SoldierLvUpGet(soldierId, level)
}

func (this *UserSoldiers) getStageScheme(soldierId, stage int32) *scheme.SoldierStageUp {
	return scheme.SoldierStageUpGet(soldierId, stage)
}

//无消耗生成
func (this *UserSoldiers) SoldierCreateFree(soldierId, num int32) bool {
	if _, ok := this.Soldiers[soldierId]; ok {
		this.Soldiers[soldierId].num += num
		this.SoldierSummary()

		this.staticSoldier(this.Soldiers[soldierId])
		galaxy.LogDebug("SoldierCreateFree : ", this.Soldiers[soldierId])
		return this.Soldiers[soldierId].save(this.user)
	}
	return false
}

func (this *UserSoldiers) SoldierEditNum(soldierId, num int32) bool {
	if _, ok := this.Soldiers[soldierId]; ok {
		this.Soldiers[soldierId].num = num
		this.SoldierSummary()
		return this.Soldiers[soldierId].save(this.user)
	}
	return false
}

func (this *UserSoldiers) Init(user IRole) {
	this.user = user
	this.Soldiers = make(map[int32]*Soldier)
	this.Events = make(map[int32]*SoldierEvents)
	this.SummaryInMap = make(map[int32]int32)
}

func (this *UserSoldiers) Load() error {
	roleId := this.user.GetUid()

	err := this.loadSoldiersInCamp(roleId)
	if err != nil {
		return err
	}
	err = this.loadSoldiersInMap(roleId)
	if err != nil {
		return err
	}

	this.SoldierSummary()
	return err
}

func (this *UserSoldiers) loadSoldiersInCamp(roleId int64) error {
	this.Soldiers = make(map[int32]*Soldier)

	//获取所有keys
	resp, err := RedisCmd("SMEMBERS", GenSoldierListKey(roleId))
	if err != nil {
		galaxy.LogError(err)
		return err //没有
	}

	keys, _ := resp.List()
	for _, key := range keys {
		//获取数据
		resp, err := RedisCmd("GET", key)
		if err != nil {
			galaxy.LogError(err)
			return err //没有
		}
		if buff, _ := resp.Bytes(); buff != nil {
			soldierCache := &SoldierCache{}
			err = proto.Unmarshal(buff, soldierCache)
			soldier := &Soldier{CacheKey: key}
			soldier.readCache(soldierCache)
			this.Soldiers[soldierCache.GetSchemeId()] = soldier
		} else {
			RedisCmd("SREM", GenSoldierListKey(roleId), key)
		}
	}

	if err != nil {
		galaxy.LogError(err)
	}

	return err
}

//载入地图中的魔物
func (this *UserSoldiers) loadSoldiersInMap(roleId int64) error {
	this.SoldiersInMap = make(map[int32]*Soldier)

	//获取所有keys
	resp, err := RedisCmd("SMEMBERS", GenSoldierInMapListKey(roleId))
	if err != nil {
		galaxy.LogError(err)
		return err //没有
	}

	keys, _ := resp.List()
	for _, key := range keys {
		//获取数据
		resp, _ := RedisCmd("GET", key)

		if buff, _ := resp.Bytes(); buff != nil {
			soldierCache := &SoldierCache{}
			err = proto.Unmarshal(buff, soldierCache)
			soldier := &Soldier{CacheKey: key}
			soldier.readCache(soldierCache)
			this.SoldiersInMap[int32(soldierCache.GetAutoId())] = soldier
		} else {
			RedisCmd("SREM", GenSoldierInMapListKey(roleId), key)
		}
	}
	if err != nil {
		galaxy.LogError(err)
	}
	return err
}

func (this *UserSoldiers) SoldierUnlock(soldierId int32) {
	soldierStage := this.getStageScheme(soldierId, 1)
	if soldierStage == nil {
		galaxy.LogError("SoldierUnlock Failed Id:", soldierId)
		return
	}
	skillIds := strings.Split(soldierStage.SkillId, ";")

	soldier := new(Soldier)
	soldier.skillLevel = make(map[int32]int32)

	for _, v := range skillIds {
		temp, _ := strconv.Atoi(v)
		soldier.skillLevel[int32(temp)] = 1
	}

	soldier.CacheKey = GenSoldierKey(this.user.GetUid(), soldierId)
	soldier.schemeId = soldierId
	soldier.level = scheme.Commonmap[d.InitialSoldierLv].Value
	soldier.exp = scheme.Commonmap[d.InitialSoldierExp].Value
	soldier.stage = 1
	soldier.active = 1

	this.Soldiers[soldierId] = soldier
	soldier.save(this.user)
}

//魔物开放，在玩家升级后调用，或者创建新玩家后调用
func (this *UserSoldiers) SoldierFreshLv(kingLv int32) {
	for _, soldier := range this.Soldiers {
		soldierLvUp := this.getLvUpScheme(soldier.schemeId, soldier.level)
		soldierStage := this.getStageScheme(soldier.schemeId, soldier.stage)

		if soldier.exp == soldierLvUp.NeedExp && this.user.GetKingLv() >= soldierLvUp.LvUpKingLv && soldier.level != soldierStage.LvLimit {
			soldier.level += 1
			soldier.exp = 0
			this.Soldiers[soldier.schemeId].save(this.user)
			this.staticSoldier(soldier)
			galaxy.LogDebug("SoldierFreshLv : ", soldier)
			this.SyncSoldierInMapFromCamp(soldier.schemeId)
		}
	}
}

//魔物等级提升 吃勇士经验
func (this *UserSoldiers) SoldierLevelUp(soldierId int32, addedExp int32) bool {
	galaxy.LogDebug("into SoldierLevelUp:", soldierId, addedExp)

	if soldier, ok := this.Soldiers[soldierId]; ok {
		soldierLvUp := this.getLvUpScheme(soldierId, soldier.level)
		soldierStage := this.getStageScheme(soldierId, soldier.stage)

		max_lv := scheme.Commonmap[d.SoldierLvMax].Value
		if soldier.level >= max_lv {
			return false
		}

		if soldier.level > soldierStage.LvLimit {
			return false
		}

		if soldier.level == soldierStage.LvLimit && soldier.exp >= soldierLvUp.NeedExp {
			return false
		}

		if this.user.GetKingLv() < soldierLvUp.LvUpKingLv && soldier.exp >= soldierLvUp.NeedExp {
			return false
		}

		oldLevel := soldier.level
		total_exp := soldier.exp + addedExp
		for total_exp > 0 {
			total_exp -= soldierLvUp.NeedExp
			if total_exp >= 0 {
				if this.user.GetKingLv() < soldierLvUp.LvUpKingLv {
					soldier.exp = soldierLvUp.NeedExp
					break
				}

				if soldier.level == soldierStage.LvLimit {
					soldier.exp = soldierLvUp.NeedExp
					break
				}

				soldierLvUp = this.getLvUpScheme(soldierId, soldier.level+1)
				if soldierLvUp == nil {
					return false
				}

				soldier.level += 1
				if soldier.level >= max_lv {
					soldier.exp = 0
					break
				}

				if total_exp == 0 {
					soldier.exp = total_exp
				}
			} else {
				soldier.exp = total_exp + this.getLvUpScheme(soldierId, soldier.level).NeedExp
			}
		}

		var flag int32 = 0
		if false == soldier.save(this.user) {
			flag = 1
		}

		if soldier.level > oldLevel {
			this.SyncSoldierInMapFromCamp(soldierId)
			this.SoldierSendAll()
			this.staticSoldier(soldier)
			galaxy.LogDebug("SoldierLevelUp : ", soldier)
		}

		//send msg
		ret := &protocol.MsgSoldierLvRet{}
		ret.SetRetCode(flag)
		ret.SetLevel(soldier.level)
		ret.SetExp(soldier.exp)

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return false
		}

		global.SendMsg(int32(protocol.MsgCode_SoldierLvRet), this.user.GetSid(), buf)
		//end
		return true
	}

	return false
}

func (this *UserSoldiers) SoldierEditLv(soldierId int32, lv int32) bool {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		soldierLvUp := this.getLvUpScheme(soldierId, lv)
		if soldierLvUp == nil {
			return false
		}

		soldierStage := this.getStageScheme(soldierId, soldierLvUp.Stage)
		if soldierStage == nil {
			return false
		}
		soldier.level = lv
		soldier.timestamp = 0
		soldier.evoSpeedTimeStamp = 0
		soldier.stage = soldierLvUp.Stage
		if lv >= scheme.Commonmap[d.SoldierLvMax].Value {
			soldier.exp = 0
		}

		if soldier.exp > soldierLvUp.NeedExp {
			soldier.exp = soldierLvUp.NeedExp - 1
		}

		soldier.save(this.user)
		this.staticSoldier(soldier)
		galaxy.LogDebug("SoldierEditLv : ", soldier)
		this.SyncSoldierInMapFromCamp(soldierId)
		this.SoldierSendAll()
		if magic_type, magic_id := this.user.GetMagicQueue(); magic_type == common.RTYPE_MAGIC_SOLDIER && int32(magic_id) == soldierId {
			this.user.ResetMagicQueue(true)
		}

		this.user.AchievementAddNum(12, 1, false)
		return true
	}
	return false
}

//魔物等级升阶
//选中的魔使id
func (this *UserSoldiers) SoldierUpgrade(soldierId int32, heroUid int64, usedCoin bool) (common.RetCode, int64, int64) {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		galaxy.LogDebug("upgrade before: stage:", this.Soldiers[soldierId].stage)

		soldierStage := this.getStageScheme(soldierId, soldier.stage)

		galaxy.LogDebug("LvLimit:", soldierStage.LvLimit)
		//是否达到最大级
		if soldier.level != soldierStage.LvLimit && soldier.exp != this.getLvUpScheme(soldierId, soldier.level).NeedExp-1 {
			galaxy.LogDebug("是否达到最大级:", soldier.level, soldierStage.LvLimit, soldier.exp, this.getLvUpScheme(soldierId, soldier.level).NeedExp-1)
			return common.RetCode_Fail, 0, 0
		}

		//是否有下一阶
		if soldierStage.NextStageId == -1 {
			galaxy.LogDebug("soldierStage.NextStageId:", soldierStage.NextStageId)
			return common.RetCode_Fail, 0, 0
		}

		//是否有坑
		if usedCoin == false {
			if rtype, id := this.user.GetMagicQueue(); id != 0 {
				galaxy.LogDebug("GetMagicQueue:", rtype, id)
				return common.RetCode_QueueFull, 0, 0 //同时只能有一个魔物或魔使升级
			}
		}

		itemids := strings.Split(soldierStage.EvoNeedItemId, ";")
		itemnums := strings.Split(soldierStage.EvoNeedItemNum, ";")

		for k, id := range itemids {
			sid, _ := strconv.Atoi(id)
			num, _ := strconv.Atoi(itemnums[k])
			if !this.user.ItemIsEnough(int32(sid), int32(num)) {
				galaxy.LogDebug("item is not enough:", sid, " num:", num)
				return common.RetCode_Fail, 0, 0
			}
		}

		//检查魔使是否达标
		hero := this.user.HeroGet(heroUid)
		if hero == nil || hero.GetSchemeId() != soldierStage.EvoNeedMagicHeroId || hero.GetLv() < soldierStage.EvoNeedMagicHeroLv || hero.GetRank() < soldierStage.EvoNeedMagicHeroRank {
			galaxy.LogDebug("hero is not fixes:")
			return common.RetCode_Fail, 0, 0
		}

		for k, id := range itemids {
			sid, _ := strconv.Atoi(id)
			num, _ := strconv.Atoi(itemnums[k])
			this.user.ItemCost(int32(sid), int32(num), true)
		}

		this.user.HeroCost(heroUid, true)

		//增加升阶时间戳
		soldier.timestamp = Time() + int64(soldierStage.EvoNeedTime)
		soldier.evoSpeedTimeStamp = Time()
		this.user.SetMagicQueue(common.RTYPE_MAGIC_SOLDIER, int64(soldierId), true)
		galaxy.LogDebug("SetMagicQueue:", common.RTYPE_MAGIC_SOLDIER, int64(soldierId))
		soldier.save(this.user)
		if usedCoin == true && this.SoldierRemoveUpgradeTime(soldierId, true) {
			return common.RetCode_Success, 0, 0
		}

		return common.RetCode_Success, soldier.timestamp, soldier.evoSpeedTimeStamp

	} else {
		return common.RetCode_Unable, 0, 0
	}
}

//减少升阶时间 每次扣一点 减少300秒
func (this *UserSoldiers) SoldierCutDownUpgradeTime(soldierId int32) (common.RetCode, int64) {
	//是否升级完成，时间未到
	if soldier, ok := this.Soldiers[soldierId]; ok {
		if soldier.timestamp == 0 {
			galaxy.LogDebug("soldier.timestamp:", soldier.timestamp)
			return common.RetCode_Fail, 0
		} else if soldier.timestamp <= Time() {
			//已经进化完成
			this.SoldierFinishUpgrade(soldierId)
			galaxy.LogDebug("soldier.timestamp <= Time():", soldier.timestamp, Time())
			return common.RetCode_Fail, 0
		}

		EvoSpeedLimit := scheme.KingLvUpmap[this.user.GetKingLv()].EvoSpeedLimit
		EvoSpeedValue := scheme.Commonmap[d.EvoSpeedValue].Value
		EvoSpeedRecover := scheme.Commonmap[d.EvoSpeedRecover].Value

		if soldier.evoSpeedTimeStamp+int64(EvoSpeedRecover*EvoSpeedLimit) < Time() {
			galaxy.LogDebug("fix evo time before:", soldier.evoSpeedTimeStamp)
			soldier.evoSpeedTimeStamp = Time() - int64(EvoSpeedRecover*EvoSpeedLimit)
			galaxy.LogDebug("fix evo time after:", soldier.evoSpeedTimeStamp, EvoSpeedRecover, EvoSpeedLimit, EvoSpeedRecover*EvoSpeedLimit)
		}

		diffTime := Time() - soldier.evoSpeedTimeStamp - int64(EvoSpeedRecover)
		if diffTime < 0 {
			galaxy.LogError("diffTime:", diffTime, soldier.evoSpeedTimeStamp, int64(EvoSpeedRecover))
			return common.RetCode_Fail, 0
		}

		soldier.evoSpeedTimeStamp += int64(EvoSpeedRecover)
		soldier.timestamp -= int64(EvoSpeedValue)
		if soldier.timestamp <= Time() {
			//进化完成
			this.SoldierFinishUpgrade(soldierId)
			galaxy.LogDebug("soldier.timestamp <= Time()", soldier.timestamp, Time())
			return common.RetCode_Success, 0
		}

		soldier.save(this.user)
		return common.RetCode_Success, soldier.evoSpeedTimeStamp
	}
	return common.RetCode_Fail, 0
}

//使用金币移除升级时间，马上升级
func (this *UserSoldiers) SoldierRemoveUpgradeTime(soldierId int32, fromSoldierUpgrade bool) bool {
	if soldier, ok := this.Soldiers[soldierId]; ok {

		if soldier.timestamp == 0 || soldier.timestamp < Time() {
			return false
		}
		soldierStage := this.getStageScheme(soldierId, soldier.stage)
		coin := ResourceToCoin(common.RTYPE_TIME, int32(soldier.timestamp-Time()))
		if this.user.IsEnoughGold(coin) {
			this.user.CostGold(coin, true, false)
			soldier.level++
			soldier.timestamp = 0
			soldier.evoSpeedTimeStamp = 0
			soldier.exp = 0
			soldier.stage = scheme.SoldierStageUpmap[soldierStage.NextStageId].Stage

			this.staticSoldier(soldier)
			if fromSoldierUpgrade {
				this.user.StaticPayLog(int32(static.PayType_evolutionOnekey), 0, coin)
			} else {
				this.user.StaticPayLog(int32(static.PayType_evolutionSpeedup), 0, coin)
			}

			soldier.save(this.user)
			//更新地图中魔物阶数
			this.SyncSoldierInMapFromCamp(soldierId)
			this.user.ResetMagicQueue(true)
			//完成成就
			this.user.AchievementAddNum(12, 1, false)
			return true
		}
	}
	return false
}

//完成升阶
func (this *UserSoldiers) SoldierFinishUpgrade(soldierId int32) (bool, int32) {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		soldierStage := this.getStageScheme(soldierId, soldier.stage)

		//是否达到最大级
		if soldier.level != soldierStage.LvLimit {
			return false, 0
		}

		if soldier.timestamp > Time() {
			return false, 0
		}

		soldier.timestamp = 0
		soldier.evoSpeedTimeStamp = 0
		soldier.stage = scheme.SoldierStageUpmap[soldierStage.NextStageId].Stage
		soldier.exp = 0

		this.staticSoldier(soldier)
		galaxy.LogDebug("SoldierFinishUpgrade : ", soldier)
		soldier.save(this.user)
		//更新地图中魔物阶数
		this.SyncSoldierInMapFromCamp(soldierId)
		this.user.ResetMagicQueue(true)

		//完成成就
		this.user.AchievementAddNum(12, 1, false)

		return true, soldier.stage
	}
	return false, 0
}

// 技能强化
func (this *UserSoldiers) SoldierSkillLevelUp(soldierId int32, skillId int32) bool {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		galaxy.LogDebug("end SoldierSkillLevelUp: ", soldier)

		if _, ok := soldier.skillLevel[skillId]; !ok {
			galaxy.LogDebug("skillId error:", skillId)
			return false
		}

		//是否达到最大级
		if soldier.skillLevel[skillId] >= scheme.Skillmap[skillId].LvMax {
			galaxy.LogDebug("skill level error:", soldier.skillLevel[skillId], scheme.Skillmap[skillId].LvMax)
			return false
		}

		//技能等级不能大于魔物等级
		if soldier.skillLevel[skillId] >= soldier.level {
			galaxy.LogDebug("技能等级不能大于魔物等级")
			return false
		}

		//材料是否够
		var skillLvUp *scheme.SkillLvUp
		for _, v := range scheme.SkillLvUpmap {
			if v.BaseId == skillId && v.Lv == soldier.skillLevel[skillId] {
				skillLvUp = v
				galaxy.LogDebug("scheme found skillLvUp :", v)
				break
			}
		}

		if skillLvUp == nil {
			galaxy.LogDebug("scheme skillLvUp is empty:", skillId, soldier.skillLevel[skillId], skillLvUp)
			return false
		}
		if !this.user.ItemIsEnough(skillLvUp.NeedItemId, skillLvUp.NeedItemNum) {
			galaxy.LogDebug("item is not enough")
			return false
		}

		this.user.ItemCost(skillLvUp.NeedItemId, skillLvUp.NeedItemNum, true)

		//技能升级
		soldier.skillLevel[skillId]++
		soldier.save(this.user)

		//完成成就
		this.user.AchievementAddNum(11, soldier.skillLevel[skillId], true)

		galaxy.LogDebug("end SoldierSkillLevelUp:", soldier)
		return true

	} else {
		return false
	}
}

func (this *UserSoldiers) SoldierEditSkillLv(soldierId int32, skillId int32, lv int32) bool {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		if _, ok := soldier.skillLevel[skillId]; !ok {
			return false
		}

		//是否达到最大级
		if lv >= 20 {
			return false
		}

		var skillLvUp *scheme.SkillLvUp
		for _, v := range scheme.SkillLvUpmap {
			if v.BaseId == skillId && v.Lv == lv {
				skillLvUp = v
				break
			}
		}

		if skillLvUp == nil {
			return false
		}

		//技能升级
		soldier.skillLevel[skillId] = lv
		this.staticSoldier(soldier)
		galaxy.LogDebug("SoldierEditSkillLv : ", soldier)
		soldier.save(this.user)

		//完成成就
		this.user.AchievementAddNum(11, soldier.skillLevel[skillId], true)
		return true

	} else {
		return false
	}
}

//派兵
func (this *UserSoldiers) SoldierDispatchByAutoId(soldierId int32, autoId int64) ISoldier {
	var newSoldier *Soldier

	if this.SoldiersInMap == nil {
		this.SoldiersInMap = make(map[int32]*Soldier)
	}

	var soldierNum int32 = 1

	if soldier, ok := this.Soldiers[soldierId]; ok {
		if soldier.num >= soldierNum {
			soldier.num -= soldierNum
			_, err := RedisCmd("SET", genSoldierInMapKeyAutoId(this.user.GetUid()), autoId)
			if err != nil {
				galaxy.LogError(err)
				return nil
			}

			newSoldier = soldier.clone()
			newSoldier.autoId = autoId
			newSoldier.num = soldierNum
			newSoldier.CacheKey = GenSoldierInMapKey(this.user.GetUid(), autoId)
			this.SoldiersInMap[int32(autoId)] = newSoldier

			soldier.save(this.user)
			if newSoldier.save(this.user) == false {
				return nil
			}

			return newSoldier
		}
	}
	return nil
}

//派兵
func (this *UserSoldiers) SoldierDispatch(soldierId int32) ISoldier {
	var newSoldier *Soldier

	if this.SoldiersInMap == nil {
		this.SoldiersInMap = make(map[int32]*Soldier)
	}

	var soldierNum int32 = 1

	if soldier, ok := this.Soldiers[soldierId]; ok {
		if soldier.GetNum() >= soldierNum {
			soldier.num -= soldierNum
			resp, err := RedisCmd("INCR", genSoldierInMapKeyAutoId(this.user.GetUid()))
			if err != nil {
				galaxy.LogError(err)
				return nil
			}

			temp, err := resp.Int64()
			if err != nil {
				galaxy.LogError(err)
				return nil
			}

			autoId := int64(temp)
			newSoldier = soldier.clone()
			newSoldier.autoId = autoId
			newSoldier.num = soldierNum
			newSoldier.CacheKey = GenSoldierInMapKey(this.user.GetUid(), autoId)
			this.SoldiersInMap[int32(autoId)] = newSoldier

			soldier.save(this.user)
			if newSoldier.save(this.user) == false {
				return nil
			}

			return newSoldier
		}
	}
	return nil
}

//收兵
func (this *UserSoldiers) SoldierWithdraw() {
	for autoId, soldier := range this.SoldiersInMap {
		soldierId := soldier.schemeId
		this.Soldiers[soldierId].num += soldier.num
		this.Soldiers[soldierId].save(this.user)
		this.DeleteSoldierInMapFromDb(autoId)
	}
	this.SoldierSummary()
}

func (this *UserSoldiers) SoldierGetInMap(autoId int32) ISoldier {
	if soldier, ok := this.SoldiersInMap[autoId]; ok {
		return soldier
	}
	return nil
}

func (this *UserSoldiers) FillSoldierFightInfo(auto_id int32) *protocol.SoldierFightInfo {
	soldier := this.SoldierGetInMap(auto_id)
	if soldier == nil {
		return nil
	}

	msg := new(protocol.SoldierFightInfo)
	msg.SetAutoId(soldier.GetAutoId())
	msg.SetSchemeId(soldier.GetSchemeId())
	msg.SetLevel(soldier.GetLevel())
	msg.SetStage(soldier.GetStage())

	skillLevels := make([]*protocol.SkillLevel, 0)
	var i int32 = 0
	for k, v := range soldier.GetSkillLevel() {
		skillLevel := &protocol.SkillLevel{}
		skillLevel.SkillId = proto.Int32(k)
		skillLevel.SkillLevel = proto.Int32(v)
		skillLevels = append(skillLevels, skillLevel)
		i++
	}
	msg.SetSkillLevel(skillLevels)

	return msg
}

func (this *UserSoldiers) SoldierGetInCamp(id int32) ISoldier {
	if soldier, ok := this.Soldiers[id]; ok {
		return soldier
	}
	return nil
}

//地图上魔物信息
func (this *UserSoldiers) SoldierSummary() {
	count := make(map[int32]int32)
	for _, v := range this.SoldiersInMap {
		count[v.GetSchemeId()] += v.GetNum()
	}

	for _, v := range this.Soldiers {
		count[v.GetSchemeId()] += v.GetNum()
	}

	this.SummaryInMap = count
}

func (this *UserSoldiers) SoldierNum(schemeId int32) int32 {
	if v, ok := this.SummaryInMap[schemeId]; ok {
		return v
	}
	return 0
}

//同步地图中魔物的等级和阶数
func (this *UserSoldiers) SyncSoldierInMapFromCamp(soldierIdInCamp int32) {
	soldierId := soldierIdInCamp
	if soldier, ok := this.Soldiers[soldierId]; ok {
		for _, v := range this.SoldiersInMap {
			if v.GetSchemeId() == soldierId {
				v.skillLevel = soldier.skillLevel
				v.exp = soldier.exp
				v.timestamp = soldier.timestamp
				v.evoSpeedTimeStamp = soldier.evoSpeedTimeStamp
				v.active = soldier.active
				v.level = soldier.level
				v.stage = soldier.stage
				v.save(this.user)
			}
		}
	}
}

//删除地图上的魔物
func (this *UserSoldiers) DeleteSoldierInMapFromDb(autoId int32) bool {
	if soldier, ok := this.SoldiersInMap[autoId]; ok {
		if _, err := RedisCmd("DEL", GenSoldierInMapKey(this.user.GetUid(), soldier.GetAutoId())); err != nil {
			galaxy.LogError(err)
			return false
		}
		if _, err := RedisCmd("SREM", GenSoldierInMapListKey(this.user.GetUid()), soldier.CacheKey); err != nil {
			galaxy.LogError(err)
			return false
		}
		delete(this.SoldiersInMap, autoId)
		return true
	}
	return false
}

func (this *UserSoldiers) SoldierSendAll() {
	list1 := make([]*protocol.Soldier, len(this.Soldiers))

	var i int32 = 0
	for _, v := range this.Soldiers {
		list1[i] = v.toProtocol()
		i++
	}

	list2 := make([]*protocol.Soldier, len(this.SoldiersInMap))

	i = 0
	for _, v := range this.SoldiersInMap {
		list2[i] = v.toProtocol()
		i++
	}

	m := new(protocol.MsgSoldierAllRet)
	m.SetSoldiersInCamp(list1)
	m.SetSoldiersInMap(list2)

	buf, err := proto.Marshal(m)
	if err != nil {
		galaxy.LogError(err)
		return
	}

	global.SendMsg(int32(protocol.MsgCode_SoldierAllRet), this.user.GetSid(), buf)
}

func (this *UserSoldiers) FillAllSoldiersInfo() *protocol.AllSoldiers {
	list1 := make([]*protocol.Soldier, len(this.Soldiers))

	var i int32 = 0
	for _, v := range this.Soldiers {
		list1[i] = v.toProtocol()
		i++
	}

	list2 := make([]*protocol.Soldier, len(this.SoldiersInMap))

	i = 0
	for _, v := range this.SoldiersInMap {
		list2[i] = v.toProtocol()
		i++
	}

	m := new(protocol.AllSoldiers)
	m.SetSoldiersInCamp(list1)
	m.SetSoldiersInMap(list2)
	return m
}

func (this *UserSoldiers) staticSoldier(soldier *Soldier) {
	msg := &static.MsgStaticSoldier{}
	msg.SetRoleUid(this.user.GetUid())
	msg.SetSchemeId(soldier.GetSchemeId())
	msg.SetNum(soldier.GetNum())
	msg.SetLv(soldier.GetLevel())
	msg.SetStage(soldier.GetStage())

	buf, err := proto.Marshal(msg)
	if err != nil {
		galaxy.LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_Soldier), buf)
}
