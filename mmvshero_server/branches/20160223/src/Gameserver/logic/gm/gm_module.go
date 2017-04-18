package gm

import (
	"Gameserver/logic"
	"common/gm"
	. "galaxy"
	"galaxy/nets/packet"
	"galaxy/nets/tcp"

	"github.com/golang/protobuf/proto"
)

func InitGmModule() {
	init_protocol()
}

func init_protocol() {
	GxService().RegisterMsg("CentreServer", int32(gm.GmMsgCode_GmCommandNotify), func(s *tcp.Session, p *packet.GxPacket) {
		LogDebug("Receive Msg From Centre")
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
}
