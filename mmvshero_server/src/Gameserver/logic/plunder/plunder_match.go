package plunder

import (
	. "Gameserver/logic"
	. "Gameserver/logic/achievement"
	. "Gameserver/logic/award"
	"common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	attackTimesForTeamKey_ = "PlunderTimes:%v:%v:%v" //%v attack id, defender id, team id, value: attack times
)

func attackTimesForTeamKey(attackerId, roleUid, teamId int64) string {
	return fmt.Sprintf(attackTimesForTeamKey_, attackerId, roleUid, teamId)
}

func (this *PlunderSys) EnrollPool(roleUid, teamId, totalExp, endTime int64) {
	plunderPool.add(roleUid, teamId, totalExp, endTime)
}

func (this *PlunderSys) RemoveFromPool(roleUid, teamId int64) {
	plunderPool.remove(roleUid, teamId)
}

func (this *PlunderSys) searchTroop() (*PlunderTeam, common.RetCode) {
	if retcode := this.searchCheck(); retcode != common.RetCode_Success {
		return nil, retcode
	}

	totalExp := scheme.RoleLvUpmap[this.owner.GetLv()].TotalExp + this.owner.GetExp()
	LogDebug("totalExp:", totalExp)

	this.owner.CostSoul(scheme.RoleLvUpmap[this.owner.GetLv()].PlunderSearchNeedSoul, true, true)
	if node := plunderPool.getByScoreRange(int64(totalExp), this.owner.GetUid()); node != nil {
		team, err := this.loadTeam(node.GetRoleUid(), node.GetTeamId())
		if err != nil {
			return nil, common.RetCode_Redis_Error
		}
		return team, common.RetCode_Success
	}

	//机器人
	team_robot := this.genRobot(true)
	return team_robot, common.RetCode_Success
}

func (this *PlunderSys) RevengeSearchLock(roleUid int64) (*PlunderTeam, common.RetCode) {
	if retcode := this.searchCheck(); retcode != common.RetCode_Success {
		return nil, retcode
	}

	teams, err := this.LoadAllTeamByUid(roleUid)
	if err != nil {
		LogDebug("err:", err)
		return nil, common.RetCode_Redis_Error
	}

	LogDebug("teams:", teams)

	var maxAwardTeamId int64
	var maxAward int32
	if len(teams) > 0 {
		for i := 0; i < len(teams); i++ {
			maxAwardTeamId = 0
			maxAward = 0
			for teamId, team := range teams {
				award := team.GetLeastAward()
				if award[0].GetAmount() > maxAward {
					maxAward = award[0].GetAmount()
					maxAwardTeamId = teamId
				}
			}
			if teams[maxAwardTeamId].IsSheild() == false {
				ret := plunderPool.addSearchLock(roleUid, maxAwardTeamId, this.owner.GetUid())
				if ret {
					this.owner.CostSoul(scheme.RoleLvUpmap[this.owner.GetLv()].PlunderSearchNeedSoul, true, true)
					return teams[maxAwardTeamId], common.RetCode_Success
				}
			}

			delete(teams, maxAwardTeamId)
		}
	}

	return nil, common.RetCode_PlunderRevengeInvalid
}

func (this *PlunderSys) ChangeTroop(selectedRoleUid, selectedTeamId int64) (*PlunderTeam, common.RetCode) {
	this.CancelMatch(selectedRoleUid, selectedTeamId, false)
	return this.searchTroop()
}

func (this *PlunderSys) searchCheck() common.RetCode {
	this.checkAttackTimesForDaily()
	if this.GetAttackTimesForDaily() >= scheme.RoleLvUpmap[this.owner.GetLv()].PlunderDailyTimes {
		LogDebug("this.GetAttackTimesForDaily():", this.GetAttackTimesForDaily())
		return common.RetCode_PlunderReachMaxAttackTimes
	}

	if this.owner.IsEnoughSoul(scheme.RoleLvUpmap[this.owner.GetLv()].PlunderSearchNeedSoul) == false {
		LogDebug("RoleNotEnoughSoul")
		return common.RetCode_RoleNotEnoughSoul
	}
	return common.RetCode_Success
}

