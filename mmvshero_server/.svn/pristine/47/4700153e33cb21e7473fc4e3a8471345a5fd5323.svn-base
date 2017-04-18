package achievement

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

	"github.com/golang/protobuf/proto"
)

type Achievement struct {
	schemeId    int32
	reachedNum  int32           //达成数量
	finishLevel int32           //完成等级，完成指领取奖励
	targetNum   map[int32]int32 // key Achievement level - 1
	awardId     map[int32]int32 // key Achievement level - 1
}

type UserAchievements struct {
	user         IRole
	Achievements map[int32]*Achievement //key Mission schemeId
}

const (
	_key  = "Role:%v:Achievement:%v"
	_keys = "Role:%v:Achievements"
)

func key(roleId int64, schemeId int32) string {
	return fmt.Sprintf(_key, roleId, schemeId)
}

func keys(roleId int64) string {
	return fmt.Sprintf(_keys, roleId)
}

func (this *UserAchievements) Init(user IRole) {
	this.user = user
	this.Achievements = make(map[int32]*Achievement)
}

func (this *UserAchievements) Load() error {
	roleId := this.user.GetUid()
	if this.Achievements == nil {
		this.Achievements = make(map[int32]*Achievement)
	}
	//获取所有keys
	resp, err := RedisCmd("SMEMBERS", keys(roleId))
	if err != nil {
		galaxy.LogError(err)
		return err //没有
	}

	keyList, _ := resp.List()
	for _, key := range keyList {
		//获取数据
		resp, err := RedisCmd("GET", key)
		if err != nil {
			galaxy.LogError(err)
			return nil
		}

		if buff, _ := resp.Bytes(); buff != nil {
			achievement := &Achievement{}
			cache := &AchievementCache{}
			err = proto.Unmarshal(buff, cache)
			achievement.readCache(cache)
			this.Achievements[cache.GetSchemeId()] = achievement
		} else {
			galaxy.LogError("SREM: ", keys(roleId), key)
			RedisCmd("SREM", keys(roleId), key)
		}
	}

	if err != nil {
		galaxy.LogError(err)
	}

	this.LoadFirst()
	return nil
}

func (this *UserAchievements) LoadFirst() {
	if len(this.Achievements) > 0 {
		return
	}

	for id, v := range scheme.Achievementmap {
		achievement := &Achievement{}
		achievement.schemeId = id
		achievement.finishLevel = 0
		achievement.reachedNum = 0
		achievement.awardId = ExtractIntMap(v.AwardId)
		achievement.targetNum = ExtractIntMap(v.TargetNum)

		this.Achievements[v.Id] = achievement
		achievement.save(this.user.GetUid())
	}
}

func (this *UserAchievements) AchievementAll() *protocol.MsgAchievementAllRet {
	this.LoadFirst()

	list1 := make([]*protocol.Achievement, len(this.Achievements))

	var i int32 = 0
	for _, v := range this.Achievements {
		list1[i] = v.FillInfo()
		i++
	}

	m := new(protocol.MsgAchievementAllRet)
	m.SetAchievements(list1)
	return m
}

func (this *Achievement) FillInfo() *protocol.Achievement {
	m := new(protocol.Achievement)
	m.SchemeId = proto.Int32(this.schemeId)
	m.ReachedNum = proto.Int32(this.reachedNum)
	m.FinishLevel = proto.Int32(this.finishLevel)
	m.RealFinishLevel = proto.Int32(this.getRealFinishLevel())
	return m
}

func (this *Achievement) toCache() *AchievementCache {
	c := &AchievementCache{}
	c.SchemeId = proto.Int32(this.schemeId)
	c.ReachedNum = proto.Int32(this.reachedNum)
	c.FinishLevel = proto.Int32(this.finishLevel)
	c.AwardId = make(map[int32]int32)
	for k, v := range this.awardId {
		c.AwardId[k] = v
	}
	c.TargetNum = make(map[int32]int32)
	for k, v := range this.targetNum {
		c.TargetNum[k] = v
	}

	return c
}

