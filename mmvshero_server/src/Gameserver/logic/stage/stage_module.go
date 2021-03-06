package stage

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"
	"github.com/golang/protobuf/proto"
)

func InitStageModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_StageAllReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgStageAllReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.StageAll())
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_StageAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_StagePlayAnimationReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgStagePlayAnimationReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := &protocol.MsgStagePlayAnimationRet{}
		ret.SetRetCode(int32(role.StagePlayAnimation(req.GetSchemeId())))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_StagePlayAnimationRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_StageBeginReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgStageBeginReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := &protocol.MsgStageBeginRet{}
		ret.SetRetCode(int32(role.StageBegin(req.GetSchemeId())))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_StageBeginRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_StageFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgStageFinishReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		stars := make(map[string]int32)
		for _, v := range req.GetStars() {
			stars[v.GetMissionId()] = v.GetIsFinish()
		}

		ret := role.StageFinish(req.GetSchemeId(), req.GetIsPassed(), stars, req.GetIsSweep(), req.GetSweepTimes())

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_StageFinishRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_StagePurchaseReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgStagePurchaseReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.StagePurchase(req.GetSchemeId()))
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_StagePurchaseRet), sid, buf)
	})
}
