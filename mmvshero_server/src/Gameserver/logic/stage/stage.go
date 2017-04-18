package stage

import (
	"Gameserver/global"
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	. "common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"fmt"
	"galaxy"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
)

type Stage struct {
	schemeId          int32
	isBeginning       bool //是否开始
	isPassed          bool //是否通关
	isPlayedAnimation bool
	stars             map[string]int32 //任务星级
	lastBattleTime    int64            //最近战斗时间
	dailyBattleTimes  int32            //当天累计战斗次数
	purchasedTimes    int32            //当天累计购买次数
}

type UserStages struct {
	user            IRole
	Stages          map[int32]*Stage //key stage schemeId
	currentStageIds map[int32]int32  //当前进度
}

const (
	stageKey  = "Role:%v:Stage:%v"
	stageKeys = "Role:%v:Stages"
)

func key(roleId int64, schemeId int32) string {
	return fmt.Sprintf(stageKey, roleId, schemeId)
}

//获取一个用户的所有魔物key
func keys(roleId int64) string {
	return fmt.Sprintf(stageKeys, roleId)
}

func (this *UserStages) Init(user IRole) {
	this.user = user
	this.Stages = make(map[int32]*Stage)
	this.currentStageIds = make(map[int32]int32)
	this.currentStageIds[0] = 0
	this.currentStageIds[1] = 0
}

func (this *UserStages) Load() error {
	roleId := this.user.GetUid()
	if this.Stages == nil {
		this.Stages = make(map[int32]*Stage)
	}
	//获取所有keys
	resp, err := RedisCmd("SMEMBERS", keys(roleId))
	if err != nil {
		galaxy.LogError(err)
		return err //没有
	}

	_keys, _ := resp.List()
	for _, key := range _keys {
		//获取数据
		resp, err := RedisCmd("GET", key)
		if err != nil {
			galaxy.LogError(err)
			return err //没有
		}
		if buff, _ := resp.Bytes(); buff != nil {
			stageCache := &StageCache{}
			stage := &Stage{}
			err = proto.Unmarshal(buff, stageCache)
			stage.readCache(stageCache)
			this.Stages[stage.schemeId] = stage
			if stage.isPassed == false {
				if _, ok := scheme.Stagemap[stage.schemeId]; !ok {
					galaxy.LogError("Not found schemeId:", stage.schemeId)
				}
			}

			StageType := scheme.Stagemap[stage.schemeId].StageType
			if this.currentStageIds[StageType] < stage.schemeId {
				this.currentStageIds[StageType] = stage.schemeId
			}

		} else {
			RedisCmd("SREM", keys(roleId), key)
		}
	}

	if err != nil {
		galaxy.LogError(err)
	}

	return nil
}

func (this *UserStages) StageAll() *protocol.MsgStageAllRet {
	//载入第一个关
	if len(this.Stages) == 0 {
		this.Stages[1] = &Stage{schemeId: 1}
		this.currentStageIds[0] = 1
		this.Stages[1].save(this.user)
	}

	list1 := make([]*protocol.Stage, len(this.Stages))

	var i int32 = 0
	for _, v := range this.Stages {
		list1[i] = v.FillInfo()
		i++
	}

	m := new(protocol.MsgStageAllRet)
	m.SetStages(list1)

	return m
}

func (this *Stage) FillInfo() *protocol.Stage {
	s := make([]*protocol.Stars, len(this.stars))

	m := new(protocol.Stage)
	m.SetSchemeId(this.schemeId)
	m.IsBeginning = proto.Bool(this.isBeginning)
	m.IsPassed = proto.Bool(this.isPassed)
	m.IsPlayedAnimation = proto.Bool(this.isPlayedAnimation)
	m.LastBattleTime = proto.Int64(this.lastBattleTime)
	m.DailyBattleTimes = proto.Int32(this.dailyBattleTimes)
	m.PurchasedTimes = proto.Int32(this.purchasedTimes)

	var i int32 = 0
	for k, v := range this.stars {
		s[i] = &protocol.Stars{}
		s[i].SetMissionId(k)
		s[i].SetIsFinish(v)
		i++
	}

	if s != nil {
		m.SetStars(s)
	}

	return m
}

func (this *UserStages) StagePlayAnimation(schemeId int32) RetCode {
	if stage, ok := this.Stages[schemeId]; ok {
		if stage.isPlayedAnimation {
			return RetCode_Success
		}
		stage.isPlayedAnimation = true
		stage.save(this.user)
		return RetCode_Success
	}
	return RetCode_StageNotFoundError
}

