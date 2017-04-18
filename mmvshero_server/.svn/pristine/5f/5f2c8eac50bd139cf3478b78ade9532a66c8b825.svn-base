package static

import (
	. "galaxy"
	"galaxy/define"
	"galaxy/nets/packet"
	"galaxy/nets/tcp"
)

func Init() {
	GxService().RegisterMsg("GameServer", define.MSGCODE_HELLO_REQ, func(s *tcp.Session, p *packet.GxPacket) {
		s.SetLegal()
	})

	init_protocol()
}

func RegisterMsg(msgCode int32, f func(msg []byte)) {
	GxService().RegisterMsg("GameServer", msgCode, func(s *tcp.Session, p *packet.GxPacket) {
		f(p.Content())
	})
}
