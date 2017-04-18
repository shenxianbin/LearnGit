package chat

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	. "galaxy"
	"galaxy/timer"

	"github.com/golang/protobuf/proto"
)

const (
	world_chat_cache = 1000
)

func InitChatModule() {
	init_protocol()

	timer.AddCdTimerEvent(30, timer.CDTimeNoCount, func() {
		_, err := GxService().Redis().Cmd("LTRIM", cache_chat_world_t, 0, world_chat_cache)
		if err != nil {
			LogError(err)
			return
		}

		//LogInfo("WorldChat Reset Success")
	})
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_ChatQueryReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgChatQueryReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		retcode, infos := role.ChatQuery(protocol.ChatType(req.GetChatType()))
		ret := &protocol.MsgChatQueryRet{
			Retcode:  proto.Int32(int32(retcode)),
			ChatType: proto.Int32(req.GetChatType()),
			Infos:    infos,
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