func (this *Stage) addDailyBattleTimes(times int32) RetCode {
	if this.lastBattleTime < RefreshTime(scheme.Commonmap[define.SysResetTime].Value) {
		//reset daily
		this.dailyBattleTimes = 0
		this.purchasedTimes = 0
	}

	if _, has := scheme.Stagemap[this.schemeId]; !has {
		galaxy.LogDebug("Stagemap SchemeData_Error:", this.schemeId)
		return RetCode_SchemeData_Error
	}

	if scheme.Stagemap[this.schemeId].DailyTimes != -1 && this.dailyBattleTimes+times > scheme.Stagemap[this.schemeId].DailyTimes {
		//次数不够
		return RetCode_CD
	}

	this.lastBattleTime = Time()
	this.dailyBattleTimes += times
	return RetCode_Success
}

func (this *UserStages) StageBegin(schemeId int32) RetCode {
	if stage, ok := this.Stages[schemeId]; ok {
		//等级限制
		if this.user.GetLv() < scheme.Stagemap[schemeId].LevelLimit {
			return RetCode_Failed
		}

		if ret := stage.addDailyBattleTimes(1); ret != RetCode_Success {
			return ret
		}

		order := scheme.Stagemap[schemeId].LeastCostOrder
		winOrder := scheme.Stagemap[schemeId].VictoryCostOrder
		if !this.user.IsEnoughOrder(winOrder) {
			return RetCode_RoleNotEnoughOrder
		}

		this.user.CostOrder(order, true, true)
		stage.isBeginning = true
		stage.save(this.user)
		this.static_stage_log(stage, int32(static.StageStatus_begin))
		return RetCode_Success
	}

	return RetCode_StageNotFoundError
}

