package global

import (
	. "galaxy"
	"galaxy/nets/packet"
	"galaxy/nets/tcp"
)

func RegisterMsg(msgCode int32, f func(sid int64, msg []byte)) {
	GxService().RegisterMsg("GateServer", msgCode, func(s *tcp.Session, p *packet.GxPacket) {
		f(p.Sid(), p.Content())
	})
}

func SendMsg(msgCode int32, sid int64, msg []byte) {
	p := packet.NewPacket(sid, msgCode)
	p.SetContent(msg)
	GxService().SendToConnector("GateServer", GxService().Config().GetConnectorConfig(0).ConnectServerId(), p)
}

func SendBroadCast(msgCode int32, space_sid int64, msg []byte) {
	p := packet.NewPacket(space_sid, msgCode)
	p.SetBroadcast(true)
	p.SetContent(msg)
	GxService().SendToConnector("GateServer", GxService().Config().GetConnectorConfig(0).ConnectServerId(), p)
}

func SendBroadCastExcept(msgCode int32, space_sid int64, except_sid uint64, msg []byte) {
	p := packet.NewPacket(space_sid, msgCode)
	p.SetBroadcast(true)
	p.SetExceptSid(except_sid)
	p.SetContent(msg)
	GxService().SendToConnector("GateServer", GxService().Config().GetConnectorConfig(0).ConnectServerId(), p)
}

func GetServerId() int32 {
	return GxService().Config().GetServerId()
}

func SendToStaticServer(msgCode int32, msg []byte) {
	p := packet.NewPacket(0, msgCode)
	p.SetContent(msg)
	GxService().SendToConnector("StaticServer", GxService().Config().GetConnectorConfig(2).ConnectServerId(), p)
}
