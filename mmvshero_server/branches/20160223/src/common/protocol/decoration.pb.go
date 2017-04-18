// Code generated by protoc-gen-go.
// source: decoration.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DecorationInfo struct {
	SchemeId         *int32 `protobuf:"varint,1,req,name=scheme_id" json:"scheme_id,omitempty"`
	Num              *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DecorationInfo) Reset()         { *m = DecorationInfo{} }
func (m *DecorationInfo) String() string { return proto.CompactTextString(m) }
func (*DecorationInfo) ProtoMessage()    {}

func (m *DecorationInfo) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *DecorationInfo) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *DecorationInfo) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *DecorationInfo) SetNum(value int32) {
	if m != nil {
		if m.Num != nil {
			*m.Num = value
			return
		}
		m.Num = proto.Int32(value)
	}
}

type DecorationListInfo struct {
	DecorationList   []*DecorationInfo `protobuf:"bytes,1,rep,name=decoration_list" json:"decoration_list,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DecorationListInfo) Reset()         { *m = DecorationListInfo{} }
func (m *DecorationListInfo) String() string { return proto.CompactTextString(m) }
func (*DecorationListInfo) ProtoMessage()    {}

func (m *DecorationListInfo) GetDecorationList() []*DecorationInfo {
	if m != nil {
		return m.DecorationList
	}
	return nil
}

func (m *DecorationListInfo) SetDecorationList(value []*DecorationInfo) {
	if m != nil {
		m.DecorationList = value
	}
}

func init() {
}