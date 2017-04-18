package mission

import (
	"Gameserver/global"
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	"common"
	. "common/cache"
	"common/protocol"
	"common/scheme"
	"fmt"
	"galaxy"
	"math"

	"github.com/golang/protobuf/proto"
)

type Mission struct {
	schemeId   int32
	reachedNum int32 //达成数量
	targetNum  int32 //目标数量 计算之后的值
	timestamp  int64 //任务生成的时间戳
	finished   bool  //是否完成，完成指领取奖励
	level      int32 //任务刷新时玩家等级
}

type UserMissions struct {
	user     IRole
	Missions map[int32]*Mission //key Mission schemeId
}

const (
	key  = "Role:%v:Mission:%v"
	keys = "Role:%v:Missions"
)

func (this *UserMissions) key(roleId int64, schemeId int32) string {
	return fmt.Sprintf(key, roleId, schemeId)
}

//获取一个用户的所有魔物key
func (this *UserMissions) keys(roleId int64) string {
	return fmt.Sprintf(keys, roleId)
}

func (this *UserMissions) Init(user IRole) {
	this.user = user
	this.Missions = make(map[int32]*Mission)
	this.Load()
}

func (this *UserMissions) Load() error {
	roleId := this.user.GetUid()
	if this.Missions == nil {
		this.Missions = make(map[int32]*Mission)
	}
	// galaxy.LogDebug("load this.Missions:", this.Missions, len(this.Missions))
	//获取所有keys
	resp, err := galaxy.GxService().Redis().Cmd("SMEMBERS", this.keys(roleId))
	if err != nil {
		galaxy.LogError(err)
		return err //没有
	}

	keys, _ := resp.List()
	for _, key := range keys {
		//获取数据
		resp, err := galaxy.GxService().Redis().Cmd("GET", key)
		if err != nil {
			galaxy.LogError(err)
			return err //没有
		}
		if buff, _ := resp.Bytes(); buff != nil {
			missionCache := &MissionCache{}
			err = proto.Unmarshal(buff, missionCache)
			this.Missions[missionCache.GetSchemeId()] = this.newMission(missionCache)
		} else {
			galaxy.LogError("SREM: ", this.keys(roleId), key)
			galaxy.GxService().Redis().Cmd("SREM", this.keys(roleId), key)
		}
	}

	if err != nil {
		galaxy.LogError(err)
	}

	// galaxy.LogDebug("missions:", this.Missions)

	this.refreshMissions()

	return nil
}

func (this *UserMissions) MissionAll() *protocol.MsgMissionAllRet {
	this.refreshMissions()
	// galaxy.LogDebug("this.Missions:", this.Missions)

	list1 := make([]*protocol.Mission, len(this.Missions))

	var i int32 = 0
	for _, v := range this.Missions {
		list1[i] = v.FillInfo()
		i++
	}

	m := new(protocol.MsgMissionAllRet)
	m.SetMissions(list1)
	return m
}

func (this *Mission) FillInfo() *protocol.Mission {
	// var awards []*protocol.MissionAward
	m := new(protocol.Mission)
	m.SchemeId = proto.Int32(this.schemeId)
	m.ReachedNum = proto.Int32(this.reachedNum)
	m.Timestamp = proto.Int64(this.timestamp)
	m.Finished = proto.Bool(this.finished)
	m.Level = proto.Int32(this.level)
	m.TargetNum = proto.Int32(this.targetNum)
	// if s, ok := scheme.Missionmap[this.schemeId]; ok {
	// 	temp := ExtractIntPairMap(s.AwardId)
	// 	if len(temp) > 0 {
	// 		awards = make([]*protocol.MissionAward, len(temp))
	// 		var i int32 = 0
	// 		for k, v := range ExtractIntPairMap(s.AwardId) {
	// 			awards[i].AwardId = proto.Int32(k)
	// 			awards[i].LvParam = proto.Int32(v)
	// 			i++
	// 		}
	// 		m.SetAwards(awards)
	// 	}
	// }
	return m
}

func (this *Mission) Notify(user IRole) {
	ret := &protocol.MsgMissionNotifyRet{}
	ret.SetMission(this.FillInfo())
	buf, err := proto.Marshal(ret)
	if err != nil {
		galaxy.LogError(err)
		return
	}

	global.SendMsg(int32(protocol.MsgCode_MissionNotifyRet), user.GetSid(), buf)
}

func (this *UserMissions) newMissionCache(m *Mission) *MissionCache {
	c := &MissionCache{}
	c.SchemeId = proto.Int32(m.schemeId)
	c.ReachedNum = proto.Int32(m.reachedNum)
	c.Timestamp = proto.Int64(m.timestamp)
	c.Finished = proto.Bool(m.finished)
	c.Level = proto.Int32(m.level)
	c.TargetNum = proto.Int32(m.targetNum)
	return c
}

