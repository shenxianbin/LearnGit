package client

import (
	"galaxy/nets/packet"
	"galaxy/nets/tcp"
)

type ServerClient struct {
	session *tcp.Session
}

func (this *ServerClient) SendPacket(packet *packet.GxPacket) {
	this.session.SendPacket(packet)
}
