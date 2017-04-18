package achievement

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitAchievementModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_AchievementAllReq), func(sid int64, msg []byte) {
		//galaxy.LogDebug("enter MsgCode_AchievementAllReq -------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgAchievementAllReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := role.AchievementAll()
		//galaxy.LogDebug("AchievementAll: ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		//galaxy.LogDebug("end MsgCode_AchievementAllReq -------")
		global.SendMsg(int32(protocol.MsgCode_AchievementAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_AchievementAddNumReq), func(sid int64, msg []byte) {
		//galaxy.LogDebug("enter AchievementAddNumReq----------")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgAchievementAddNumReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		//galaxy.LogDebug("req:", req.GetSchemeId())
		retCode, reachedNum := role.AchievementAddNum(req.GetSchemeId(), req.GetNum(), req.GetIsRepace())

		//galaxy.LogDebug("retCode:", retCode)
		ret := &protocol.MsgAchievementAddNumRet{}
		ret.SetRetCode(int32(retCode))
		ret.ReachedNum = proto.Int32(reachedNum)

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		//galaxy.LogDebug("end AchievementAddNumReq----------")
		global.SendMsg(int32(protocol.MsgCode_AchievementAddNumRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_AchievementFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgAchievementFinishReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgAchievementFinishRet{}
		retCode := role.AchievementFinish(req.GetSchemeId())

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_AchievementFinishRet), sid, buf)
	})
}