func (this *PlunderSys) CancelMatch(selectedRoleUid, selectedTeamId int64, addPostWarLock bool) bool {
	this.resetCurrentMatch()
	this.resetAttackTimesForTeam(selectedRoleUid, selectedTeamId)
	if selectedRoleUid == -1 {
		return true
	}

	return plunderPool.cancelMatch(selectedRoleUid, selectedTeamId, this.owner.GetUid(), addPostWarLock)
}

func (this *PlunderSys) BattleLock(selectedRoleUid, selectedTeamId int64, isRevenge bool) common.RetCode {
	this.checkAttackTimesForDaily()
	if this.GetAttackTimesForDaily() >= scheme.RoleLvUpmap[this.owner.GetLv()].PlunderDailyTimes {
		LogDebug("this.GetAttackTimesForDaily():", this.GetAttackTimesForDaily())
		return common.RetCode_PlunderReachMaxAttackTimes
	}
	battleLockTime := scheme.Commonmap[define.PlunderBattleLockTime].Value

	//robot
	if selectedRoleUid == -1 {
		this.incrAttackTimesForDaily()
		this.setCurrentMatch(selectedRoleUid, selectedTeamId, int64(battleLockTime), isRevenge)
		RedisCmd("SETEX", attackTimesForTeamKey(this.owner.GetUid(), selectedRoleUid, selectedTeamId), battleLockTime, 0)
		return common.RetCode_Success
	}

	if plunderPool.addBattleLock(selectedRoleUid, selectedTeamId, this.owner.GetUid()) == false {
		return common.RetCode_Failed
	}

	if isRevenge == false {
		this.incrAttackTimesForDaily()
	}

	this.setCurrentMatch(selectedRoleUid, selectedTeamId, int64(battleLockTime), isRevenge)
	RedisCmd("SETEX", attackTimesForTeamKey(this.owner.GetUid(), selectedRoleUid, selectedTeamId), battleLockTime, 0)

	return common.RetCode_Success
}

func (this *PlunderSys) loadTeam(roleUid, teamId int64) (*PlunderTeam, error) {
	return LoadPlunderTeam(genPlunderTeamCacheKey(roleUid, teamId))
}

func (this *PlunderSys) LoadAllTeamByUid(roleUid int64) (map[int64]*PlunderTeam, error) {
	cache_list_key := fmt.Sprintf(cache_plunder_team_list_t, roleUid)
	resp, err := GxService().Redis().Cmd("SMEMBERS", cache_list_key)
	if err != nil {
		return nil, err
	}

	cacheKeys, err := resp.List()
	if err != nil {
		return nil, err
	}

	team_list := make(map[int64]*PlunderTeam)
	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", cache_list_key, key)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			team := new(PlunderTeam)
			err := proto.Unmarshal(buf, &team.PlunderTeamCache)
			if err != nil {
				return nil, err
			}

			err = team.LoadSheild(key)
			if err != nil {
				return nil, err
			}

			team_list[team.GetTeamId()] = team
		}
	}

	return team_list, nil
}

func (this *PlunderSys) checkCurrentMatch() {
	if this.GetCurrentMatchExpireTime() == 0 {
		return
	}

	if this.GetCurrentMatchExpireTime() < TimeNano() {
		LogDebug("this.GetCurrentMatchExpireTime() < TimeNano():", this.GetCurrentMatchExpireTime(), TimeNano())
		this.resetCurrentMatch()
	}
}

func (this *PlunderSys) resetCurrentMatch() {
	this.SetCurrentMatchExpireTime(0)
	this.SetCurrentMatchRoleUid(0)
	this.SetCurrentMatchTeamId(0)
	this.SetCurrentMatchIsRevenge(false)
	this.SetReportId(0)
	this.Save()
	LogDebug("resetCurrentMatch !")
}

