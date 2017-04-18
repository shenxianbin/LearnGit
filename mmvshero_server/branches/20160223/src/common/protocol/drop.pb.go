// Code generated by protoc-gen-go.
// source: drop.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MsgDropReq struct {
	Type             *int32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgDropReq) Reset()         { *m = MsgDropReq{} }
func (m *MsgDropReq) String() string { return proto.CompactTextString(m) }
func (*MsgDropReq) ProtoMessage()    {}

func (m *MsgDropReq) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *MsgDropReq) SetType(value int32) {
	if m != nil {
		if m.Type != nil {
			*m.Type = value
			return
		}
		m.Type = proto.Int32(value)
	}
}

type MsgDropRet struct {
	Retcode          *int32       `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Infos            []*AwardInfo `protobuf:"bytes,2,rep,name=infos" json:"infos,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MsgDropRet) Reset()         { *m = MsgDropRet{} }
func (m *MsgDropRet) String() string { return proto.CompactTextString(m) }
func (*MsgDropRet) ProtoMessage()    {}

func (m *MsgDropRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgDropRet) GetInfos() []*AwardInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (m *MsgDropRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgDropRet) SetInfos(value []*AwardInfo) {
	if m != nil {
		m.Infos = value
	}
}

func init() {
}
