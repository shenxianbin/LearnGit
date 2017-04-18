// Code generated by protoc-gen-go.
// source: achievement.proto
// DO NOT EDIT!

/*
	Package protocol is a generated protocol buffer package.

	It is generated from these files:
		achievement.proto
		activity.proto
		arena.proto
		award.proto
		building.proto
		challenge.proto
		chat.proto
		drop.proto
		fight_report.proto
		friend.proto
		hero.proto
		item.proto
		lantern.proto
		login.proto
		mall.proto
		map.proto
		mission.proto
		msgcode.proto
		plunder.proto
		purchase.proto
		role.proto
		sign.proto
		soldier.proto
		stage.proto

	It has these top-level messages:
		Achievement
		MsgAchievementAllReq
		MsgAchievementAllRet
		MsgAchievementAddNumReq
		MsgAchievementAddNumRet
		MsgAchievementFinishReq
		MsgAchievementFinishRet
		MsgAchievementNotifyRet
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Achievement struct {
	SchemeId         *int32 `protobuf:"varint,1,req,name=schemeId" json:"schemeId,omitempty"`
	ReachedNum       *int32 `protobuf:"varint,2,req,name=reachedNum" json:"reachedNum,omitempty"`
	FinishLevel      *int32 `protobuf:"varint,3,req,name=finishLevel" json:"finishLevel,omitempty"`
	RealFinishLevel  *int32 `protobuf:"varint,4,req,name=realFinishLevel" json:"realFinishLevel,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Achievement) Reset()         { *m = Achievement{} }
func (m *Achievement) String() string { return proto.CompactTextString(m) }
func (*Achievement) ProtoMessage()    {}

func (m *Achievement) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *Achievement) GetReachedNum() int32 {
	if m != nil && m.ReachedNum != nil {
		return *m.ReachedNum
	}
	return 0
}

func (m *Achievement) GetFinishLevel() int32 {
	if m != nil && m.FinishLevel != nil {
		return *m.FinishLevel
	}
	return 0
}

func (m *Achievement) GetRealFinishLevel() int32 {
	if m != nil && m.RealFinishLevel != nil {
		return *m.RealFinishLevel
	}
	return 0
}

func (m *Achievement) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *Achievement) SetReachedNum(value int32) {
	if m != nil {
		if m.ReachedNum != nil {
			*m.ReachedNum = value
			return
		}
		m.ReachedNum = proto.Int32(value)
	}
}

func (m *Achievement) SetFinishLevel(value int32) {
	if m != nil {
		if m.FinishLevel != nil {
			*m.FinishLevel = value
			return
		}
		m.FinishLevel = proto.Int32(value)
	}
}

func (m *Achievement) SetRealFinishLevel(value int32) {
	if m != nil {
		if m.RealFinishLevel != nil {
			*m.RealFinishLevel = value
			return
		}
		m.RealFinishLevel = proto.Int32(value)
	}
}

type MsgAchievementAllReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgAchievementAllReq) Reset()         { *m = MsgAchievementAllReq{} }
func (m *MsgAchievementAllReq) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementAllReq) ProtoMessage()    {}

type MsgAchievementAllRet struct {
	Achievements     []*Achievement `protobuf:"bytes,1,rep,name=achievements" json:"achievements,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MsgAchievementAllRet) Reset()         { *m = MsgAchievementAllRet{} }
func (m *MsgAchievementAllRet) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementAllRet) ProtoMessage()    {}

func (m *MsgAchievementAllRet) GetAchievements() []*Achievement {
	if m != nil {
		return m.Achievements
	}
	return nil
}

func (m *MsgAchievementAllRet) SetAchievements(value []*Achievement) {
	if m != nil {
		m.Achievements = value
	}
}

// 增加成就达成的数量
type MsgAchievementAddNumReq struct {
	SchemeId         *int32 `protobuf:"varint,1,req,name=schemeId" json:"schemeId,omitempty"`
	Num              *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	IsRepace         *bool  `protobuf:"varint,3,req,name=isRepace" json:"isRepace,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgAchievementAddNumReq) Reset()         { *m = MsgAchievementAddNumReq{} }
func (m *MsgAchievementAddNumReq) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementAddNumReq) ProtoMessage()    {}

func (m *MsgAchievementAddNumReq) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *MsgAchievementAddNumReq) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *MsgAchievementAddNumReq) GetIsRepace() bool {
	if m != nil && m.IsRepace != nil {
		return *m.IsRepace
	}
	return false
}

func (m *MsgAchievementAddNumReq) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *MsgAchievementAddNumReq) SetNum(value int32) {
	if m != nil {
		if m.Num != nil {
			*m.Num = value
			return
		}
		m.Num = proto.Int32(value)
	}
}

func (m *MsgAchievementAddNumReq) SetIsRepace(value bool) {
	if m != nil {
		if m.IsRepace != nil {
			*m.IsRepace = value
			return
		}
		m.IsRepace = proto.Bool(value)
	}
}

type MsgAchievementAddNumRet struct {
	RetCode          *int32 `protobuf:"varint,1,req,name=retCode" json:"retCode,omitempty"`
	ReachedNum       *int32 `protobuf:"varint,2,req,name=reachedNum" json:"reachedNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgAchievementAddNumRet) Reset()         { *m = MsgAchievementAddNumRet{} }
func (m *MsgAchievementAddNumRet) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementAddNumRet) ProtoMessage()    {}

func (m *MsgAchievementAddNumRet) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *MsgAchievementAddNumRet) GetReachedNum() int32 {
	if m != nil && m.ReachedNum != nil {
		return *m.ReachedNum
	}
	return 0
}

func (m *MsgAchievementAddNumRet) SetRetCode(value int32) {
	if m != nil {
		if m.RetCode != nil {
			*m.RetCode = value
			return
		}
		m.RetCode = proto.Int32(value)
	}
}

func (m *MsgAchievementAddNumRet) SetReachedNum(value int32) {
	if m != nil {
		if m.ReachedNum != nil {
			*m.ReachedNum = value
			return
		}
		m.ReachedNum = proto.Int32(value)
	}
}

// 成就完成，领取奖励
type MsgAchievementFinishReq struct {
	SchemeId         *int32 `protobuf:"varint,1,req,name=schemeId" json:"schemeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgAchievementFinishReq) Reset()         { *m = MsgAchievementFinishReq{} }
func (m *MsgAchievementFinishReq) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementFinishReq) ProtoMessage()    {}

func (m *MsgAchievementFinishReq) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *MsgAchievementFinishReq) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

type MsgAchievementFinishRet struct {
	RetCode          *int32 `protobuf:"varint,1,req,name=retCode" json:"retCode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgAchievementFinishRet) Reset()         { *m = MsgAchievementFinishRet{} }
func (m *MsgAchievementFinishRet) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementFinishRet) ProtoMessage()    {}

func (m *MsgAchievementFinishRet) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *MsgAchievementFinishRet) SetRetCode(value int32) {
	if m != nil {
		if m.RetCode != nil {
			*m.RetCode = value
			return
		}
		m.RetCode = proto.Int32(value)
	}
}

// 成就达成通知
type MsgAchievementNotifyRet struct {
	Achievement      *Achievement `protobuf:"bytes,1,req,name=achievement" json:"achievement,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MsgAchievementNotifyRet) Reset()         { *m = MsgAchievementNotifyRet{} }
func (m *MsgAchievementNotifyRet) String() string { return proto.CompactTextString(m) }
func (*MsgAchievementNotifyRet) ProtoMessage()    {}

func (m *MsgAchievementNotifyRet) GetAchievement() *Achievement {
	if m != nil {
		return m.Achievement
	}
	return nil
}

func (m *MsgAchievementNotifyRet) SetAchievement(value *Achievement) {
	if m != nil {
		m.Achievement = value
	}
}

func init() {
}
