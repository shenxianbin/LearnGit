package client

import (
	. "galaxy"
	"galaxy/define"
	"galaxy/nets/packet"
	"galaxy/nets/tcp"
	"github.com/golang/protobuf/proto"
)

var instance *clientM

type clientM struct {
	serverClientMap      map[string]map[int]*ServerClient
	serverClientMapIndex map[int64]int
}

func ClientManager() *clientM {
	if instance == nil {
		instance = new(clientM)
		instance.Init()
	}
	return instance
}

func (this *clientM) SendToServer(serverType string, serverId int, packet *packet.GxPacket) {
	if c, has := this.serverClientMap[serverType][serverId]; has {
		c.SendPacket(packet)
		LogDebug("send to ", serverType, serverId)
	}
}

func (this *clientM) Init() {
	this.serverClientMap = make(map[string]map[int]*ServerClient)
	this.serverClientMapIndex = make(map[int64]int)

	GxService().RegisterListenerLoseLogicCallBack("GameServer", func(sid int64) {
		if serverId, has := this.serverClientMapIndex[sid]; has {
			delete(this.serverClientMapIndex, sid)
			if _, has := this.serverClientMap["GameServer"]; has {
				delete(this.serverClientMap["GameServer"], int(serverId))
			}
		}
	})

	GxService().RegisterMsg("GameServer", define.MSGCODE_HELLO_REQ, func(s *tcp.Session, p *packet.GxPacket) {
		s.SetLegal()
		msg := &define.MsgHelloReq{}
		err := proto.Unmarshal(p.Content(), msg)
		if err != nil {
			return
		}

		serverClient := new(ServerClient)
		serverClient.session = s
		if _, has := this.serverClientMap[*msg.ServerType]; has {
			this.serverClientMap[*msg.ServerType][int(*msg.ServerId)] = serverClient
		} else {
			this.serverClientMap[*msg.ServerType] = make(map[int]*ServerClient)
			this.serverClientMap[*msg.ServerType][int(*msg.ServerId)] = serverClient
		}

		this.serverClientMapIndex[s.Sid()] = int(msg.GetServerId())
	})
}