func (this *PlunderSys) setCurrentMatch(roleUid, teamId, expire int64, isRevenge bool) {
	this.SetCurrentMatchExpireTime(TimeNano() + expire*time.Second.Nanoseconds())
	this.SetCurrentMatchRoleUid(roleUid)
	this.SetCurrentMatchTeamId(teamId)
	this.SetCurrentMatchIsRevenge(isRevenge)
	this.Save()
	LogDebug("setCurrentMatch:", roleUid, teamId, this.GetCurrentMatchExpireTime())
}

func (this *PlunderSys) getAttackTimesForTeam(roleUid, teamId int64) (int32, bool) {
	resp, err := RedisCmd("GET", attackTimesForTeamKey(this.owner.GetUid(), roleUid, teamId))
	if err != nil {
		return 0, false
	}

	times, err := resp.Int()
	if err != nil {
		return 0, false
	}

	return int32(times), true
}

func (this *PlunderSys) incrAttackTimesForTeam(roleUid, teamId int64) (int32, bool) {
	key := attackTimesForTeamKey(this.owner.GetUid(), roleUid, teamId)
	resp, err := RedisCmd("INCR", key)
	if err != nil {
		return 0, false
	}

	times, err := resp.Int()
	if err != nil {
		return 0, false
	}

	return int32(times), true
}

func (this *PlunderSys) resetAttackTimesForTeam(roleUid, teamId int64) {
	RedisCmd("DEL", attackTimesForTeamKey(this.owner.GetUid(), roleUid, teamId))
}

func (this *PlunderSys) PlunderFightStart(selectedRoleUid, selectedTeamId int64) (ret *protocol.MsgPlunderFightStartRet) {
	ret = &protocol.MsgPlunderFightStartRet{}
	ret.SetRetcode(int32(common.RetCode_Failed))

	var times int32
	times, _ = this.getAttackTimesForTeam(selectedRoleUid, selectedTeamId)
	ret.SetAttackTimesForTeam(times)

	if times < scheme.Commonmap[define.PlunderChanceFree].Value {
		//free

	} else if times < scheme.Commonmap[define.PlunderChanceFree].Value+scheme.Commonmap[define.PlunderChanceGold].Value {
		if this.owner.IsEnoughGold(scheme.Commonmap[define.PlunderChanceCost].Value) == false {
			return
		}

		this.owner.CostGold(scheme.Commonmap[define.PlunderChanceCost].Value, true, true)
	} else {
		return
	}

	times, _ = this.incrAttackTimesForTeam(selectedRoleUid, selectedTeamId)
	ret.SetAttackTimesForTeam(times)
	ret.SetRetcode(int32(common.RetCode_Success))
	return
}