func (this *UserMissions) newMission(cache *MissionCache) *Mission {
	s := &Mission{}
	s.schemeId = cache.GetSchemeId()
	s.reachedNum = cache.GetReachedNum()
	s.timestamp = cache.GetTimestamp()
	s.finished = cache.GetFinished()
	s.level = cache.GetLevel()
	s.targetNum = cache.GetTargetNum()
	return s
}

func (this *UserMissions) saveMissionToDb(schemeId int32) bool {
	if mission, ok := (*this).Missions[schemeId]; ok {
		missionCache := this.newMissionCache(mission)

		buff, err := proto.Marshal(missionCache)
		if err != nil {
			galaxy.LogError(err)
			return false
		}
		if _, err = galaxy.GxService().Redis().Cmd("SET", this.key(this.user.GetUid(), mission.schemeId), buff); err != nil {
			galaxy.LogError(err)
			return false
		}

		if _, err := galaxy.GxService().Redis().Cmd("SADD", this.keys(this.user.GetUid()), this.key(this.user.GetUid(), mission.schemeId)); err != nil {
			galaxy.LogError(err)
			return false
		}

		return true
	}
	return false
}

func (this UserMissions) calcTargetNum(schemeId, level int32) int32 {
	if m, ok := scheme.Missionmap[schemeId]; ok {
		TargetNum := m.TargetNum
		LvParam := m.LvParam
		if LvParam > 0 {
			TargetNum *= int32(math.Ceil(1 / float64(LvParam) * float64(level)))
		}
		return TargetNum
	}
	return 0
}

func (this UserMissions) calcAwards(schemeId, level int32) map[int32]int32 {
	if m, ok := scheme.Missionmap[schemeId]; ok {
		awards := ExtractIntPairMap(m.AwardId)
		for k, v := range awards {
			if v > 0 {
				awards[k] = int32(math.Ceil(1 / float64(v) * float64(level)))
			} else {
				awards[k] = 1
			}
		}

		return awards
	}
	return nil
}

//任务刷新 新号、
func (this *UserMissions) refreshMissions() {
	for id, v := range scheme.Missionmap {
		if mission, ok := this.Missions[id]; ok {
			// galaxy.LogDebug("refreshMissions 0")
			//检查是否过期 ，当天早上5点
			if v.Once == 0 && mission.timestamp < RefreshTime(5) {
				mission.timestamp = Time()
				mission.finished = false
				mission.reachedNum = 0
				mission.level = this.user.GetLv()
				mission.targetNum = this.calcTargetNum(id, mission.level)
				this.saveMissionToDb(id)
			}
		} else {
			//检查是否达到开放等级
			// galaxy.LogDebug("refreshMissions 1")
			if this.user.GetLv() >= v.NeedRoleLv {
				// galaxy.LogDebug("this.user.GetLv() >= v.NeedRoleLv;", v.NeedRoleLv, this.user.GetLv(), v.Id)
				m := &Mission{}
				m.level = this.user.GetLv()
				m.finished = false
				m.reachedNum = 0
				m.schemeId = v.Id
				m.timestamp = Time()
				m.targetNum = this.calcTargetNum(v.Id, m.level)
				this.Missions[m.schemeId] = m
				this.saveMissionToDb(m.schemeId)
			}
		}
	}
}

//添加达成数量
func (this *UserMissions) MissionAddNum(schemeId, num, targetLevel int32) (common.RetCode, int32) {
	this.refreshMissions()
	if mission, ok := this.Missions[schemeId]; ok {
		//判断是否符合条件
		needLv := scheme.Missionmap[schemeId].TargetLv
		if needLv > 0 && targetLevel < needLv {
			// galaxy.LogDebug("target level too small.", needLv, targetLevel)
			return common.RetCode_MissionTargetLvUnable, 0
		}

		if mission.reachedNum < mission.targetNum {
			mission.reachedNum += num
			this.saveMissionToDb(schemeId)
			mission.Notify(this.user)
			return common.RetCode_Success, mission.reachedNum
		}
	} else {
		return common.RetCode_MissionIdError, 0
	}
	return common.RetCode_Failed, 0
}

func (this *UserMissions) MissionFinish(schemeId int32) common.RetCode {
	this.refreshMissions()
	if mission, ok := this.Missions[schemeId]; ok {
		if mission.reachedNum >= mission.targetNum && mission.finished == false {
			//检查是否vip
			if this.user.IsVip() == false && scheme.Missionmap[schemeId].OnlyVip == 1 {
				return common.RetCode_MissionMustVip
			}

			//发放奖励
			awards := this.calcAwards(schemeId, mission.level)
			// galaxy.LogDebug("mission awards:", schemeId, mission.level, awards, mission.finished)

			for awardId, times := range awards {
				for i := 0; i < int(times); i++ {
					Award(awardId, this.user, true)
				}
			}

			mission.finished = true
			this.saveMissionToDb(schemeId)
			mission.Notify(this.user)

			//添加成就
			// this.user.AchievementAddNum(18, 1, false)

			return common.RetCode_Success
		}
	} else {
		return common.RetCode_MissionIdError
	}
	return common.RetCode_Failed
}
