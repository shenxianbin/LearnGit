// Code generated by protoc-gen-go.
// source: decoration.proto
// DO NOT EDIT!

package cache

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DecorationCache struct {
	SchemeId         *int32 `protobuf:"varint,1,req,name=scheme_id" json:"scheme_id,omitempty"`
	Num              *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DecorationCache) Reset()         { *m = DecorationCache{} }
func (m *DecorationCache) String() string { return proto.CompactTextString(m) }
func (*DecorationCache) ProtoMessage()    {}

func (m *DecorationCache) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *DecorationCache) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *DecorationCache) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *DecorationCache) SetNum(value int32) {
	if m != nil {
		if m.Num != nil {
			*m.Num = value
			return
		}
		m.Num = proto.Int32(value)
	}
}

func init() {
}
