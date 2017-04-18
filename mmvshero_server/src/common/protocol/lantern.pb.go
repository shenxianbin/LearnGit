// Code generated by protoc-gen-go.
// source: lantern.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MsgLanternNotify struct {
	Type             *int32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Id               *int32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Time             *int64 `protobuf:"varint,3,req,name=time" json:"time,omitempty"`
	Content          []byte `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgLanternNotify) Reset()         { *m = MsgLanternNotify{} }
func (m *MsgLanternNotify) String() string { return proto.CompactTextString(m) }
func (*MsgLanternNotify) ProtoMessage()    {}

func (m *MsgLanternNotify) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *MsgLanternNotify) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *MsgLanternNotify) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *MsgLanternNotify) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MsgLanternNotify) SetType(value int32) {
	if m != nil {
		if m.Type != nil {
			*m.Type = value
			return
		}
		m.Type = proto.Int32(value)
	}
}

func (m *MsgLanternNotify) SetId(value int32) {
	if m != nil {
		if m.Id != nil {
			*m.Id = value
			return
		}
		m.Id = proto.Int32(value)
	}
}

func (m *MsgLanternNotify) SetTime(value int64) {
	if m != nil {
		if m.Time != nil {
			*m.Time = value
			return
		}
		m.Time = proto.Int64(value)
	}
}

func (m *MsgLanternNotify) SetContent(value []byte) {
	if m != nil {
		if m.Content != nil {
			m.Content = value
			return
		}
		m.Content = value
	}
}

type MsgLanternReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgLanternReq) Reset()         { *m = MsgLanternReq{} }
func (m *MsgLanternReq) String() string { return proto.CompactTextString(m) }
func (*MsgLanternReq) ProtoMessage()    {}

func init() {
}
