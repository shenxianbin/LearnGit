package fightreport

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitFightReportModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_FightReportReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_FightReportReq -------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		ret := &protocol.MsgFightReportRet{
			Infos: role.FightReportQuery(),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FightReportRet), sid, buf)
		galaxy.LogDebug("end MsgCode_FightReportReq -------")
	})

	global.RegisterMsg(int32(protocol.MsgCode_FightReportIdReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_FightReportIdReq----------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFightReportIdReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retcode, info := role.FightReportQueryById(req.GetReportUid())

		ret := &protocol.MsgFightReportIdRet{
			Retcode: proto.Int32(int32(retcode)),
			Info:    info,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FightReportIdRet), sid, buf)
		galaxy.LogDebug("end MsgCode_FightReportIdRet----------")
	})

	global.RegisterMsg(int32(protocol.MsgCode_FightReportAdd), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_FightReportAdd----------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFightReportAdd{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		role.FightReportAdd(req.GetActiveUid(), req.GetPassiveUid(), req.GetInfo())
		galaxy.LogDebug("end MsgCode_FightReportAdd----------")
	})

	global.RegisterMsg(int32(protocol.MsgCode_FightReportUpdate), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_FightReportUpdate----------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFightReportUpdate{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		role.FightReportUpdate(req.GetReportUid(), req.GetInfo())
		galaxy.LogDebug("end MsgCode_FightReportUpdate----------")
	})
}
