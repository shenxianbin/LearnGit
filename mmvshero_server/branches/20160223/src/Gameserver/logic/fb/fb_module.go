package fb

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitFbModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_FbAllReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_FbAllReq -------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFbAllReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := role.FbAll()
		galaxy.LogDebug("FbALl: ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		galaxy.LogDebug("end MsgCode_FbAllReq -------")
		global.SendMsg(int32(protocol.MsgCode_FbAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FbBeginReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter FbBeginReq----------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFbBeginReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		galaxy.LogDebug("req:", req.GetSchemeId())

		retCode := 0
		if !role.FbBegin(req.GetSchemeId(), req.GetDifficulty()) {
			retCode = 1
		}
		galaxy.LogDebug("retCode:", retCode)
		ret := &protocol.MsgFbBeginRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		galaxy.LogDebug("end FbBeginReq----------")
		global.SendMsg(int32(protocol.MsgCode_FbBeginRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FbFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFbFinishReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := role.FbFinish(req.GetSchemeId(), req.GetDifficulty(), req.GetIsPassed(), req.GetCaughtNpc())

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FbFinishRet), sid, buf)
	})

}