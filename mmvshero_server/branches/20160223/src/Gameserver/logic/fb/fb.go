package fb

import (
	. "Gameserver/cache"
	// "Gameserver/global"
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	// "common"
	"common/protocol"
	"common/scheme"
	"fmt"
	"galaxy"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
)

type Fb struct {
	schemeId      int32
	timestamp     int64           //the index of time interval, example 201509101,201509101
	attackedTimes map[int32]int32 //key difficulty -1
}

type UserFbs struct {
	user IRole
	Fbs  map[int32]*Fb //key Fb schemeId
}

const (
	key  = "Role:%v:Fb:%v"
	keys = "Role:%v:Fbs"
)

func (this *UserFbs) key(roleId int64, schemeId int32) string {
	return fmt.Sprintf(key, roleId, schemeId)
}

func (this *UserFbs) keys(roleId int64) string {
	return fmt.Sprintf(keys, roleId)
}
func (this *UserFbs) Init(user IRole) {
	this.user = user
	this.Fbs = make(map[int32]*Fb)
}

func (this *UserFbs) Load() error {
	roleId := this.user.GetUid()
	if this.Fbs == nil {
		this.Fbs = make(map[int32]*Fb)
	}
	// galaxy.LogDebug("load this.Fbs:", this.Fbs, len(this.Fbs))
	//获取所有keys
	resp, err := RedisCmd("SMEMBERS", this.keys(roleId))
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
			cache := &FbCache{}
			err = proto.Unmarshal(buff, cache)
			this.Fbs[cache.GetSchemeId()] = this.newFb(cache)
		} else {
			galaxy.LogError("SREM: ", this.keys(roleId), key)
			RedisCmd("SREM", this.keys(roleId), key)
		}
	}

	if err != nil {
		galaxy.LogError(err)
	}

	return nil
}

func (this *UserFbs) newFbCache(m *Fb) *FbCache {
	c := &FbCache{}
	c.SchemeId = proto.Int32(m.schemeId)
	c.Timestamp = proto.Int64(m.timestamp)

	c.AttackedTimes = make(map[int32]int32)
	for k, v := range m.attackedTimes {
		c.AttackedTimes[k] = v
	}

	return c
}

func (this *UserFbs) newFb(cache *FbCache) *Fb {
	s := &Fb{}
	s.schemeId = cache.GetSchemeId()
	s.timestamp = cache.GetTimestamp()

	s.attackedTimes = make(map[int32]int32)
	for k, v := range cache.GetAttackedTimes() {
		s.attackedTimes[k] = v
	}

	return s
}

func (this *UserFbs) LoadFirst() {
	if this.Fbs != nil {
		return
	}

	galaxy.LogDebug("loadFirst", this.Fbs)
	for id, v := range scheme.Fbmap {
		fb := &Fb{}
		fb.schemeId = id
		fb.timestamp = 0
		fb.attackedTimes = make(map[int32]int32)

		this.Fbs[v.Id] = fb
		this.save(id)
	}
	galaxy.LogDebug("end loadFirst", this.Fbs)
}

//update attacked times
func (this *UserFbs) refreshFbs() {
	timestamp := this.getTimestamp()
	for _, fb := range this.Fbs {
		//todo cache current time interval
		temp := timestamp + this.getCurrentTimeInterval(fb.schemeId)

		if fb.timestamp != temp {
			//reset attackedTimes
			for k, _ := range fb.attackedTimes {
				fb.attackedTimes[k] = 0
			}
		}
	}
}

func (this *UserFbs) FbAll() *protocol.MsgFbAllRet {
	this.LoadFirst()
	// galaxy.LogDebug("this.Fbs:", this.Fbs)

	list1 := make([]*protocol.Fb, len(this.Fbs))

	var i int32 = 0
	for _, v := range this.Fbs {
		list1[i] = v.FillInfo()
		i++
	}

	m := new(protocol.MsgFbAllRet)
	m.SetFbs(list1)
	return m
}

func (this *Fb) FillInfo() *protocol.Fb {
	m := new(protocol.Fb)
	m.SchemeId = proto.Int32(this.schemeId)
	m.AttackedTimes = make([]*protocol.AttackedTimes, len(this.attackedTimes))

	var i int32 = 0
	for k, v := range this.attackedTimes {
		temp := &protocol.AttackedTimes{}
		temp.Difficulty = proto.Int32(k)
		temp.Times = proto.Int32(v)
		m.AttackedTimes[i] = temp
		i++
	}
	return m
}