//check...
func (this *PlunderSys) PlunderFight(selectedRoleUid, selectedTeamId int64, win bool, fight_type int32) (ret *protocol.MsgPlunderFightRet) {
	ret = &protocol.MsgPlunderFightRet{}
	ret.SetRetcode(int32(common.RetCode_Failed))

	//robot
	if selectedRoleUid == -1 {
		if win != false {
			this.robotAward(selectedTeamId)
			//添加成就
			this.owner.AchievementAddNum(13, 1, false)
		}
		ret.SetRetcode(int32(common.RetCode_Success))
		this.processReport(this.genRobot(false), Time(), win, fight_type, make([]*PlunderAwardCache, 0))

		times, wrong := this.getAttackTimesForTeam(selectedRoleUid, selectedTeamId)
		if wrong == false {
			LogError(wrong)
			return
		}

		if win != false || times == scheme.Commonmap[define.PlunderChanceFree].Value+scheme.Commonmap[define.PlunderChanceGold].Value {
			this.resetCurrentMatch()
			this.resetAttackTimesForTeam(selectedRoleUid, selectedTeamId)
		}
		return
	}

	var err error

	var team *PlunderTeam
	now := time.Now()

	if win == false {
		times, wrong := this.getAttackTimesForTeam(selectedRoleUid, selectedTeamId)
		if wrong == false {
			LogError(wrong)
			return
		}

		team, err = this.loadTeam(selectedRoleUid, selectedTeamId)
		if err != nil {
			LogDebug("err:", err)
			return
		}

		this.processReport(team, now.Unix(), win, fight_type, make([]*PlunderAwardCache, 0))

		if times == scheme.Commonmap[define.PlunderChanceFree].Value+scheme.Commonmap[define.PlunderChanceGold].Value {
			plunderPool.addPostWarLock(selectedRoleUid, selectedTeamId, this.owner.GetUid(), win)
			this.CancelMatch(selectedRoleUid, selectedTeamId, false)
		}

		ret.SetRetcode(int32(common.RetCode_Success))

		//添加成就
		AchievementAddNumByUid(selectedRoleUid, 12, 1, false)
		return
	}

	//begin
	ttl := plunderPool.getBattleLock(selectedRoleUid, selectedTeamId)
	if ttl <= 0 {
		return
	}

	if team.IsSheild() == false {
		key := fmt.Sprintf("%v:%v", selectedRoleUid, selectedTeamId)
		allLockers[key].Lock()
		defer allLockers[key].Unlock()
	}

	team, err = this.loadTeam(selectedRoleUid, selectedTeamId)
	if err != nil {
		LogDebug("err:", err)
		return
	}

	onceQuantity := scheme.PlunderAwardmap[team.GetPlunderId()].OnceQuantity
	awards := team.GetMoreAward()
	gainAwards := make([]*PlunderAwardCache, 0)
	var chanaged bool = false
	for k, v := range awards {
		if v.GetAmount() == 0 {
			continue
		}
		gainAmount := v.GetAmount() * onceQuantity / 100
		if gainAmount == 0 {
			gainAmount = 1
		}
		awards[k].SetAmount(v.GetAmount() - gainAmount)

		gainAward := new(PlunderAwardCache)
		gainAward.SetType(v.GetType())
		gainAward.SetCode(v.GetCode())
		gainAward.SetAmount(gainAmount)
		gainAwards = append(gainAwards, gainAward)

		chanaged = true
	}

	if chanaged == false {
		//没抢到东西
		return
	}

	empty := true
	for _, v := range awards {
		if v.GetAmount() > 0 {
			empty = false
			break
		}
	}
	//没有战利品可以被抢
	if empty == true {
		this.RemoveFromPool(selectedRoleUid, selectedTeamId)
	}

	if team.IsSheild() == false {
		team.SetMoreAward(awards)
		this.SaveTeam(team)
	}

	this.processReport(team, now.Unix(), win, fight_type, gainAwards)

	this.plunder_award(gainAwards)
	plunderPool.addPostWarLock(selectedRoleUid, selectedTeamId, this.owner.GetUid(), win)
	ret.SetRetcode(int32(common.RetCode_Success))

	//添加成就
	this.owner.AchievementAddNum(13, 1, false)
	return
}

func (this *PlunderSys) PlunderGiveUp(selectedRoleUid, selectedTeamId int64) (ret *protocol.MsgPlunderGiveUpRet) {
	ret = &protocol.MsgPlunderGiveUpRet{}
	ret.SetRetcode(int32(common.RetCode_Failed))

	this.resetCurrentMatch()
	this.resetAttackTimesForTeam(selectedRoleUid, selectedTeamId)

	//robot
	if selectedRoleUid == -1 {
		ret.SetRetcode(int32(common.RetCode_Success))
		AchievementAddNumByUid(selectedRoleUid, 12, 1, false)
		return
	}

	if plunderPool.cancelMatch(selectedRoleUid, selectedTeamId, this.owner.GetUid(), true) {
		ret.SetRetcode(int32(common.RetCode_Success))

		//添加成就
		AchievementAddNumByUid(selectedRoleUid, 12, 1, false)
	}

	return
}

