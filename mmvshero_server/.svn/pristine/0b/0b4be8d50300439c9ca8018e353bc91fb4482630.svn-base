package plunder

import (
	. "Gameserver/logic"
	"common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	TEAM_LOG_LIMIT = 3
	REPORT_LIMIT   = 10
)

const (
	cache_plunder_report_autokey_t = "Role:%v:PlunderAutoKey"
	cache_plunder_report_key_t     = "Role:%v:PlunderReport:%v"
)

func genPlunderReportKey(role_uid int64, report_id int64) string {
	return fmt.Sprintf(cache_plunder_report_key_t, role_uid, report_id)
}

func genPlunderReportAutoKey(role_uid int64) string {
	return fmt.Sprintf(cache_plunder_report_autokey_t, role_uid)
}

func NewPlunderReport(attacker IRole, team *PlunderTeam, fight_time int64, fight_result []int32, fight_type int32, plunder_award []*PlunderAwardCache) (*PlunderReportCache, error) {
	resp, err := GxService().Redis().Cmd("INCR", genPlunderReportAutoKey(attacker.GetUid()))
	if err != nil {
		return nil, err
	}

	uid, _ := resp.Int64()
	report := new(PlunderReportCache)
	report.SetReportId(uid)
	report.SetAttackRoleUid(attacker.GetUid())
	report.SetAttackRoleName(attacker.GetNickname())
	report.SetAttackRoleLv(attacker.GetLv())
	report.SetDefenceRoleUid(team.GetRoleUid())
	report.SetDefenceRoleName(team.GetRoleName())
	report.SetDefenceRoleLv(team.GetRoleLv())
	report.SetHeros(team.GetHeros())
	report.SetFightTime(fight_time)
	report.SetFightResult(fight_result)
	report.SetFightType(fight_type)
	report.SetHasShield(team.sheild)
	report.SetPlunderAward(plunder_award)
	report.SetIsRevenged(false)
	report.SetRevengeCd(0)
	report.SetExData(team.GetTeamId())

	return report, nil
}

func SavePlunderReport(report *PlunderReportCache) (string, error) {
	key := genPlunderReportKey(report.GetAttackRoleUid(), report.GetReportId())
	buf, err := proto.Marshal(report)
	if err != nil {
		return key, err
	}

	_, err = GxService().Redis().Cmd("SETEX", key, int64(scheme.Commonmap[define.PlunderReportLife].Value)-(time.Now().Unix()-report.GetFightTime()), buf)
	if err != nil {
		return key, err
	}

	return key, nil
}

func LoadPlunderReport(key string) (*PlunderReportCache, error) {
	resp, err := GxService().Redis().Cmd("GET", key)
	if err != nil {
		return nil, err
	}

	cache := new(PlunderReportCache)
	if buf, _ := resp.Bytes(); buf != nil {
		err := proto.Unmarshal(buf, cache)
		if err != nil {
			return nil, err
		}
		return cache, nil
	}

	return nil, nil
}

func FillPlunderReport(report_cache *PlunderReportCache) *protocol.PlunderReport {
	report_info := new(protocol.PlunderReport)
	report_info.SetReportId(report_cache.GetReportId())
	report_info.SetAttackRoleUid(report_cache.GetAttackRoleUid())
	report_info.SetAttackRoleLv(report_cache.GetAttackRoleLv())
	report_info.SetAttackRoleName(report_cache.GetAttackRoleName())
	report_info.SetDefenceRoleUid(report_cache.GetDefenceRoleUid())
	report_info.SetDefenceRoleLv(report_cache.GetDefenceRoleLv())
	report_info.SetDefenceRoleName(report_cache.GetDefenceRoleName())
	hero_infos := make([]*protocol.PlunderHero, 0)
	for _, hero := range report_cache.GetHeros() {
		hero_info := new(protocol.PlunderHero)
		hero_info.SetSchemeId(hero.GetSchemeId())
		hero_info.SetLv(hero.GetLv())
		hero_info.SetStage(hero.GetStage())
		hero_info.SetRank(hero.GetRank())
		hero_skill_infos := make([]*protocol.PlunderHeroSkill, 0)
		for _, hero_skill := range hero_info.GetSkillList() {
			hero_skill_info := new(protocol.PlunderHeroSkill)
			hero_skill_info.SetSkillId(hero_skill.GetSkillId())
			hero_skill_info.SetSkillLv(hero_skill.GetSkillLv())
			hero_skill_infos = append(hero_skill_infos, hero_skill_info)
		}
		hero_info.SetSkillList(hero_skill_infos)
		hero_info.SetProperties(hero.GetProperties())
		hero_infos = append(hero_infos, hero_info)
	}
	report_info.SetHeros(hero_infos)
	report_info.SetFightTime(report_cache.GetFightTime())
	report_info.SetFightResult(report_cache.GetFightResult())
	report_info.SetFightType(report_cache.GetFightType())
	report_info.SetHasShield(report_cache.GetHasShield())
	award_infos := make([]*protocol.PlunderAward, 0)
	for _, award := range report_cache.GetPlunderAward() {
		award_info := new(protocol.PlunderAward)
		award_info.SetType(award.GetType())
		award_info.SetCode(award.GetCode())
		award_info.SetAmount(award.GetAmount())
		award_infos = append(award_infos, award_info)
	}
	report_info.SetPlunderAward(award_infos)
	report_info.SetIsRevenged(report_cache.GetIsRevenged())
	report_info.SetRevengeCd(report_cache.GetRevengeCd())
	report_info.SetExData(report_cache.GetExData())

	return report_info
}

