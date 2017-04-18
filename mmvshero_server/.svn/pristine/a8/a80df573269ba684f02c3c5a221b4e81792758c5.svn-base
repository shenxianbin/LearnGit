package soldier

import (
	"Gameserver/global"
	. "Gameserver/logic"
	. "common/cache"
	d "common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"fmt"
	. "galaxy"
	"math/rand"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

type Soldier struct {
	schemeId   int32
	num        int32
	level      int32
	stage      int32
	skillLevel map[int32]int32

	CacheKey string
}

type UserSoldiers struct {
	user     IRole
	Soldiers map[int32]*Soldier
}

const (
	soldierInCampKey  = "Role:%v:SoldierInCamp:%v"
	soldierInCampKeys = "RoleSoldiers:%v"
)

func GenSoldierKey(roleId int64, soldierId int32) string {
	return fmt.Sprintf(soldierInCampKey, roleId, soldierId)
}

//获取一个用户的所有魔物key
func GenSoldierListKey(roleId int64) string {
	return fmt.Sprintf(soldierInCampKeys, roleId)
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

		this.staticSoldier(this.Soldiers[soldierId])
		return this.Soldiers[soldierId].save(this.user)
	}

	return false
}

func (this *UserSoldiers) SoldierEditNum(soldierId, num int32) bool {
	if _, ok := this.Soldiers[soldierId]; ok {
		this.Soldiers[soldierId].num = num
		return this.Soldiers[soldierId].save(this.user)
	}
	return false
}

func (this *UserSoldiers) Init(user IRole) {
	this.user = user
	this.Soldiers = make(map[int32]*Soldier)
}

func (this *UserSoldiers) Load() error {
	this.Soldiers = make(map[int32]*Soldier)
	roleId := this.user.GetUid()
	//获取所有keys
	resp, err := RedisCmd("SMEMBERS", GenSoldierListKey(roleId))
	if err != nil {
		LogError(err)
		return err //没有
	}

	keys, _ := resp.List()
	for _, key := range keys {
		//获取数据
		resp, err := RedisCmd("GET", key)
		if err != nil {
			LogError(err)
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
		LogError(err)
	}

	return err
}

func (this *UserSoldiers) SoldierRandomType(ban_id []int32) int32 {
	ban := make(map[int32]bool)
	if ban_id != nil {
		for _, v := range ban_id {
			ban[v] = true
		}
	}

	list := make([]int32, 0)
	for _, v := range this.Soldiers {
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

func (this *UserSoldiers) SoldierUnlock(soldierId int32) {
	if _, has := this.Soldiers[soldierId]; has {
		return
	}

	soldierStage := this.getStageScheme(soldierId, 1)
	if soldierStage == nil {
		LogError("SoldierUnlock Failed Id:", soldierId)
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
	soldier.level = scheme.Commonmap[d.SoldierLvInitial].Value
	soldier.stage = 1

	this.Soldiers[soldierId] = soldier
	soldier.save(this.user)
}

//魔物等级提升 吃勇士经验
func (this *UserSoldiers) SoldierLevelUp(soldierId int32, item_scheme_id int32, num int32) bool {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		soldierBase, _ := scheme.Soldiermap[soldierId]
		soldierLvUp := this.getLvUpScheme(soldierId, soldier.level)

		if item_scheme_id != soldierBase.NeedItemId {
			return false
		}

		max_lv := scheme.Commonmap[d.SoldierLvMax].Value
		if soldier.level >= max_lv {
			return false
		}

		if this.user.GetLv() < soldierLvUp.LvUpRoleLv {
			return false
		}

		if num < soldierLvUp.NeedExp {
			return false
		}

		soldierLvUp = this.getLvUpScheme(soldierId, soldier.level+1)
		if soldierLvUp == nil {
			return false
		}

		soldierStage := this.getStageScheme(soldierId, soldierLvUp.Stage)
		if soldierStage == nil {
			LogError("SoldierStage Null : SchemeId [", soldier.schemeId, "] Stage[", soldierLvUp.Stage, "]")
			return false
		}

		soldier.level += 1
		soldier.stage = soldierLvUp.Stage
		this.user.AddExp(soldierLvUp.AddRoleExp, true, true)

		var flag int32 = 0
		if false == soldier.save(this.user) {
			flag = 1
		}

		this.staticSoldier(soldier)

		//send msg
		ret := &protocol.MsgSoldierLvRet{}
		ret.SetSchemeId(soldier.schemeId)
		ret.SetRetCode(flag)
		ret.SetLevel(soldier.level)

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return false
		}

		global.SendMsg(int32(protocol.MsgCode_SoldierLvRet), this.user.GetSid(), buf)
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
		soldier.stage = soldierLvUp.Stage

		soldier.save(this.user)
		this.staticSoldier(soldier)

		// this.user.AchievementAddNum(12, 1, false)
		return true
	}
	return false
}

// 技能强化
func (this *UserSoldiers) SoldierSkillLevelUp(soldierId int32, skillId int32) bool {
	if soldier, ok := this.Soldiers[soldierId]; ok {
		if _, ok := soldier.skillLevel[skillId]; !ok {
			return false
		}

		//是否达到最大级
		if soldier.skillLevel[skillId] >= scheme.Skillmap[skillId].LvMax {
			return false
		}

		//技能等级不能大于魔物等级
		if soldier.skillLevel[skillId] >= soldier.level {
			return false
		}

		//材料是否够
		var skillLvUp *scheme.SkillLvUp
		for _, v := range scheme.SkillLvUpmap {
			if v.BaseId == skillId && v.Lv == soldier.skillLevel[skillId] {
				skillLvUp = v
				break
			}
		}

		if skillLvUp == nil {
			return false
		}

		if !this.user.IsEnoughSoul(skillLvUp.NeedSoul) {
			return false
		}

		this.user.CostSoul(skillLvUp.NeedSoul, true, true)

		//技能升级
		soldier.skillLevel[skillId]++
		soldier.save(this.user)

		//完成成就
		this.user.AchievementAddNum(3, soldier.skillLevel[skillId], true)

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
		soldier.save(this.user)

		//完成成就
		this.user.AchievementAddNum(3, soldier.skillLevel[skillId], true)
		return true

	} else {
		return false
	}
}

func (this *UserSoldiers) SoldierGetInCamp(id int32) ISoldier {
	if soldier, ok := this.Soldiers[id]; ok {
		return soldier
	}
	return nil
}

func (this *UserSoldiers) SoldierAllId() []int32 {
	var ret []int32 = make([]int32, 0)
	for _, soldier := range this.Soldiers {
		ret = append(ret, soldier.GetSchemeId())
	}
	return ret
}

func (this *UserSoldiers) SoldierNum(schemeId int32) int32 {
	if v, ok := this.Soldiers[schemeId]; ok {
		return v.GetNum()
	}
	return 0
}

func (this *UserSoldiers) FillAllSoldiersInfo() *protocol.AllSoldiers {
	list1 := make([]*protocol.Soldier, len(this.Soldiers))

	var i int32 = 0
	for _, v := range this.Soldiers {
		list1[i] = v.toProtocol()
		i++
	}

	m := new(protocol.AllSoldiers)
	m.SetSoldiersInCamp(list1)
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
		LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_Soldier), buf)
}