func (this *PlunderSys) checkAttackTimesForDaily() {
	if this.GetLastAttackTimeForDaily() < RefreshTime(scheme.Commonmap[define.SysResetTime].Value) {
		this.SetLastAttackTimeForDaily(Time())
		this.SetAttackTimesForDaily(0)
		this.SetPurchasedTimes(0)
	}
}

func (this *PlunderSys) incrAttackTimesForDaily() bool {
	this.checkAttackTimesForDaily()
	if this.GetAttackTimesForDaily() >= scheme.RoleLvUpmap[this.owner.GetLv()].PlunderDailyTimes {
		return false
	}
	this.SetAttackTimesForDaily(this.GetAttackTimesForDaily() + 1)
	this.Save()
	return true
}

func (this *PlunderSys) PlunderSearchQuery() (ret *protocol.MsgPlunderSearchQueryRet) {
	ret = &protocol.MsgPlunderSearchQueryRet{}
	ret.SetRetcode(int32(common.RetCode_Failed))
	this.checkAttackTimesForDaily()
	ret.SetAttackTimesForDaily(this.GetAttackTimesForDaily())
	ret.SetPurchasedTimes(this.GetPurchasedTimes())

	this.checkCurrentMatch()
	if this.GetCurrentMatchRoleUid() == 0 {
		return
	}

	var team *PlunderTeam
	var err error
	if this.GetCurrentMatchRoleUid() == -1 {
		team = this.genRobot(false)

	} else {
		roleUid, teamId := this.GetCurrentMatchRoleUid(), this.GetCurrentMatchTeamId()
		team, err = this.loadTeam(roleUid, teamId)
		if err != nil {
			LogDebug("err:", err)
			return
		}
	}

	ret.SetLocktime(int64(float64(this.GetCurrentMatchExpireTime()) * time.Nanosecond.Seconds()))

	times, _ := this.getAttackTimesForTeam(this.GetCurrentMatchRoleUid(), this.GetCurrentMatchTeamId())
	ret.SetAttackTimesForTeam(times)

	ret.SetIsRevenge(this.GetCurrentMatchIsRevenge())
	ret.SetTeam(team.FillPlunderTeamInfo())
	ret.SetRetcode(int32(common.RetCode_Success))
	return
}

func (this *PlunderSys) PlunderSearch() (ret *protocol.MsgPlunderSearchRet) {
	ret = &protocol.MsgPlunderSearchRet{}
	team, retCode := this.searchTroop()
	if retCode == common.RetCode_Success {
		ret.SetTeam(team.FillPlunderTeamInfo())
	}

	ret.SetRetcode(int32(retCode))
	return ret
}

func (this *PlunderSys) PlunderRevengeSearch(roleUid, reportId int64) (ret *protocol.MsgPlunderRevengeSearchRet) {
	ret = &protocol.MsgPlunderRevengeSearchRet{}
	ret.SetRetcode(int32(common.RetCode_Failed))

	report, err := LoadPlunderReport(genPlunderReportKey(roleUid, reportId))
	if err != nil {
		LogDebug("err:", err)
		return
	}

	if report.GetDefenceRoleUid() != this.owner.GetUid() {
		LogDebug("report.GetAttackRoleUid() != roleUid :", report.GetAttackRoleUid(), roleUid)
		return
	}

	if report.GetFightType() == int32(protocol.PlunderFightType_Revenge) {
		LogDebug("report.GetFightType() :", report.GetFightType())
		return
	}

	if report.GetIsRevenged() {
		LogDebug("err:report.GetIsRevenged()")
		ret.SetRetcode(int32(common.RetCode_PlunderRevengeDone))
		return
	}

	if Time()-report.GetRevengeCd() < int64(scheme.Commonmap[define.PlunderRevengeCD].Value) {
		LogDebug("err:RevengeCd", Time()-report.GetRevengeCd())
		ret.SetRetcode(int32(common.RetCode_PlunderRevengeCD))
		return
	}

	team, retCode := this.RevengeSearchLock(roleUid)
	if retCode != common.RetCode_Success {
		LogDebug("err:", retCode, team)
		ret.SetRetcode(int32(retCode))
		return
	}

	report.SetRevengeCd(Time())
	SavePlunderReport(report)

	ret.SetRetcode(int32(common.RetCode_Success))
	ret.SetTeam(team.FillPlunderTeamInfo())
	return
}

