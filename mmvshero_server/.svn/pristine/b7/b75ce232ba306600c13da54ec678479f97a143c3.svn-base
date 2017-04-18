package chat

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	. "galaxy"

	"github.com/golang/protobuf/proto"
)

func InitChatModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_ChatQueryReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode, infos := role.ChatQuery()
		ret := &protocol.MsgChatQueryRet{
			Retcode: proto.Int32(int32(retcode)),
			Infos:   infos,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ChatQueryRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ChatReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgChatReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		retcode, info := role.Chat(protocol.ChatType(req.GetChatType()), req.GetRoleUid(), req.GetContent())
		ret := &protocol.MsgChatRet{
			Retcode:  proto.Int32(int32(retcode)),
			ChatType: proto.Int32(req.GetChatType()),
			Info:     info,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ChatRet), sid, buf)
	})
}
