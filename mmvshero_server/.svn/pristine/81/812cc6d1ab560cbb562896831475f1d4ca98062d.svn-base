package item

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	. "galaxy"

	"github.com/golang/protobuf/proto"
)

func InitItemModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_ItemUseReq), func(sid int64, msg []byte) {
		LogDebug("Enter MsgCode_ItemUseReq")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		reqMsg := &protocol.MsgItemUseReq{}
		err := proto.Unmarshal(msg, reqMsg)
		if err != nil {
			LogDebug(err)
			return
		}

		retcode := role.ItemUse(reqMsg.GetInfos(), reqMsg.GetUserType(), reqMsg.GetUserId())
		LogDebug("retcode:", retcode)
		retMsg := &protocol.MsgItemUseRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(retMsg)
		if err != nil {
			LogDebug(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ItemUseRet), sid, buf)
		LogDebug("End MsgCode_ItemUseReq")
	})

	global.RegisterMsg(int32(protocol.MsgCode_ItemSellReq), func(sid int64, msg []byte) {
		LogDebug("Enter MsgCode_ItemSellReq")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		reqMsg := &protocol.MsgItemSellReq{}
		err := proto.Unmarshal(msg, reqMsg)
		if err != nil {
			LogDebug(err)
			return
		}

		retcode := role.ItemSell(reqMsg.GetUid(), reqMsg.GetNum(), true)
		LogDebug("retcode:", retcode)
		retMsg := &protocol.MsgItemSellRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(retMsg)
		if err != nil {
			LogDebug(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ItemSellRet), sid, buf)
		LogDebug("End MsgCode_ItemSellRet")
	})
}
