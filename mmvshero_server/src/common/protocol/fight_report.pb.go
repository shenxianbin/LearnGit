// Code generated by protoc-gen-go.
// source: fight_report.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MsgFightReportReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgFightReportReq) Reset()         { *m = MsgFightReportReq{} }
func (m *MsgFightReportReq) String() string { return proto.CompactTextString(m) }
func (*MsgFightReportReq) ProtoMessage()    {}

type FightReportInfo struct {
	ReportUid        *int64 `protobuf:"varint,1,req,name=report_uid" json:"report_uid,omitempty"`
	FightReport      []byte `protobuf:"bytes,2,req,name=fightReport" json:"fightReport,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FightReportInfo) Reset()         { *m = FightReportInfo{} }
func (m *FightReportInfo) String() string { return proto.CompactTextString(m) }
func (*FightReportInfo) ProtoMessage()    {}

func (m *FightReportInfo) GetReportUid() int64 {
	if m != nil && m.ReportUid != nil {
		return *m.ReportUid
	}
	return 0
}

func (m *FightReportInfo) GetFightReport() []byte {
	if m != nil {
		return m.FightReport
	}
	return nil
}

func (m *FightReportInfo) SetReportUid(value int64) {
	if m != nil {
		if m.ReportUid != nil {
			*m.ReportUid = value
			return
		}
		m.ReportUid = proto.Int64(value)
	}
}

func (m *FightReportInfo) SetFightReport(value []byte) {
	if m != nil {
		if m.FightReport != nil {
			m.FightReport = value
			return
		}
		m.FightReport = value
	}
}

type MsgFightReportRet struct {
	Infos            []*FightReportInfo `protobuf:"bytes,1,rep,name=infos" json:"infos,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *MsgFightReportRet) Reset()         { *m = MsgFightReportRet{} }
func (m *MsgFightReportRet) String() string { return proto.CompactTextString(m) }
func (*MsgFightReportRet) ProtoMessage()    {}

func (m *MsgFightReportRet) GetInfos() []*FightReportInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (m *MsgFightReportRet) SetInfos(value []*FightReportInfo) {
	if m != nil {
		m.Infos = value
	}
}

type MsgFightReportIdReq struct {
	ReportUid        *int64 `protobuf:"varint,1,req,name=report_uid" json:"report_uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgFightReportIdReq) Reset()         { *m = MsgFightReportIdReq{} }
func (m *MsgFightReportIdReq) String() string { return proto.CompactTextString(m) }
func (*MsgFightReportIdReq) ProtoMessage()    {}

func (m *MsgFightReportIdReq) GetReportUid() int64 {
	if m != nil && m.ReportUid != nil {
		return *m.ReportUid
	}
	return 0
}

func (m *MsgFightReportIdReq) SetReportUid(value int64) {
	if m != nil {
		if m.ReportUid != nil {
			*m.ReportUid = value
			return
		}
		m.ReportUid = proto.Int64(value)
	}
}

type MsgFightReportIdRet struct {
	Retcode          *int32           `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Info             *FightReportInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MsgFightReportIdRet) Reset()         { *m = MsgFightReportIdRet{} }
func (m *MsgFightReportIdRet) String() string { return proto.CompactTextString(m) }
func (*MsgFightReportIdRet) ProtoMessage()    {}

func (m *MsgFightReportIdRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgFightReportIdRet) GetInfo() *FightReportInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgFightReportIdRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgFightReportIdRet) SetInfo(value *FightReportInfo) {
	if m != nil {
		m.Info = value
	}
}

type MsgFightReportAdd struct {
	ActiveUid        *int64           `protobuf:"varint,1,req,name=active_uid" json:"active_uid,omitempty"`
	PassiveUid       *int64           `protobuf:"varint,2,req,name=passive_uid" json:"passive_uid,omitempty"`
	Info             *FightReportInfo `protobuf:"bytes,3,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MsgFightReportAdd) Reset()         { *m = MsgFightReportAdd{} }
func (m *MsgFightReportAdd) String() string { return proto.CompactTextString(m) }
func (*MsgFightReportAdd) ProtoMessage()    {}

func (m *MsgFightReportAdd) GetActiveUid() int64 {
	if m != nil && m.ActiveUid != nil {
		return *m.ActiveUid
	}
	return 0
}

func (m *MsgFightReportAdd) GetPassiveUid() int64 {
	if m != nil && m.PassiveUid != nil {
		return *m.PassiveUid
	}
	return 0
}

func (m *MsgFightReportAdd) GetInfo() *FightReportInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgFightReportAdd) SetActiveUid(value int64) {
	if m != nil {
		if m.ActiveUid != nil {
			*m.ActiveUid = value
			return
		}
		m.ActiveUid = proto.Int64(value)
	}
}

func (m *MsgFightReportAdd) SetPassiveUid(value int64) {
	if m != nil {
		if m.PassiveUid != nil {
			*m.PassiveUid = value
			return
		}
		m.PassiveUid = proto.Int64(value)
	}
}

func (m *MsgFightReportAdd) SetInfo(value *FightReportInfo) {
	if m != nil {
		m.Info = value
	}
}

type MsgFightReportUpdate struct {
	ReportUid        *int64           `protobuf:"varint,1,req,name=report_uid" json:"report_uid,omitempty"`
	Info             *FightReportInfo `protobuf:"bytes,2,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MsgFightReportUpdate) Reset()         { *m = MsgFightReportUpdate{} }
func (m *MsgFightReportUpdate) String() string { return proto.CompactTextString(m) }
func (*MsgFightReportUpdate) ProtoMessage()    {}

func (m *MsgFightReportUpdate) GetReportUid() int64 {
	if m != nil && m.ReportUid != nil {
		return *m.ReportUid
	}
	return 0
}

func (m *MsgFightReportUpdate) GetInfo() *FightReportInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgFightReportUpdate) SetReportUid(value int64) {
	if m != nil {
		if m.ReportUid != nil {
			*m.ReportUid = value
			return
		}
		m.ReportUid = proto.Int64(value)
	}
}

func (m *MsgFightReportUpdate) SetInfo(value *FightReportInfo) {
	if m != nil {
		m.Info = value
	}
}

func init() {
}
