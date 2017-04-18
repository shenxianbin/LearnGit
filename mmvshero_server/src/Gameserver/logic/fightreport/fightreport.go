package fightreport

import (
	. "Gameserver/logic"
	"common"
	. "common/cache"
	"common/protocol"
	"fmt"
	. "galaxy"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	cache_fightreport_id_list_t = "Role:%v:FightReportIdList"
	cache_fightreport_key_t     = "FightReport:%v"
	cache_fightreport_autoid_t  = "FightReport:AutoId"
)

const (
	fight_report_timeout = 3600 * 24 * 3
)

func genFightReportIdListKey(role_uid int64) string {
	return fmt.Sprintf(cache_fightreport_id_list_t, role_uid)
}

func genFightReportAutoId() int64 {
	resp, err := GxService().Redis().Cmd("INCR", cache_fightreport_autoid_t)
	if err != nil {
		return 0
	}

	uid, _ := resp.Int64()
	return uid
}

func genFightReportCacheKey(uid int64) string {
	return fmt.Sprintf(cache_fightreport_key_t, uid)
}

type FightReportSys struct {
	owner             IRole
	report_list       map[int64]*FightReportCache
	cache_list_id_key string
}

func (this *FightReportSys) Init(owner IRole) {
	this.owner = owner
	this.report_list = make(map[int64]*FightReportCache)
	this.cache_list_id_key = genFightReportIdListKey(this.owner.GetUid())
}

func (this *FightReportSys) Load() error {
	resp, err := GxService().Redis().Cmd("SMEMBERS", this.cache_list_id_key)
	if err != nil {
		return err
	}

	cacheKeys, err := resp.List()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", this.cache_list_id_key, key)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			cache := &FightReportCache{}
			err = proto.Unmarshal(buf, cache)
			if err != nil {
				LogFatal(err)
				continue
			}

			if cache.GetTimeStamp()+fight_report_timeout > now {
				GxService().Redis().Cmd("DEL", key)
				GxService().Redis().Cmd("SREM", this.cache_list_id_key, key)
				continue
			}

			this.report_list[cache.GetUid()] = cache
		}
	}

	return nil
}

func (this *FightReportSys) FightReportQuery() []*protocol.FightReportInfo {
	infos := make([]*protocol.FightReportInfo, len(this.report_list))
	index := 0
	for _, v := range this.report_list {
		report := &protocol.FightReportInfo{
			ReportUid:   proto.Int64(v.GetUid()),
			FightReport: v.GetReport(),
		}

		infos[index] = report
		index++
	}

	return infos
}

func (this *FightReportSys) FightReportQueryById(uid int64) (common.RetCode, *protocol.FightReportInfo) {
	if v, has := this.report_list[uid]; has {
		report := &protocol.FightReportInfo{
			ReportUid:   proto.Int64(v.GetUid()),
			FightReport: v.GetReport(),
		}
		return common.RetCode_Success, report
	}

	key := genFightReportCacheKey(uid)
	resp, err := GxService().Redis().Cmd("GET", key)
	if err != nil {
		return common.RetCode_Redis_Error, nil
	}

	if buf, _ := resp.Bytes(); buf != nil {
		cache := &FightReportCache{}
		err = proto.Unmarshal(buf, cache)
		if err != nil {
			LogFatal(err)
			return common.RetCode_Proto_Error, nil
		}

		if cache.GetTimeStamp()+fight_report_timeout > time.Now().Unix() {
			GxService().Redis().Cmd("DEL", key)
			return common.RetCode_TimeOut_Error, nil
		}

		report := &protocol.FightReportInfo{
			ReportUid:   proto.Int64(cache.GetUid()),
			FightReport: cache.GetReport(),
		}

		return common.RetCode_Success, report
	}

	return common.RetCode_TimeOut_Error, nil
}

func (this *FightReportSys) FightReportAdd(active_uid int64, passive_uid int64, info *protocol.FightReportInfo) {
	if active_uid != this.owner.GetUid() {
		return
	}

	if info == nil {
		return
	}

	auto_id := genFightReportAutoId()
	if auto_id == 0 {
		return
	}

	cache := &FightReportCache{
		Uid:       proto.Int64(auto_id),
		Report:    info.GetFightReport(),
		TimeStamp: proto.Int64(time.Now().Unix()),
	}

	buf, err := proto.Marshal(cache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := genFightReportCacheKey(auto_id)
	_, err = GxService().Redis().Cmd("SET", key, buf)
	if err != nil {
		LogFatal(err)
		return
	}

	_, err = GxService().Redis().Cmd("SADD", this.cache_list_id_key, key)
	if err != nil {
		LogFatal(err)
		return
	}

	if passive_uid != 0 {
		_, err = GxService().Redis().Cmd("SADD", genFightReportIdListKey(passive_uid), key)
		if err != nil {
			LogFatal(err)
			return
		}
	}

	this.report_list[auto_id] = cache
}

func (this *FightReportSys) FightReportUpdate(report_uid int64, info *protocol.FightReportInfo) {
	_, has := this.report_list[report_uid]
	if !has {
		return
	}

	cache := &FightReportCache{
		Uid:       proto.Int64(report_uid),
		Report:    info.GetFightReport(),
		TimeStamp: proto.Int64(time.Now().Unix()),
	}

	buf, err := proto.Marshal(cache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := genFightReportCacheKey(report_uid)
	_, err = GxService().Redis().Cmd("SET", key, buf)
	if err != nil {
		LogFatal(err)
		return
	}

	this.report_list[report_uid] = cache
}
