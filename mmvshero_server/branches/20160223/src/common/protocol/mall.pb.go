// Code generated by protoc-gen-go.
// source: mall.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MallInfo struct {
	MallId           *int32 `protobuf:"varint,1,req,name=mall_id" json:"mall_id,omitempty"`
	Args             *int64 `protobuf:"varint,2,req,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MallInfo) Reset()         { *m = MallInfo{} }
func (m *MallInfo) String() string { return proto.CompactTextString(m) }
func (*MallInfo) ProtoMessage()    {}

func (m *MallInfo) GetMallId() int32 {
	if m != nil && m.MallId != nil {
		return *m.MallId
	}
	return 0
}

func (m *MallInfo) GetArgs() int64 {
	if m != nil && m.Args != nil {
		return *m.Args
	}
	return 0
}

func (m *MallInfo) SetMallId(value int32) {
	if m != nil {
		if m.MallId != nil {
			*m.MallId = value
			return
		}
		m.MallId = proto.Int32(value)
	}
}

func (m *MallInfo) SetArgs(value int64) {
	if m != nil {
		if m.Args != nil {
			*m.Args = value
			return
		}
		m.Args = proto.Int64(value)
	}
}

type MsgMallInitReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgMallInitReq) Reset()         { *m = MsgMallInitReq{} }
func (m *MsgMallInitReq) String() string { return proto.CompactTextString(m) }
func (*MsgMallInitReq) ProtoMessage()    {}

type MsgMallInitRet struct {
	Infos            []*MallInfo `protobuf:"bytes,1,rep,name=infos" json:"infos,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MsgMallInitRet) Reset()         { *m = MsgMallInitRet{} }
func (m *MsgMallInitRet) String() string { return proto.CompactTextString(m) }
func (*MsgMallInitRet) ProtoMessage()    {}

func (m *MsgMallInitRet) GetInfos() []*MallInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (m *MsgMallInitRet) SetInfos(value []*MallInfo) {
	if m != nil {
		m.Infos = value
	}
}

type MsgMallBuyReq struct {
	MallId           *int32 `protobuf:"varint,1,req,name=mall_id" json:"mall_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgMallBuyReq) Reset()         { *m = MsgMallBuyReq{} }
func (m *MsgMallBuyReq) String() string { return proto.CompactTextString(m) }
func (*MsgMallBuyReq) ProtoMessage()    {}

func (m *MsgMallBuyReq) GetMallId() int32 {
	if m != nil && m.MallId != nil {
		return *m.MallId
	}
	return 0
}

func (m *MsgMallBuyReq) SetMallId(value int32) {
	if m != nil {
		if m.MallId != nil {
			*m.MallId = value
			return
		}
		m.MallId = proto.Int32(value)
	}
}

type MsgMallBuyRet struct {
	Retcode          *int32    `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Info             *MallInfo `protobuf:"bytes,2,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MsgMallBuyRet) Reset()         { *m = MsgMallBuyRet{} }
func (m *MsgMallBuyRet) String() string { return proto.CompactTextString(m) }
func (*MsgMallBuyRet) ProtoMessage()    {}

func (m *MsgMallBuyRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgMallBuyRet) GetInfo() *MallInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgMallBuyRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgMallBuyRet) SetInfo(value *MallInfo) {
	if m != nil {
		m.Info = value
	}
}

func init() {
}