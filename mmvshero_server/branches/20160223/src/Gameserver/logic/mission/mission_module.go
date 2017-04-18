package mission

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitMissionModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_MissionAllReq), func(sid int64, msg []byte) {
		// galaxy.LogDebug("enter MsgCode_MissionAllReq -------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMissionAllReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := role.MissionAll()
		// galaxy.LogDebug("MissionAll: ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		// galaxy.LogDebug("end MsgCode_MissionAllReq -------")
		global.SendMsg(int32(protocol.MsgCode_MissionAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MissionAddNumReq), func(sid int64, msg []byte) {
		// galaxy.LogDebug("enter MissionAddNumReq----------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMissionAddNumReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		galaxy.LogDebug("req:", req.GetSchemeId())
		retCode, reachedNum := role.MissionAddNum(req.GetSchemeId(), req.GetNum(), req.GetTargetLevel())

		galaxy.LogDebug("retCode:", retCode)
		ret := &protocol.MsgMissionAddNumRet{}
		ret.SetRetCode(int32(retCode))
		ret.ReachedNum = proto.Int32(reachedNum)

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		// galaxy.LogDebug("end MissionAddNumReq----------")
		global.SendMsg(int32(protocol.MsgCode_MissionAddNumRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MissionFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMissionFinishReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgMissionFinishRet{}
		retCode := role.MissionFinish(req.GetSchemeId())

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_MissionFinishRet), sid, buf)
	})

}