func (this *UserStages) StageFinish(schemeId int32, isPassed bool, stars map[string]int32, isSweep bool, sweepTimes int32) *protocol.MsgStageFinishRet {
	ret := &protocol.MsgStageFinishRet{}
	ret.SetRetCode(int32(RetCode_Failed))
	ret.CurrentStageId1 = proto.Int32(this.currentStageIds[0])
	ret.CurrentStageId2 = proto.Int32(this.currentStageIds[1])

	if stage, ok := this.Stages[schemeId]; ok {
		//奖励次数
		rewardTimes := 1
		if isSweep == true {
			rewardTimes = int(sweepTimes)
		} else {
			if stage.isBeginning == false {
				galaxy.LogDebug("stage.isBeginning == false")
				return ret
			}
		}

		//添加完成任务
		//普通
		this.user.MissionAddNum(6, int32(rewardTimes), 0)

		stage.isBeginning = false
		if isPassed == false {
			ret.SetRetCode(int32(RetCode_Success))
			stage.save(this.user)
			return ret
		}

		this.user.AddExp(scheme.Stagemap[schemeId].VictoryRoleExp*int32(rewardTimes), true, true)

		//首次通关
		if stage.isPassed == false {

			stage.isPassed = true
			//解锁掠夺特性
			id_str := strings.Split(scheme.Stagemap[schemeId].PlunderFeatures, ";")
			for _, v := range id_str {
				id, _ := strconv.Atoi(v)
				this.user.PlunderUnLockProperties(int32(id))
			}
		}

		RetExtraAwards := make([]*protocol.AwardInfo, 0)
		if isSweep == false {
			//合并星级
			if stage.stars == nil {
				stage.stars = make(map[string]int32)
			}

			for k, v := range stars {
				if v == 1 || stage.stars[k] == 1 {
					stage.stars[k] = 1
				}
			}

			//完成星级成就
			this.user.AchievementAddNum(1, this.calcAllStars(), true)
		} else {
			for _, v := range stage.stars {
				if v == 0 {
					galaxy.LogDebug("stage.stars:", stage.stars)
					return ret
				}
			}

			if len(stage.stars) < 3 {
				galaxy.LogDebug("len(stage.stars) < 3 ", stage.stars)
				return ret
			}

			//检查体力是否足够
			if false == this.user.IsEnoughOrder(scheme.Stagemap[schemeId].VictoryCostOrder*sweepTimes) {
				ret.SetRetCode(int32(RetCode_RoleNotEnoughOrder))
				return ret
			}

			//检查道具是否足够
			SweepNeedItemID := scheme.Stagemap[schemeId].SweepNeedItemID
			SweepNeedItemNum := scheme.Stagemap[schemeId].SweepNeedItemNum
			if SweepNeedItemID > 0 {
				if false == this.user.ItemIsEnough(SweepNeedItemID, SweepNeedItemNum*sweepTimes) {
					ret.SetRetCode(int32(RetCode_ItemNotEnough))
					return ret
				}
				this.user.ItemCost(SweepNeedItemID, SweepNeedItemNum*sweepTimes, true)
			}

			//检查每日可攻略次数是否足够
			if ret1 := stage.addDailyBattleTimes(sweepTimes); ret1 != RetCode_Success {
				ret.SetRetCode(int32(RetCode_CD))
				galaxy.LogDebug("addDailyBattleTimes failed :", stage.dailyBattleTimes, sweepTimes)
				return ret
			}

			// ExtraBonus := 0
			for i := 0; i < int(sweepTimes); i++ {
				temp, _ := Award(int32(scheme.Stagemap[schemeId].SweepExtraBonus), this.user, true)
				for _, awardInfo := range temp {
					RetExtraAwards = append(RetExtraAwards, awardInfo)
				}
			}
		}

		for i := 0; i < rewardTimes; i++ {
			//固定奖励
			RetFixedAwards, _ := Award(int32(scheme.Stagemap[schemeId].FixedAwardId), this.user, true)
			RetItemAwards, _ := Award(int32(scheme.Stagemap[schemeId].ItemAwardId), this.user, true)
			RetHeroAwards, _ := Award(int32(scheme.Stagemap[schemeId].HeroAwardId), this.user, true)
			RetSoldierAwards, _ := Award(int32(scheme.Stagemap[schemeId].SoldierAwardId), this.user, true)

			var countHeros int32
			for _, awardInfo := range RetHeroAwards {
				countHeros += awardInfo.GetAmount()
			}

			//完成任务
			this.user.MissionAddNum(5, countHeros, 0)
			//完成成就
			// this.user.AchievementAddNum(1, countHeros, false)

			stageAward := &protocol.StageAward{}
			stageAward.SetExtraBonus(RetExtraAwards)
			stageAward.SetFixedAwards(RetFixedAwards)
			stageAward.SetHeroAwards(RetHeroAwards)
			stageAward.SetItemAwards(RetItemAwards)
			stageAward.SetSoldierAwards(RetSoldierAwards)
			if ret.Awards == nil {
				ret.Awards = make([]*protocol.StageAward, 0)
			}
			ret.Awards = append(ret.Awards, stageAward)

			RetExtraAwards = make([]*protocol.AwardInfo, 0)
		}

		//扣除体力
		order := scheme.Stagemap[schemeId].VictoryCostOrder - scheme.Stagemap[schemeId].LeastCostOrder
		if isSweep {
			order = scheme.Stagemap[schemeId].VictoryCostOrder * int32(rewardTimes)
		}
		this.user.CostOrder(order, true, true)

		stage.save(this.user)
		this.checkNextStage(schemeId)
		ret.SetRetCode(int32(RetCode_Success))

		ret.CurrentStageId1 = proto.Int32(this.currentStageIds[0])
		ret.CurrentStageId2 = proto.Int32(this.currentStageIds[1])

		this.static_stage_log(stage, int32(static.StageStatus_end))

		return ret
	}

	galaxy.LogDebug("not found.", schemeId)
	return ret
}

//检查下一个是否开启
func (this *UserStages) checkNextStage(schemeId int32) {
	nextStageId := scheme.Stagemap[schemeId].NextStageId
	if nextStageId == -1 {
		return
	}

	//已经开启
	if _, ok := this.Stages[nextStageId]; ok {
		return
	}

	if stage, ok := this.Stages[schemeId]; ok {
		if stage.isPassed == false {
			return
		}

		//普通关卡 直接开放
		if scheme.Stagemap[schemeId].StageType == 0 {
			this.Stages[nextStageId] = &Stage{schemeId: nextStageId}
			this.currentStageIds[0] = nextStageId
			this.Stages[nextStageId].save(this.user)

			//开启精英关1
			if this.currentStageIds[1] == 0 {
				//获得精英关卡第一关
				id := scheme.StageEliteGetFirst()
				this.Stages[id] = &Stage{schemeId: id}
				this.currentStageIds[1] = id
				this.Stages[id].save(this.user)
				return
			}
		}

		//已通当前关
		//判断是否开启精英关
		if _, ok := this.Stages[this.currentStageIds[1]]; !ok {
			galaxy.LogError("not found stages:", this.currentStageIds[1])
			return
		}

		if this.Stages[this.currentStageIds[1]].isPassed == false {
			return
		}

		nextId := scheme.Stagemap[this.currentStageIds[1]].NextStageId
		if nextId == -1 {
			return
		}

		if scheme.Stagemap[this.currentStageIds[1]].StageNo < scheme.Stagemap[this.currentStageIds[0]].StageNo {
			//开放
			this.Stages[nextId] = &Stage{schemeId: nextId}
			this.currentStageIds[1] = nextId
			this.Stages[nextId].save(this.user)
		}
	}
}

