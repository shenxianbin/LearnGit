package mall

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	. "galaxy"

	"github.com/golang/protobuf/proto"
)

func InitMallModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_MallInitReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		LogDebug("Enter1")

		ret := &protocol.MsgMallInitRet{
			Infos: role.FillMallInfo(),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		LogDebug(buf)
		global.SendMsg(int32(protocol.MsgCode_MallInitRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MallBuyReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		LogDebug("Enter2")

		req := &protocol.MsgMallBuyReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		retcode, retinfo := role.MallBuy(req.GetMallId())
		ret := &protocol.MsgMallBuyRet{
			Retcode: proto.Int32(int32(retcode)),
			Info:    retinfo,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_MallBuyRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_MallGoldFillReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgMallGoldFillReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		retcode := role.MallGoldFill(req.GetSoul())
		ret := &protocol.MsgMallGoldFillRet{
			Retcode: proto.Int32(int32(retcode)),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_MallGoldFillRet), sid, buf)
	})
}
