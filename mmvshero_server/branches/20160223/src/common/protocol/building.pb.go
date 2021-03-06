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
	CollectTimestamp *int64 `protobuf:"varint,4,req,name=collect_timestamp" json:"collect_timestamp,omitempty"`
	BuildTimestamp   *int64 `protobuf:"varint,5,req,name=build_timestamp" json:"build_timestamp,omitempty"`
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

func (m *BuildingInfo) GetCollectTimestamp() int64 {
	if m != nil && m.CollectTimestamp != nil {
		return *m.CollectTimestamp
	}
	return 0
}

func (m *BuildingInfo) GetBuildTimestamp() int64 {
	if m != nil && m.BuildTimestamp != nil {
		return *m.BuildTimestamp
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

func (m *BuildingInfo) SetCollectTimestamp(value int64) {
	if m != nil {
		if m.CollectTimestamp != nil {
			*m.CollectTimestamp = value
			return
		}
		m.CollectTimestamp = proto.Int64(value)
	}
}

func (m *BuildingInfo) SetBuildTimestamp(value int64) {
	if m != nil {
		if m.BuildTimestamp != nil {
			*m.BuildTimestamp = value
			return
		}
		m.BuildTimestamp = proto.Int64(value)
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
type MsgBuildingStartLvUpReq struct {
	BuildingUid      *int64 `protobuf:"varint,1,req,name=building_uid" json:"building_uid,omitempty"`
	UsedCoin         *bool  `protobuf:"varint,2,req,name=used_coin" json:"used_coin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingStartLvUpReq) Reset()         { *m = MsgBuildingStartLvUpReq{} }
func (m *MsgBuildingStartLvUpReq) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingStartLvUpReq) ProtoMessage()    {}

func (m *MsgBuildingStartLvUpReq) GetBuildingUid() int64 {
	if m != nil && m.BuildingUid != nil {
		return *m.BuildingUid
	}
	return 0
}

func (m *MsgBuildingStartLvUpReq) GetUsedCoin() bool {
	if m != nil && m.UsedCoin != nil {
		return *m.UsedCoin
	}
	return false
}

func (m *MsgBuildingStartLvUpReq) SetBuildingUid(value int64) {
	if m != nil {
		if m.BuildingUid != nil {
			*m.BuildingUid = value
			return
		}
		m.BuildingUid = proto.Int64(value)
	}
}

func (m *MsgBuildingStartLvUpReq) SetUsedCoin(value bool) {
	if m != nil {
		if m.UsedCoin != nil {
			*m.UsedCoin = value
			return
		}
		m.UsedCoin = proto.Bool(value)
	}
}

type MsgBuildingStartLvUpRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingStartLvUpRet) Reset()         { *m = MsgBuildingStartLvUpRet{} }
func (m *MsgBuildingStartLvUpRet) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingStartLvUpRet) ProtoMessage()    {}

func (m *MsgBuildingStartLvUpRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgBuildingStartLvUpRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

// 建筑升级取消
type MsgBuildingCancelLvUpReq struct {
	BuildingUid      *int64 `protobuf:"varint,1,req,name=building_uid" json:"building_uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingCancelLvUpReq) Reset()         { *m = MsgBuildingCancelLvUpReq{} }
func (m *MsgBuildingCancelLvUpReq) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingCancelLvUpReq) ProtoMessage()    {}

func (m *MsgBuildingCancelLvUpReq) GetBuildingUid() int64 {
	if m != nil && m.BuildingUid != nil {
		return *m.BuildingUid
	}
	return 0
}

func (m *MsgBuildingCancelLvUpReq) SetBuildingUid(value int64) {
	if m != nil {
		if m.BuildingUid != nil {
			*m.BuildingUid = value
			return
		}
		m.BuildingUid = proto.Int64(value)
	}
}

type MsgBuildingCancelLvUpRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingCancelLvUpRet) Reset()         { *m = MsgBuildingCancelLvUpRet{} }
func (m *MsgBuildingCancelLvUpRet) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingCancelLvUpRet) ProtoMessage()    {}

func (m *MsgBuildingCancelLvUpRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgBuildingCancelLvUpRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

// 检查建筑升级是否完成
type MsgBuildingFinishLvUpReq struct {
	BuildingUid      *int64 `protobuf:"varint,1,req,name=building_uid" json:"building_uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingFinishLvUpReq) Reset()         { *m = MsgBuildingFinishLvUpReq{} }
func (m *MsgBuildingFinishLvUpReq) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingFinishLvUpReq) ProtoMessage()    {}

func (m *MsgBuildingFinishLvUpReq) GetBuildingUid() int64 {
	if m != nil && m.BuildingUid != nil {
		return *m.BuildingUid
	}
	return 0
}

func (m *MsgBuildingFinishLvUpReq) SetBuildingUid(value int64) {
	if m != nil {
		if m.BuildingUid != nil {
			*m.BuildingUid = value
			return
		}
		m.BuildingUid = proto.Int64(value)
	}
}

type MsgBuildingFinishLvUpRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingFinishLvUpRet) Reset()         { *m = MsgBuildingFinishLvUpRet{} }
func (m *MsgBuildingFinishLvUpRet) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingFinishLvUpRet) ProtoMessage()    {}

func (m *MsgBuildingFinishLvUpRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgBuildingFinishLvUpRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

// 使用货币减少升级时间
type MsgBuildingLvUpRemoveTimeReq struct {
	BuildingUid      *int64 `protobuf:"varint,1,req,name=building_uid" json:"building_uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingLvUpRemoveTimeReq) Reset()         { *m = MsgBuildingLvUpRemoveTimeReq{} }
func (m *MsgBuildingLvUpRemoveTimeReq) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingLvUpRemoveTimeReq) ProtoMessage()    {}

func (m *MsgBuildingLvUpRemoveTimeReq) GetBuildingUid() int64 {
	if m != nil && m.BuildingUid != nil {
		return *m.BuildingUid
	}
	return 0
}

func (m *MsgBuildingLvUpRemoveTimeReq) SetBuildingUid(value int64) {
	if m != nil {
		if m.BuildingUid != nil {
			*m.BuildingUid = value
			return
		}
		m.BuildingUid = proto.Int64(value)
	}
}

type MsgBuildingLvUpRemoveTimeRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingLvUpRemoveTimeRet) Reset()         { *m = MsgBuildingLvUpRemoveTimeRet{} }
func (m *MsgBuildingLvUpRemoveTimeRet) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingLvUpRemoveTimeRet) ProtoMessage()    {}

func (m *MsgBuildingLvUpRemoveTimeRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgBuildingLvUpRemoveTimeRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

type MsgBuildingCollectReq struct {
	Uid              *int64 `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingCollectReq) Reset()         { *m = MsgBuildingCollectReq{} }
func (m *MsgBuildingCollectReq) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingCollectReq) ProtoMessage()    {}

func (m *MsgBuildingCollectReq) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MsgBuildingCollectReq) SetUid(value int64) {
	if m != nil {
		if m.Uid != nil {
			*m.Uid = value
			return
		}
		m.Uid = proto.Int64(value)
	}
}

type MsgBuildingCollectRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgBuildingCollectRet) Reset()         { *m = MsgBuildingCollectRet{} }
func (m *MsgBuildingCollectRet) String() string { return proto.CompactTextString(m) }
func (*MsgBuildingCollectRet) ProtoMessage()    {}

func (m *MsgBuildingCollectRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgBuildingCollectRet) SetRetcode(value int32) {
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