func (this *PlunderSys) PlunderConfirm(roleUid, teamId int64, isRevenge bool) (ret *protocol.MsgPlunderConfirmRet) {
	ret = &protocol.MsgPlunderConfirmRet{}
	ret.SetRetcode(int32(this.BattleLock(roleUid, teamId, isRevenge)))
	this.checkAttackTimesForDaily()
	ret.SetAttackTimesForDaily(this.GetAttackTimesForDaily())
	ret.SetLocktime(Time() + int64(scheme.Commonmap[define.PlunderBattleLockTime].Value))
	return
}

func (this *PlunderSys) PlunderChange(roleUid, teamId int64) (ret *protocol.MsgPlunderChangeRet) {
	ret = &protocol.MsgPlunderChangeRet{}
	team, retCode := this.ChangeTroop(roleUid, teamId)
	if retCode == common.RetCode_Success {
		ret.SetTeam(team.FillPlunderTeamInfo())
	}
	ret.SetRetcode(int32(retCode))
	return
}

func (this *PlunderSys) PlunderPurchase() (ret *protocol.MsgPlunderPurchaseRet) {
	ret = &protocol.MsgPlunderPurchaseRet{}
	ret.SetRetCode(int32(common.RetCode_Failed))
	this.checkAttackTimesForDaily()
	ret.SetPurchasedTimes(this.GetPurchasedTimes())

	if this.GetPurchasedTimes() >= scheme.Commonmap[define.PlunderBuyNum].Value {
		LogDebug("this.GetPurchasedTimes() >= scheme.Commonmap[define.PlunderBuyNum].Value:", this.GetPurchasedTimes(), scheme.Commonmap[define.PlunderBuyNum].Value)
		ret.SetRetCode(int32(common.RetCode_PlunderReachMaxPurchaseTimes))
		return
	}

	if this.GetAttackTimesForDaily() < scheme.RoleLvUpmap[this.owner.GetLv()].PlunderDailyTimes {
		LogDebug("this.GetAttackTimesForDaily() < scheme.RoleLvUpmap[this.owner.GetLv()].PlunderDailyTimes:", this.GetAttackTimesForDaily(), scheme.RoleLvUpmap[this.owner.GetLv()].PlunderDailyTimes)
		ret.SetRetCode(int32(common.RetCode_PlunderReachMaxAttackTimes))
		return
	}

	this.SetLastAttackTimeForDaily(Time())
	this.SetAttackTimesForDaily(0)
	this.SetPurchasedTimes(this.GetPurchasedTimes() + 1)

	ret.SetRetCode(int32(common.RetCode_Success))
	ret.SetPurchasedTimes(this.GetPurchasedTimes())

	ret.SetAttackTimesForDaily(this.GetAttackTimesForDaily())
	this.Save()
	return
}

func (this *PlunderSys) PlunderSearchCancel(roleUid, teamId int64) (ret *protocol.MsgPlunderSearchCancelRet) {
	ret = &protocol.MsgPlunderSearchCancelRet{}
	if this.CancelMatch(roleUid, teamId, false) == true {
		ret.SetRetcode(int32(common.RetCode_Success))
	} else {
		ret.SetRetcode(int32(common.RetCode_Failed))
	}
	return
}

