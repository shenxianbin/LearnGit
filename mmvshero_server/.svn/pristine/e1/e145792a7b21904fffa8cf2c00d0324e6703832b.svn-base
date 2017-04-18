package pvp

/*
1.整理离线玩家战斗数据（地图，魔使，魔物，建筑，障碍，装饰）战后 掠夺数据处理 杯数计算
2.处理玩家战斗锁定 （30秒 + 5分钟）
3.排行榜，匹配（redis有序集合 zadd等），杯数 * 100 + 联赛等级 做排序 id + name
4.匹配规则 只看杯数 830 630-1030 出现没有两边扩大 530- 1130 5次极限
*/

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	. "galaxy"
	"galaxy/timer"
	"time"

	"github.com/golang/protobuf/proto"
)

func InitPvpModule() error {
	init_protocol()
	return init_data()
}

func init_data() error {
	resp, err := GxService().Redis().Cmd("GET", cache_pvprank_flag_key)
	if err != nil {
		LogError(err)
		return err
	}

	flag, err := resp.Int()
	if err != nil && err.Error() != "could not convert to int" {
		LogError(err)
	}

	now := int(time.Now().Month())
	if flag != now {
		_, err := GxService().Redis().Cmd("ZREMRANGEBYRANK", cache_pvprank_key, 0, -1)
		if err != nil {
			LogError(err)
			return err
		}
		flag = now
		_, err = GxService().Redis().Cmd("SET", cache_pvprank_flag_key, flag)
		if err != nil {
			LogError(err)
			return err
		}
	}

	timer.AddFixedTimerEvent(timer.FIXED_TIMER_TYPE_MONTH, func() {
		_, err := GxService().Redis().Cmd("ZREMRANGEBYRANK", cache_pvprank_key, 0, -1)
		if err != nil {
			LogError(err)
			return
		}

		flag_timer := int(time.Now().Month())
		_, err = GxService().Redis().Cmd("SET", cache_pvprank_flag_key, flag_timer)
		if err != nil {
			LogError(err)
			return
		}

		LogInfo("PvpRank Reset Success Flag : ", flag_timer)
	})

	return nil
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_PvpQueryReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		rank, retcode := role.PvpQuery()
		ret := &protocol.MsgPvpQueryRet{
			Retcode: proto.Int32(int32(retcode)),
			Rank:    proto.Int32(rank),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_PvpQueryRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpRankInfoReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		rank, my_rank_info, top_rank, retcode := role.PvpRankInfo()
		ret := &protocol.MsgPvpRankInfoRet{
			Retcode: proto.Int32(int32(retcode)),
			Rank:    proto.Int32(rank),
			MyRank:  my_rank_info,
			TopRank: top_rank,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_PvpRankInfoRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpPrepareReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgPvpPrepareReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		ret := &protocol.MsgPvpPrepareRet{
			Retcode: proto.Int32(int32(role.PvpPrepare(req.GetFightHeroList()))),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_PvpPrepareRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpMatchReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgPvpMatchReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		ret := role.PvpMatch()
		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_PvpMatchRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpMatchIdReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgPvpMatchIdReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		ret := role.PvpMatchById(req.GetUid())
		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_PvpMatchRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpStartReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode := role.PvpStart()
		ret := &protocol.MsgPvpStartRet{
			Retcode: proto.Int32(int32(retcode)),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_PvpStartRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpGiveUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode, trophy, stone := role.PvpGiveUp()
		ret := &protocol.MsgPvpGiveUpRet{
			Retcode: proto.Int32(int32(retcode)),
			Trophy:  proto.Int32(trophy),
			Stone:   proto.Int32(stone),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_PvpGiveUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_PvpFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgPvpFinishReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		retcode, trophy, stone := role.PvpFinish(req)
		ret := &protocol.MsgPvpFinishRet{
			Retcode: proto.Int32(int32(retcode)),
			Trophy:  proto.Int32(trophy),
			Stone:   proto.Int32(stone),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_PvpFinishRet), sid, buf)
	})
}
