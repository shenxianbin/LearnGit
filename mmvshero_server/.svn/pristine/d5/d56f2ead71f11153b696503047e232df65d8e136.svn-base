package packet

import (
	"galaxy/utils"
)

const PACKET_HEAD_LEN = 33

type GxPacket struct {
	sid         int64
	isBroadcast bool
	exceptSid   uint64
	timestamp   int64
	msgCode     int32
	length      uint32
	content     []byte

	sendBuffer utils.BinaryBuffer // 用于发包
}

func NewRecvPacket(buffer *utils.BinaryBuffer) *GxPacket {
	p := new(GxPacket)
	p.sid = buffer.ReadInt64()
	p.isBroadcast = buffer.ReadBool()
	p.exceptSid = buffer.ReadUint64()
	p.timestamp = buffer.ReadInt64()
	p.msgCode = buffer.ReadInt32()
	p.length = buffer.ReadUint32()
	return p
}

func NewPacket(sid int64, msgCode int32) *GxPacket {
	p := new(GxPacket)
	p.sid = sid
	p.msgCode = msgCode
	return p
}

func (this *GxPacket) flush() {
	this.sendBuffer.Reset()
	this.sendBuffer.WriteInt64(this.sid)
	this.sendBuffer.WriteBool(this.isBroadcast)
	this.sendBuffer.WriteUint64(this.exceptSid)
	this.sendBuffer.WriteInt64(this.timestamp)
	this.sendBuffer.WriteInt32(this.msgCode)
	if this.content != nil {
		this.sendBuffer.WriteUint32(uint32(len(this.content)))
		this.sendBuffer.Write(this.content)
	} else {
		this.sendBuffer.WriteUint32(0)
	}
}

func (this *GxPacket) ToBytes() []byte {
	this.flush()
	return this.sendBuffer.Bytes()
}

func (this *GxPacket) Sid() int64 {
	return this.sid
}

func (this *GxPacket) IsBroadcast() bool {
	return this.isBroadcast
}

func (this *GxPacket) ExceptSid() uint64 {
	return this.exceptSid
}

func (this *GxPacket) MsgCode() int32 {
	return this.msgCode
}

func (this *GxPacket) Timestamp() int64 {
	return this.timestamp
}

func (this *GxPacket) Length() uint32 {
	return this.length
}

func (this *GxPacket) Content() []byte {
	return this.content
}

func (this *GxPacket) SetSid(sid int64) {
	this.sid = sid
}

func (this *GxPacket) SetBroadcast(isBroadcast bool) {
	this.isBroadcast = isBroadcast
}

func (this *GxPacket) SetExceptSid(exceptSid uint64) {
	this.exceptSid = exceptSid
}

func (this *GxPacket) SetTimestamp(timestamp int64) {
	this.timestamp = timestamp
}

func (this *GxPacket) SetContent(content []byte) {
	this.content = content
}