func (this *PlunderSys) processReport(team *PlunderTeam, fight_time int64, win bool, fight_type int32, plunder_award []*PlunderAwardCache) {
	LogDebug("report_id : ", this.PlunderCache.GetReportId())
	if this.PlunderCache.GetReportId() == 0 {
		fight_results := make([]int32, 1)
		if win {
			fight_results[0] = 1
		} else {
			fight_results[0] = 0
		}
		report, err := NewPlunderReport(this.owner, team, fight_time, fight_results, fight_type, plunder_award)
		if err != nil {
			LogError(err)
		} else {
			if win && fight_type == int32(protocol.PlunderFightType_Revenge) {
				report.SetIsRevenged(true)
			}

			key, err := SavePlunderReport(report)
			if err != nil {
				LogError(err)
			} else {
				this.PlunderAddLog(key, team.GetStartTime()+int64(scheme.Commonmap[define.PlunderReportLife].Value))
				this.PlunderCache.SetReportId(report.GetReportId())
				this.Save()

				team.SetReportIds(append(team.GetReportIds(), key))
				this.SaveTeam(team)
				LogDebug("new report : ", report)
				LogDebug("new team : ", team)
			}
		}
	} else {
		report, err := LoadPlunderReport(genPlunderReportKey(this.owner.GetUid(), this.PlunderCache.GetReportId()))
		if err != nil {
			LogError(err)
		} else {
			if win {
				report.SetFightResult(append(report.GetFightResult(), 1))
				if fight_type == int32(protocol.PlunderFightType_Revenge) {
					report.SetIsRevenged(true)
				}
			} else {
				report.SetFightResult(append(report.GetFightResult(), 0))
				if int32(len(report.GetFightResult())) == scheme.Commonmap[define.PlunderChanceFree].Value+scheme.Commonmap[define.PlunderChanceGold].Value {
					report.SetIsRevenged(true)
				}
			}
			report.SetPlunderAward(plunder_award)
			SavePlunderReport(report)
			LogDebug("load report : ", report)
			LogDebug("load team : ", team)
		}
	}
}

func (this *PlunderSys) genRobot(is_new bool) *PlunderTeam {
	team := new(PlunderTeam)
	team.SetPos(0)

	var team_id int64
	if is_new {
		var count uint32 = 1
		var try int32
		robots := make(map[int32]struct{})
		for try < 10 {
			try++
			LogDebug("try : ", try)

			robot := scheme.ChallengeTroopRandGet(3, int32(count), this.owner.GetLv())
			if robot == nil {
				LogError("ChallengeTroopRandGet : ", count, " Lv : ", this.owner.GetLv())
				continue
			}

			id, err := strconv.Atoi(robot.NPC)
			if err != nil {
				LogError(err)
				continue
			}

			if id == 0 {
				LogDebug("Stop")
				break
			}

			npc_scheme, has := scheme.NPCmap[int32(id)]
			if !has {
				LogError("NPC ID ERROR : ", id)
				continue
			}

			if _, has := robots[npc_scheme.BaseId]; has {
				LogDebug("Already")
				continue
			}

			robots[npc_scheme.BaseId] = struct{}{}
			var temp int64 = int64(robot.Id)
			team_id = team_id | temp<<(16*(3-count))

			LogDebug("count : ", count, "team : ", team_id)

			count++
			if count > 3 {
				break
			}
		}
	} else {
		team_id = this.GetCurrentMatchTeamId()
	}

	team.SetTeamId(team_id)
	team.SetPlunderId(0)
	team.SetRoleUid(int64(protocol.PlunderRobot_Uid))
	team.SetRoleName("")
	team.SetRoleLv(0)
	team.SetHeros(make([]*PlunderHeroCache, 0))
	team.SetStartTime(time.Now().Unix())
	team.SetLeastAward(make([]*PlunderAwardCache, 0))
	team.SetMoreAward(make([]*PlunderAwardCache, 0))
	team.SetReportIds(make([]string, 0))
	team.sheild = false

	return team
}

func (this *PlunderSys) robotAward(team_id int64) {
	wave_id := team_id >> 32
	wave, has := scheme.ChallengeTroopsmap[int32(wave_id)]
	if has {
		Award(wave.PlunderAward, this.owner, true)
	} else {
		LogError("robotAward error : ", wave_id)
	}
}
