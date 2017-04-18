package client

import (
	"common/protocol"
	"galaxy"
	"galaxy/define"
	"galaxy/nets/packet"
	"galaxy/nets/tcp"

	"github.com/golang/protobuf/proto"
)

var instance *clientM
var server_id int32

type clientM struct {
	//游戏客户端
	clientMap       map[int64]*Client
	serverClientMap map[string]map[int]*ServerClient
}

func ClientManager() *clientM {
	if instance == nil {
		instance = new(clientM)
		instance.Init()
	}
	return instance
}

func (this *clientM) Add(sid int64, client *Client) {
	if _, has := this.clientMap[sid]; has {
		galaxy.LogError("SID [ ", sid, " ] client has already exsited")
		return
	}

	this.clientMap[sid] = client
	galaxy.LogDebug("SID [ ", sid, " ] client added")
}

func (this *clientM) Remove(sid int64) {
	delete(this.clientMap, sid)
	galaxy.LogDebug("SID [ ", sid, " ] client removed")
}

func (this *clientM) RemoveAll(serverType string) {
	this.clientMap = make(map[int64]*Client)
	if _, has := this.serverClientMap[serverType]; has {
		this.serverClientMap[serverType] = make(map[int]*ServerClient)
	}
	galaxy.GxService().RemoveAllListenerSession(serverType)
	galaxy.LogDebug("all client removed")
}

func (this *clientM) SendToClient(sid int64, packet *packet.GxPacket) {
	if c, has := this.clientMap[sid]; has {
		c.SendPacket(packet)
	}
}

func (this *clientM) SendToServer(serverType string, serverId int, packet *packet.GxPacket) {
	if c, has := this.serverClientMap[serverType][serverId]; has {
		c.SendPacket(packet)
	}
}

func (this *clientM) Init() {
	this.clientMap = make(map[int64]*Client)
	this.serverClientMap = make(map[string]map[int]*ServerClient)

	galaxy.GxService().RegisterListenerLoseLogicCallBack("Client", func(sid int64) {
		this.Remove(sid)
		p := packet.NewPacket(sid, define.MSGCODE_LOSE)
		this.SendToServer("GameServer", int(server_id), p)
	})

	galaxy.GxService().RegisterListenerLoseLogicCallBack("GameServer", func(sid int64) {
		this.RemoveAll("GameServer")
	})

	galaxy.GxService().RegisterUnknowMsgCb("Client", func(s *tcp.Session, p *packet.GxPacket) {
		p.SetSid(s.Sid())
		this.SendToServer("GameServer", int(server_id), p)
	})

	galaxy.GxService().RegisterUnknowMsgCb("GameServer", func(s *tcp.Session, p *packet.GxPacket) {
		this.SendToClient(p.Sid(), p)
	})

	galaxy.GxService().RegisterMsg("Client", int32(protocol.MsgCode_LoginAuthReq), func(s *tcp.Session, p *packet.GxPacket) {
		s.SetLegal()
		msg := &protocol.MsgLoginAuthReq{}
		err := proto.Unmarshal(p.Content(), msg)
		if err != nil {
			return
		}
		galaxy.LogDebug("recv MsgCode_LoginAuthReq from", s.RemoteAddr())
		galaxy.LogDebug(msg.GetTokenKey())

		client := new(Client)
		client.session = s
		this.Add(s.Sid(), client)

		ret_msg := &protocol.MsgLoginAuthRet{
			RetCode: proto.Int32(0),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		ret_packet := packet.NewPacket(s.Sid(), int32(protocol.MsgCode_LoginAuthRet))
		ret_packet.SetContent(buf)
		s.SendPacket(ret_packet)
	})

	galaxy.GxService().RegisterMsg("GameServer", define.MSGCODE_HELLO_REQ, func(s *tcp.Session, p *packet.GxPacket) {
		s.SetLegal()
		msg := &define.MsgHelloReq{}
		err := proto.Unmarshal(p.Content(), msg)
		if err != nil {
			return
		}

		serverClient := new(ServerClient)
		serverClient.session = s
		if v, has := this.serverClientMap[*msg.ServerType]; has {
			v[int(*msg.ServerId)] = serverClient
		} else {
			this.serverClientMap[*msg.ServerType] = make(map[int]*ServerClient)
			this.serverClientMap[*msg.ServerType][int(*msg.ServerId)] = serverClient
		}

		server_id = *msg.ServerId
	})
}
