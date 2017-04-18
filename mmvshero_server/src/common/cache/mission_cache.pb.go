// Code generated by protoc-gen-go.
// source: mission_cache.proto
// DO NOT EDIT!

package cache

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MissionCache struct {
	SchemeId         *int32 `protobuf:"varint,1,req,name=schemeId" json:"schemeId,omitempty"`
	ReachedNum       *int32 `protobuf:"varint,2,req,name=reachedNum" json:"reachedNum,omitempty"`
	TargetNum        *int32 `protobuf:"varint,3,req,name=targetNum" json:"targetNum,omitempty"`
	Timestamp        *int64 `protobuf:"varint,4,req,name=timestamp" json:"timestamp,omitempty"`
	Finished         *bool  `protobuf:"varint,5,req,name=finished" json:"finished,omitempty"`
	Level            *int32 `protobuf:"varint,6,req,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MissionCache) Reset()         { *m = MissionCache{} }
func (m *MissionCache) String() string { return proto.CompactTextString(m) }
func (*MissionCache) ProtoMessage()    {}

func (m *MissionCache) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *MissionCache) GetReachedNum() int32 {
	if m != nil && m.ReachedNum != nil {
		return *m.ReachedNum
	}
	return 0
}

func (m *MissionCache) GetTargetNum() int32 {
	if m != nil && m.TargetNum != nil {
		return *m.TargetNum
	}
	return 0
}

func (m *MissionCache) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *MissionCache) GetFinished() bool {
	if m != nil && m.Finished != nil {
		return *m.Finished
	}
	return false
}

func (m *MissionCache) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *MissionCache) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *MissionCache) SetReachedNum(value int32) {
	if m != nil {
		if m.ReachedNum != nil {
			*m.ReachedNum = value
			return
		}
		m.ReachedNum = proto.Int32(value)
	}
}

func (m *MissionCache) SetTargetNum(value int32) {
	if m != nil {
		if m.TargetNum != nil {
			*m.TargetNum = value
			return
		}
		m.TargetNum = proto.Int32(value)
	}
}

func (m *MissionCache) SetTimestamp(value int64) {
	if m != nil {
		if m.Timestamp != nil {
			*m.Timestamp = value
			return
		}
		m.Timestamp = proto.Int64(value)
	}
}

func (m *MissionCache) SetFinished(value bool) {
	if m != nil {
		if m.Finished != nil {
			*m.Finished = value
			return
		}
		m.Finished = proto.Bool(value)
	}
}

func (m *MissionCache) SetLevel(value int32) {
	if m != nil {
		if m.Level != nil {
			*m.Level = value
			return
		}
		m.Level = proto.Int32(value)
	}
}

func init() {
}