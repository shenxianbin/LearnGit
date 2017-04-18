package tcp

import (
	"galaxy/define"
	. "galaxy/event"
	. "galaxy/nets/packet"
)

type Dispatcher struct {
	serverType  string
	msgHandler  map[int32]func(s *Session, packet *GxPacket) //注册协议
	unknowMsgCb func(s *Session, packet *GxPacket)
}

func NewDispatcher(serverType string) *Dispatcher {
	nd := new(Dispatcher)
	nd.Init(serverType)
	return nd
}

func (this *Dispatcher) Init(serverType string) {
	this.serverType = serverType
	this.msgHandler = make(map[int32]func(s *Session, packet *GxPacket))
}

func (this *Dispatcher) RegisterMsg(msgCode int32, f func(s *Session, packet *GxPacket)) {
	if _, has := this.msgHandler[msgCode]; !has {
		this.msgHandler[msgCode] = f
	}

}

func (this *Dispatcher) RegisterUnknowMsgCb(f func(s *Session, packet *GxPacket)) {
	this.unknowMsgCb = f
}

func (this *Dispatcher) Process(s *Session, packet *GxPacket) {
	if handle, has := this.msgHandler[packet.MsgCode()]; has {
		GxEvent().Execute(func(args ...interface{}) {
			handle(args[0].(*Session), args[1].(*GxPacket))
		}, s, packet)
	} else {
		if packet.MsgCode() != define.MSGCODE_HEART && this.unknowMsgCb != nil {
			this.unknowMsgCb(s, packet)
		}
	}
}
