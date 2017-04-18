// Code generated by protoc-gen-go.
// source: building.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type BuildingInfo struct {
	Uid              *int64 `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	SchemeId         *int32 `protobuf:"varint,2,req,name=scheme_id" json:"scheme_id,omitempty"`
	Lv               *int32 `protobuf:"varint,3,req,name=lv" json:"lv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuildingInfo) Reset()         { *m = BuildingInfo{} }
func (m *BuildingInfo) String() string { return proto.CompactTextString(m) }
func (*BuildingInfo) ProtoMessage()    {}

func (m *BuildingInfo) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *BuildingInfo) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *BuildingInfo) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *BuildingInfo) SetUid(value int64) {
	if m != nil {
		if m.Uid != nil {
			*m.Uid = value
			return
		}
		m.Uid = proto.Int64(value)
	}
}

func (m *BuildingInfo) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *BuildingInfo) SetLv(value int32) {
	if m != nil {
		if m.Lv != nil {
			*m.Lv = value
			return
		}
		m.Lv = proto.Int32(value)
	}
}

type BuildingListInfo struct {
	BuildList        []*BuildingInfo `protobuf:"bytes,1,rep,name=build_list" json:"build_list,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BuildingListInfo) Reset()         { *m = BuildingListInfo{} }
func (m *BuildingListInfo) String() string { return proto.CompactTextString(m) }
func (*BuildingListInfo) ProtoMessage()    {}

func (m *BuildingListInfo) GetBuildList() []*BuildingInfo {
	if m != nil {
		return m.BuildList
	}
	return nil
}

func (m *BuildingListInfo) SetBuildList(value []*BuildingInfo) {
	if m != nil {
		m.BuildList = value
	}
}

// 建筑變動通知
type MsgBuildingInfoNotify struct {
	Building         *BuildingInfo `protobuf:"bytes,1,req,name=building" json:"building,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *MsgBuildingInfoNotify) Reset()         { *m = MsgBuildingInfoNotify{} }
func (m *MsgBuildingInfoNotify) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingInfoNotify) ProtoMessage()    {}

func (m *MsgBuildingInfoNotify) GetBuilding() *BuildingInfo {
	if m != nil {
		return m.Building
	}
	return nil
}

func (m *MsgBuildingInfoNotify) SetBuilding(value *BuildingInfo) {
	if m != nil {
		m.Building = value
	}
}

// 建筑开始升级
type MsgBuildingLvUpReq struct {
	BuildingUid      *int64 `protobuf:"varint,1,req,name=building_uid" json:"building_uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingLvUpReq) Reset()         { *m = MsgBuildingLvUpReq{} }
func (m *MsgBuildingLvUpReq) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingLvUpReq) ProtoMessage()    {}

func (m *MsgBuildingLvUpReq) GetBuildingUid() int64 {
	if m != nil && m.BuildingUid != nil {
		return *m.BuildingUid
	}
	return 0
}

func (m *MsgBuildingLvUpReq) SetBuildingUid(value int64) {
	if m != nil {
		if m.BuildingUid != nil {
			*m.BuildingUid = value
			return
		}
		m.BuildingUid = proto.Int64(value)
	}
}

type MsgBuildingLvUpRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingLvUpRet) Reset()         { *m = MsgBuildingLvUpRet{} }
func (m *MsgBuildingLvUpRet) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingLvUpRet) ProtoMessage()    {}

func (m *MsgBuildingLvUpRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgBuildingLvUpRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func init() {
}