func (this *UserStages) StagePurchase(schemeId int32) (ret *protocol.MsgStagePurchaseRet) {
	ret = &protocol.MsgStagePurchaseRet{}
	ret.SetRetCode(int32(RetCode_Failed))
	if stage, ok := this.Stages[schemeId]; ok {
		if stage.lastBattleTime < RefreshTime(scheme.Commonmap[define.SysResetTime].Value) {
			//reset daily
			stage.dailyBattleTimes = 0
			stage.purchasedTimes = 0
			stage.lastBattleTime = Time()
		}

		if scheme.Stagemap[schemeId].DailyTimes == -1 {
			return
		}

		if stage.dailyBattleTimes < scheme.Stagemap[schemeId].DailyTimes {
			return
		}

		if stage.purchasedTimes >= scheme.Commonmap[define.StageBuyNum].Value {
			return
		}

		price := CalcPrice(2, stage.purchasedTimes)
		if this.user.IsEnoughGold(price) == false {
			return
		}

		this.user.CostGold(price, true, true)
		ret.SetRetCode(int32(RetCode_Success))
		stage.purchasedTimes++
		stage.dailyBattleTimes = 0
		stage.save(this.user)
	}
	return
}

func (this *Stage) toCache() *StageCache {
	stars := make(map[string]int32)
	for k, v := range this.stars {
		stars[k] = v
	}
	m := &StageCache{}
	if stars != nil {
		m.SetStars(stars)
	}

	m.IsBeginning = proto.Bool(this.isBeginning)
	m.IsPassed = proto.Bool(this.isPassed)
	m.IsPlayedAnimation = proto.Bool(this.isPlayedAnimation)
	m.LastBattleTime = proto.Int64(this.lastBattleTime)
	m.DailyBattleTimes = proto.Int32(this.dailyBattleTimes)
	m.PurchasedTimes = proto.Int32(this.purchasedTimes)

	m.SetSchemeId(this.schemeId)
	return m
}

func (this *Stage) readCache(stageCache *StageCache) {
	stars := make(map[string]int32)
	this.schemeId = stageCache.GetSchemeId()
	this.isPassed = stageCache.GetIsPassed()
	this.isBeginning = stageCache.GetIsBeginning()
	this.isPlayedAnimation = stageCache.GetIsPlayedAnimation()

	this.lastBattleTime = stageCache.GetLastBattleTime()
	this.dailyBattleTimes = stageCache.GetDailyBattleTimes()

	for k, v := range stageCache.GetStars() {
		stars[k] = v
	}
	if stars != nil {
		this.stars = stars
	}
}

func (this *Stage) save(user IRole) bool {
	buff, err := proto.Marshal(this.toCache())
	if err != nil {
		galaxy.LogError(err)
		return false
	}
	if _, err = RedisCmd("SET", key(user.GetUid(), this.schemeId), buff); err != nil {
		galaxy.LogError(err)
		return false
	}

	if _, err := RedisCmd("SADD", keys(user.GetUid()), key(user.GetUid(), this.schemeId)); err != nil {
		galaxy.LogError(err)
		return false
	}

	this.notify(user)
	return true
}

func (this *Stage) notify(user IRole) {
	ret := &protocol.MsgStageNotify{}
	ret.SetStage(this.FillInfo())
	buf, err := proto.Marshal(ret)
	if err != nil {
		galaxy.LogError(err)
		return
	}

	global.SendMsg(int32(protocol.MsgCode_StageNotify), user.GetSid(), buf)
}

//统计所有星
func (this *UserStages) calcAllStars() int32 {
	var allStars int32 = 0
	for _, stage := range this.Stages {
		for _, v := range stage.stars {
			if v == 1 {
				allStars++
			}
		}
	}

	return allStars
}

func (this *UserStages) static_stage_log(stage *Stage, status int32) {
	msg := &static.MsgStaticStageLog{}
	msg.SetRoleUid(this.user.GetUid())
	msg.SetLv(this.user.GetLv())
	msg.SetSchemeId(stage.schemeId)
	msg.SetStatus(status)
	if stage.isPassed {
		msg.SetIsPassed(1)
	} else {
		msg.SetIsPassed(0)
	}
	msg.SetTimeStamp(time.Now().Unix())

	buf, err := proto.Marshal(msg)
	if err != nil {
		galaxy.LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_StageLog), buf)
}