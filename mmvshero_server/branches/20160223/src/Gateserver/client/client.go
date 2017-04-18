package client

import (
	"galaxy/nets/packet"
	"galaxy/nets/tcp"
)

type Client struct {
	session *tcp.Session
}

func (this *Client) SendPacket(packet *packet.GxPacket) {
	this.session.SendPacket(packet)
}

type ServerClient struct {
	session *tcp.Session
}

func (this *ServerClient) SendPacket(packet *packet.GxPacket) {
	this.session.SendPacket(packet)
}
