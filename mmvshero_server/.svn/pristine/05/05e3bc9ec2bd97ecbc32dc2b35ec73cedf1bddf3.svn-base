package king

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitKingModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_KingSkillStartLvUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgKingSkillStartLvUpReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		retCode := role.KingSkillStartLvUp(req.GetSkillId(), req.GetUsedCoin())
		ret := &protocol.MsgKingSkillStartLvUpRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_KingSkillStartLvUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_KingSkillLvUpRemoveTimeReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgKingSkillLvUpRemoveTimeReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		retCode := role.KingSkillLvUpRemoveTime(req.GetSkillId())
		ret := &protocol.MsgKingSkillLvUpRemoveTimeRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_KingSkillLvUpRemoveTimeRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_KingSkillFinishLvUpReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_KingSkillFinishLvUpReq:-------------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgKingSkillFinishLvUpReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		retCode := role.KingSkillFinishLvUp(req.GetSkillId())
		galaxy.LogDebug("retcode:", retCode)
		ret := &protocol.MsgKingSkillFinishLvUpRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		galaxy.LogDebug("end MsgCode_KingSkillFinishLvUpReq:-------------")

		global.SendMsg(int32(protocol.MsgCode_KingSkillFinishLvUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_KingSkillCancelLvUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgKingSkillCancelLvUpReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		retCode := role.KingSkillCancelLvUp(req.GetSkillId())
		ret := &protocol.MsgKingSkillCancelLvUpRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_KingSkillCancelLvUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_KingAddLvReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter KingAddLvReq")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgKingAddLvReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		retCode := role.KingAddLv()
		ret := &protocol.MsgKingAddLvRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		galaxy.LogDebug("end KingAddLvReq")
		global.SendMsg(int32(protocol.MsgCode_KingAddLvRet), sid, buf)
	})
}