func (this *UserFbs) FbBegin(schemeId, difficulty int32) bool {
	galaxy.LogDebug("enter FbBegin ")
	if fb, ok := this.Fbs[schemeId]; ok {
		if scheme.Fbmap[schemeId].NeededLv > this.user.GetLv() {
			galaxy.LogDebug("current lv is under of neededLv:", scheme.Fbmap[schemeId].NeededLv, this.user.GetLv())
			return false
		}

		//check opening time interval and update attacked times
		currentTimeInterval := this.getCurrentTimeInterval(schemeId)
		if currentTimeInterval <= 0 {
			galaxy.LogDebug("fb has not opened")
			return false
		}
		attackTimes := ExtractIntMap(scheme.Fbmap[schemeId].AttackTimes)
		if fb.attackedTimes[difficulty-1] >= attackTimes[difficulty-1] {
			galaxy.LogDebug("reach max times of the attack:", attackTimes, difficulty)
			return false
		}

		orders := ExtractIntMap(scheme.Fbmap[schemeId].LeastCostOrder)
		winOrders := ExtractIntMap(scheme.Fbmap[schemeId].ResultCostOrder)
		LeastRoleExps := ExtractIntMap(scheme.Fbmap[schemeId].LeastRoleExp)

		order, ok := orders[difficulty-1]
		if !ok {
			galaxy.LogError("not found the order in scheme :", orders, difficulty)
			return false
		}

		winOrder, ok := winOrders[difficulty-1]
		if !ok {
			galaxy.LogError("not found the winOrder in scheme :", winOrders, difficulty)
			return false
		}

		if !this.user.IsEnoughOrder(winOrder) {
			galaxy.LogDebug("VictoryCostOrder is not enough", this.user.GetOrder())
			return false
		}

		LeastRoleExp, ok := LeastRoleExps[difficulty-1]
		if !ok {
			galaxy.LogDebug("not found the LeastRoleExp in scheme ", LeastRoleExps, difficulty)
			return false
		}

		this.user.CostOrder(order, true, true)
		this.user.AddExp(LeastRoleExp, true, true)

		fb.timestamp = this.getTimestamp() + currentTimeInterval
		fb.attackedTimes[difficulty-1]++

		this.save(schemeId)
		return true
	}
	galaxy.LogError("not found:", schemeId, this.Fbs)
	return false
}

func (this *UserFbs) FbFinish(schemeId int32, difficulty int32, isPassed bool, caughtNpcs []*protocol.CaughtNpc) *protocol.MsgFbFinishRet {
	ret := &protocol.MsgFbFinishRet{}
	ret.SetRetCode(1)

	//完成任务
	this.user.MissionAddNum(8, 1, 0)

	//添加成就
	this.user.AchievementAddNum(19, 1, false)

	if _, ok := this.Fbs[schemeId]; ok {
		//send awards for caught npc
		awards := make([]int32, 0)
		for _, npc := range caughtNpcs {
			npcScheme, ok := scheme.NPCmap[npc.GetSchemeId()]
			if !ok {
				galaxy.LogError("not found npc in scheme:", npc.GetSchemeId())
				return ret
			}

			for awardId, _ := range GetAwardsByProp(ExtractIntPairMap(npcScheme.Drop)) {
				Award(awardId, this.user, true)
				awards = append(awards, awardId)
			}
		}

		if isPassed == false {
			ret.SetRetCode(0)
			this.save(schemeId)
			return ret
		}

		//send passed awards
		awardsOdds := strings.Split(scheme.Fbmap[schemeId].AwardId, "|")

		for awardId, _ := range GetAwardsByProp(ExtractIntPairMap(awardsOdds[difficulty-1])) {
			Award(awardId, this.user, true)
			awards = append(awards, awardId)
		}

		//扣除体力
		orders := ExtractIntMap(scheme.Fbmap[schemeId].LeastCostOrder)
		ResultOrders := ExtractIntMap(scheme.Fbmap[schemeId].ResultCostOrder)

		LeastRoleExps := ExtractIntMap(scheme.Fbmap[schemeId].LeastRoleExp)
		ResultRoleExps := ExtractIntMap(scheme.Fbmap[schemeId].ResultRoleExp)

		order := ResultOrders[difficulty-1] - orders[difficulty-1]
		this.user.CostOrder(order, true, true)

		//增加经验
		exp := ResultRoleExps[difficulty-1] - LeastRoleExps[difficulty-1]
		this.user.AddExp(exp, true, true)

		this.save(schemeId)

		ret.SetRetCode(0)
		ret.SetAwards(awards)

		galaxy.LogDebug("all fbs:", this.Fbs)
		return ret
	}

	return ret
}

func (this *UserFbs) save(schemeId int32) bool {
	if ob, ok := (*this).Fbs[schemeId]; ok {
		buff, err := proto.Marshal(this.newFbCache(ob))
		if err != nil {
			galaxy.LogError(err)
			return false
		}
		if _, err = RedisCmd("SET", this.key(this.user.GetUid(), ob.schemeId), buff); err != nil {
			galaxy.LogError(err)
			return false
		}

		if _, err := RedisCmd("SADD", this.keys(this.user.GetUid()), this.key(this.user.GetUid(), ob.schemeId)); err != nil {
			galaxy.LogError(err)
			return false
		}

		return true
	}
	return false
}

func (this *UserFbs) getCurrentTimeInterval(schemeId int32) int64 {
	EveryDay := strings.Split(scheme.Fbmap[schemeId].EveryDay, ";")
	for k, v := range EveryDay {
		interval := strings.Split(v, "-")
		start := fmt.Sprintf("%d-%d-%d %s:00", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), interval[0])
		end := fmt.Sprintf("%d-%d-%d %s:00", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), interval[1])
		startTime, err := time.Parse("2006-01-02 15:04:05", start)
		if err != nil {
			galaxy.LogError(err)
			return -1
		}

		endTime, err := time.Parse("2006-01-02 15:04:05", end)
		if err != nil {
			galaxy.LogError(err)
			return -1
		}

		if Time() >= startTime.Unix() && Time() < endTime.Unix() {
			return int64(k)
		}
	}
	return -1
}

//example, return 201501010
func (this *UserFbs) getTimestamp() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	return int64(year*10000+int(month)*100+day) * 10
}
