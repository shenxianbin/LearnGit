package gm

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/gm"
	"common/protocol"
	. "galaxy"
	"galaxy/nets/packet"
	"galaxy/nets/tcp"
	"time"

	"github.com/golang/protobuf/proto"
)

var lantern_cache *protocol.MsgLanternNotify

func InitGmModule() {
	init_protocol()
}

func init_protocol() {
	GxService().RegisterMsg("CentreServer", int32(gm.GmMsgCode_GmCommandNotify), func(s *tcp.Session, p *packet.GxPacket) {
		LogDebug("Receive GmCommandNotify From Centre")
		content := p.Content()
		msg := new(gm.MsgGmCommandNotify)
		err := proto.Unmarshal(content, msg)
		if err != nil {
			LogError(err)
			return
		}

		role := logic.GetRoleByUid(msg.GetUid())
		if role == nil {
			return
		}
		role.GmProcess(true)
	})

	GxService().RegisterMsg("CentreServer", int32(gm.GmMsgCode_GmLanternNotify), func(s *tcp.Session, p *packet.GxPacket) {
		LogDebug("Receive GmLanternNotify From Centre")
		content := p.Content()
		msg := new(gm.MsgGmLanternNotify)
		err := proto.Unmarshal(content, msg)
		if err != nil {
			LogError(err)
			return
		}

		notify := new(protocol.MsgLanternNotify)
		notify.SetType(msg.GetType())
		notify.SetId(msg.GetId())
		notify.SetTime(msg.GetTime())
		notify.SetContent(msg.GetContent())

		lantern_cache = notify
		LogDebug("lantern_cache : ", lantern_cache)

		buf, err := proto.Marshal(notify)
		if err != nil {
			LogError(err)
			return
		}

		global.SendBroadCast(int32(protocol.MsgCode_LanternNotify), 0, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_LanternReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		LogDebug("MsgCode_LanternReq : Uid = ", role.GetUid())
		if lantern_cache != nil {
			LogDebug("lantern_cache : ", lantern_cache)
		} else {
			LogDebug("lantern_cache nil")
		}

		if lantern_cache != nil && time.Now().Unix() <= lantern_cache.GetTime() {
			buf, err := proto.Marshal(lantern_cache)
			if err != nil {
				LogError(err)
				return
			}
			global.SendMsg(int32(protocol.MsgCode_LanternNotify), sid, buf)
			LogDebug("lantern_cache send to : Uid = ", role.GetUid())
		}
	})
}
