package gamemap

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitMapModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_MapRefreshReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMapRefreshReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retcode := role.MapReFresh(req.GetMapInfos(), req.GetMapPointActive())
		ret := &protocol.MsgMapRefreshRet{
			Retcode: proto.Int32(int32(retcode)),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_MapRefreshRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MapInfoReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMapInfoReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retcode, info := role.MapInfo(req.GetRoleUid())
		ret := &protocol.MsgMapInfoRet{
			Retcode: proto.Int32(int32(retcode)),
			Info:    info,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_MapInfoRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MapRemoveObstacleReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMapRemoveObstacleReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retcode := role.MapRemoveObstacle(req.GetSchemeId(), req.GetPosX(), req.GetPosY())
		ret := &protocol.MsgMapRemoveObstacleRet{
			Retcode: proto.Int32(int32(retcode)),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_MapRemoveObstacleRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MapUnLockPointReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMapUnLockPointReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		retcode := role.MapUnLockPoint(req.GetPointId())
		ret := &protocol.MsgMapUnLockPointRet{
			Retcode: proto.Int32(int32(retcode)),
			PointId: proto.Int32(req.GetPointId()),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_MapUnLockPointRet), sid, buf)
	})
}