func (this *Achievement) readCache(cache *AchievementCache) {
	this.schemeId = cache.GetSchemeId()
	this.reachedNum = cache.GetReachedNum()
	this.finishLevel = cache.GetFinishLevel()
	this.awardId = make(map[int32]int32)
	for k, v := range cache.GetAwardId() {
		this.awardId[k] = v
	}

	this.targetNum = make(map[int32]int32)
	for k, v := range cache.GetTargetNum() {
		this.targetNum[k] = v
	}
}

func (this *Achievement) save(uid int64) bool {
	buff, err := proto.Marshal(this.toCache())
	if err != nil {
		galaxy.LogError(err)
		return false
	}
	if _, err = RedisCmd("SET", key(uid, this.schemeId), buff); err != nil {
		galaxy.LogError(err)
		return false
	}

	if _, err := RedisCmd("SADD", keys(uid), key(uid, this.schemeId)); err != nil {
		galaxy.LogError(err)
		return false
	}

	return true
}

//添加达成数量
func (this *UserAchievements) AchievementAddNum(schemeId, num int32, isReplace bool) (common.RetCode, int32) {
	if achievement, ok := this.Achievements[schemeId]; ok {
		return achievement.addNum(this.user.GetUid(), schemeId, num, isReplace)
	}
	return common.RetCode_AchievementIdError, 0
}

func AchievementAddNumByUid(uid int64, schemeId, num int32, isReplace bool) (common.RetCode, int32) {
	//load
	resp, err := RedisCmd("GET", key(uid, schemeId))
	if err != nil {
		galaxy.LogError(err)
		return common.RetCode_AchievementIdError, 0
	}

	if buff, _ := resp.Bytes(); buff != nil {
		achievement := &Achievement{}
		cache := &AchievementCache{}
		err = proto.Unmarshal(buff, cache)
		achievement.readCache(cache)

		return achievement.addNum(uid, schemeId, num, isReplace)
	}

	return common.RetCode_Redis_Error, 0
}

func (this *Achievement) addNum(uid int64, schemeId, num int32, isReplace bool) (common.RetCode, int32) {
	//成就达到最高级后不再累加
	if this.reachedNum < this.targetNum[int32(len(this.targetNum)-1)] {
		if isReplace == true {
			if this.reachedNum < num {
				this.reachedNum = num
			} else {
				return common.RetCode_AchievementFinished, 0
			}

		} else {
			this.reachedNum += num
		}

		this.save(uid)
		//发送成就完成通知
		if role := GetRoleByUid(uid); role != nil {
			this.notify(role)
		}

		return common.RetCode_Success, this.reachedNum
	}
	return common.RetCode_AchievementUnableFinish, 0
}

func (this *Achievement) notify(user IRole) {
	ret := &protocol.MsgAchievementNotifyRet{}
	ret.SetAchievement(this.FillInfo())
	buf, err := proto.Marshal(ret)
	if err != nil {
		galaxy.LogError(err)
		return
	}

	global.SendMsg(int32(protocol.MsgCode_AchievementNotifyRet), user.GetSid(), buf)
}

func (this *Achievement) getRealFinishLevel() int32 {
	//获取可领取奖励等级
	var realFinishLevel int32 = 0
	//从0开始
	for i := 0; i < len(this.targetNum); i++ {
		if this.reachedNum >= this.targetNum[int32(i)] {
			realFinishLevel = int32(i + 1)
		} else {
			break
		}
	}
	return realFinishLevel
}

func (this *UserAchievements) AchievementFinish(schemeId int32) common.RetCode {
	if achievement, ok := this.Achievements[schemeId]; ok {
		if achievement.getRealFinishLevel() > achievement.finishLevel {
			achievement.finishLevel++
			//发放当前完成成就的奖励
			Award(achievement.awardId[achievement.finishLevel-1], this.user, true)
			achievement.save(this.user.GetUid())
			achievement.notify(this.user)
			return common.RetCode_Success
		} else {
			return common.RetCode_AchievementUnableFinish
		}
	}
	return common.RetCode_AchievementIdError
}
