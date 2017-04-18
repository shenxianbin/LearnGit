package building

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitBuildingModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_BuildingStartLvUpReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter MsgCode_BuildingStartLvUpReq")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgBuildingStartLvUpReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retCode := role.BuildingStartLvUp(req.GetBuildingUid(), req.GetUsedCoin(), true)
		galaxy.LogDebug("retcode:", retCode)
		ret := &protocol.MsgBuildingStartLvUpRet{}
		ret.SetRetcode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_BuildingStartLvUpRet), sid, buf)
		galaxy.LogDebug("end MsgCode_BuildingStartLvUpReq")
	})

	global.RegisterMsg(int32(protocol.MsgCode_BuildingCancelLvUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgBuildingCancelLvUpReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retCode := role.BuildingCancelLvUp(req.GetBuildingUid(), true)

		ret := &protocol.MsgBuildingCancelLvUpRet{}
		ret.SetRetcode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_BuildingCancelLvUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_BuildingFinishLvUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgBuildingFinishLvUpReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retCode := role.BuildingFinishLvUp(req.GetBuildingUid(), true)

		ret := &protocol.MsgBuildingFinishLvUpRet{}
		ret.SetRetcode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_BuildingFinishLvUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_BuildingLvUpRemoveTimeReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgBuildingLvUpRemoveTimeReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retCode := role.BuildingLvUpRemoveTime(req.GetBuildingUid(), true)

		ret := &protocol.MsgBuildingLvUpRemoveTimeRet{}
		ret.SetRetcode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_BuildingLvUpRemoveTimeRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_BuildingCollectReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgBuildingCollectReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retcode := role.BuildingCollect(req.GetUid(), true)

		ret := &protocol.MsgBuildingCollectRet{
			Retcode: proto.Int32(int32(retcode)),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_BuildingCollectRet), sid, buf)
	})
}