func (this *PlunderSys) PlunderAddLog(report_id string, start_time int64) {
	log := new(PlunderLogCache)
	log.SetReportId(report_id)
	log.SetStartTime(start_time)

	if len(this.PlunderCache.PlunderLogs) < REPORT_LIMIT {
		this.PlunderCache.PlunderLogs = append(this.PlunderCache.PlunderLogs, log)
	} else {
		for i := 1; i < REPORT_LIMIT; i++ {
			this.PlunderCache.PlunderLogs[i-1] = this.PlunderCache.PlunderLogs[i]
		}
		this.PlunderCache.PlunderLogs[REPORT_LIMIT-1] = log
	}
	this.Save()
}

func (this *PlunderSys) PlunderAddTeamLog(team_id int64, start_time int64, report_ids []string) {
	teamlog := new(PlunderTeamLogCache)
	teamlog.SetTeamId(team_id)
	teamlog.SetStartTime(start_time)
	teamlog.SetReportIds(report_ids)

	if len(this.PlunderCache.PlunderTeamLogs) < TEAM_LOG_LIMIT {
		this.PlunderCache.PlunderTeamLogs = append(this.PlunderCache.PlunderTeamLogs, teamlog)
	} else {
		for i := 1; i < TEAM_LOG_LIMIT; i++ {
			this.PlunderCache.PlunderTeamLogs[i-1] = this.PlunderCache.PlunderTeamLogs[i]
		}
		this.PlunderCache.PlunderTeamLogs[TEAM_LOG_LIMIT-1] = teamlog
	}
	this.Save()
}

func (this *PlunderSys) PlunderGuardNowReport(pos int32) (common.RetCode, []*protocol.PlunderReport) {
	team, has := this.team_list[pos]
	if !has {
		return common.RetCode_PlunderTeamNotExist, nil
	}

	//team日志过期
	if time.Now().Unix() > team.GetStartTime()+int64(scheme.Commonmap[define.PlunderReportLife].Value) {
		return common.RetCode_TimeOut_Error, nil
	}

	reports := make([]*protocol.PlunderReport, 0)
	LogDebug("team : ", team)
	for _, report_id := range team.GetReportIds() {
		report_cache, err := LoadPlunderReport(report_id)
		if err != nil {
			continue
		}
		report := FillPlunderReport(report_cache)
		reports = append(reports, report)
	}

	return common.RetCode_Success, reports
}

func (this *PlunderSys) PlunderGuardBeforeReport() (common.RetCode, []*protocol.PlunderTeamLog) {
	team_logs := make([]*protocol.PlunderTeamLog, 0)
	now := time.Now().Unix()
	for _, teamlog_cache := range this.PlunderCache.GetPlunderTeamLogs() {
		if now > teamlog_cache.GetStartTime()+int64(scheme.Commonmap[define.PlunderReportLife].Value) {
			continue
		}

		team_log := new(protocol.PlunderTeamLog)
		team_log.SetTeamId(teamlog_cache.GetTeamId())
		reports := make([]*protocol.PlunderReport, 0)
		for _, report_id := range teamlog_cache.GetReportIds() {
			report_cache, err := LoadPlunderReport(report_id)
			if err != nil {
				continue
			}
			report := FillPlunderReport(report_cache)
			reports = append(reports, report)
		}
		team_log.SetReports(reports)
		team_logs = append(team_logs, team_log)
	}
	return common.RetCode_Success, team_logs
}

func (this *PlunderSys) PlunderReport() (common.RetCode, []*protocol.PlunderReport) {
	reports := make([]*protocol.PlunderReport, 0)
	now := time.Now().Unix()
	for _, log := range this.PlunderCache.GetPlunderLogs() {
		if now > log.GetStartTime()+int64(scheme.Commonmap[define.PlunderReportLife].Value) {
			continue
		}

		report_cache, err := LoadPlunderReport(log.GetReportId())
		if err != nil {
			continue
		}
		report := FillPlunderReport(report_cache)
		reports = append(reports, report)
	}
	return common.RetCode_Success, reports
}
