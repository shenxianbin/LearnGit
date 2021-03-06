// Code generated by protoc-gen-go.
// source: role.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// 属性变更通知
type RoleAttrType int32

const (
	RoleAttrType_Lv              RoleAttrType = 1
	RoleAttrType_Exp             RoleAttrType = 2
	RoleAttrType_Order           RoleAttrType = 3
	RoleAttrType_Order_Timestamp RoleAttrType = 4
	RoleAttrType_Soul            RoleAttrType = 5
	RoleAttrType_Gold            RoleAttrType = 6
	RoleAttrType_KingSkillLv     RoleAttrType = 7
	RoleAttrType_Vip             RoleAttrType = 8
	RoleAttrType_Chat_FreeTime   RoleAttrType = 9
	RoleAttrType_ArenaPoint      RoleAttrType = 10
)

var RoleAttrType_name = map[int32]string{
	1:  "Lv",
	2:  "Exp",
	3:  "Order",
	4:  "Order_Timestamp",
	5:  "Soul",
	6:  "Gold",
	7:  "KingSkillLv",
	8:  "Vip",
	9:  "Chat_FreeTime",
	10: "ArenaPoint",
}
var RoleAttrType_value = map[string]int32{
	"Lv":              1,
	"Exp":             2,
	"Order":           3,
	"Order_Timestamp": 4,
	"Soul":            5,
	"Gold":            6,
	"KingSkillLv":     7,
	"Vip":             8,
	"Chat_FreeTime":   9,
	"ArenaPoint":      10,
}

func (x RoleAttrType) Enum() *RoleAttrType {
	p := new(RoleAttrType)
	*p = x
	return p
}
func (x RoleAttrType) String() string {
	return proto.EnumName(RoleAttrType_name, int32(x))
}
func (x *RoleAttrType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RoleAttrType_value, data, "RoleAttrType")
	if err != nil {
		return err
	}
	*x = RoleAttrType(value)
	return nil
}

type RoleInfo struct {
	Uid                *int64       `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Nickname           *string      `protobuf:"bytes,2,req,name=nickname" json:"nickname,omitempty"`
	Lv                 *int32       `protobuf:"varint,3,req,name=lv" json:"lv,omitempty"`
	Exp                *int32       `protobuf:"varint,4,req,name=exp" json:"exp,omitempty"`
	Order              *int32       `protobuf:"varint,5,req,name=order" json:"order,omitempty"`
	OrderTimestamp     *int64       `protobuf:"varint,6,req,name=order_timestamp" json:"order_timestamp,omitempty"`
	Soul               *int64       `protobuf:"varint,7,req,name=soul" json:"soul,omitempty"`
	Gold               *int32       `protobuf:"varint,8,req,name=gold" json:"gold,omitempty"`
	KingSkills         []*KingSkill `protobuf:"bytes,9,rep,name=king_skills" json:"king_skills,omitempty"`
	Vip                *int64       `protobuf:"varint,10,req,name=vip" json:"vip,omitempty"`
	ChatFreeTime       *int32       `protobuf:"varint,11,req,name=chat_free_time" json:"chat_free_time,omitempty"`
	NewPlayerGuideStep *int32       `protobuf:"varint,12,req,name=new_player_guide_step" json:"new_player_guide_step,omitempty"`
	Alias              *string      `protobuf:"bytes,13,req,name=alias" json:"alias,omitempty"`
	XXX_unrecognized   []byte       `json:"-"`
}

func (m *RoleInfo) Reset()         { *m = RoleInfo{} }
func (m *RoleInfo) String() string { return proto.CompactTextString(m) }
func (*RoleInfo) ProtoMessage()    {}

func (m *RoleInfo) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *RoleInfo) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *RoleInfo) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *RoleInfo) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *RoleInfo) GetOrder() int32 {
	if m != nil && m.Order != nil {
		return *m.Order
	}
	return 0
}

func (m *RoleInfo) GetOrderTimestamp() int64 {
	if m != nil && m.OrderTimestamp != nil {
		return *m.OrderTimestamp
	}
	return 0
}

func (m *RoleInfo) GetSoul() int64 {
	if m != nil && m.Soul != nil {
		return *m.Soul
	}
	return 0
}

func (m *RoleInfo) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *RoleInfo) GetKingSkills() []*KingSkill {
	if m != nil {
		return m.KingSkills
	}
	return nil
}

func (m *RoleInfo) GetVip() int64 {
	if m != nil && m.Vip != nil {
		return *m.Vip
	}
	return 0
}

func (m *RoleInfo) GetChatFreeTime() int32 {
	if m != nil && m.ChatFreeTime != nil {
		return *m.ChatFreeTime
	}
	return 0
}

func (m *RoleInfo) GetNewPlayerGuideStep() int32 {
	if m != nil && m.NewPlayerGuideStep != nil {
		return *m.NewPlayerGuideStep
	}
	return 0
}

func (m *RoleInfo) GetAlias() string {
	if m != nil && m.Alias != nil {
		return *m.Alias
	}
	return ""
}

func (m *RoleInfo) SetUid(value int64) {
	if m != nil {
		if m.Uid != nil {
			*m.Uid = value
			return
		}
		m.Uid = proto.Int64(value)
	}
}

func (m *RoleInfo) SetNickname(value string) {
	if m != nil {
		if m.Nickname != nil {
			*m.Nickname = value
			return
		}
		m.Nickname = proto.String(value)
	}
}

func (m *RoleInfo) SetLv(value int32) {
	if m != nil {
		if m.Lv != nil {
			*m.Lv = value
			return
		}
		m.Lv = proto.Int32(value)
	}
}

func (m *RoleInfo) SetExp(value int32) {
	if m != nil {
		if m.Exp != nil {
			*m.Exp = value
			return
		}
		m.Exp = proto.Int32(value)
	}
}

func (m *RoleInfo) SetOrder(value int32) {
	if m != nil {
		if m.Order != nil {
			*m.Order = value
			return
		}
		m.Order = proto.Int32(value)
	}
}

func (m *RoleInfo) SetOrderTimestamp(value int64) {
	if m != nil {
		if m.OrderTimestamp != nil {
			*m.OrderTimestamp = value
			return
		}
		m.OrderTimestamp = proto.Int64(value)
	}
}

func (m *RoleInfo) SetSoul(value int64) {
	if m != nil {
		if m.Soul != nil {
			*m.Soul = value
			return
		}
		m.Soul = proto.Int64(value)
	}
}

func (m *RoleInfo) SetGold(value int32) {
	if m != nil {
		if m.Gold != nil {
			*m.Gold = value
			return
		}
		m.Gold = proto.Int32(value)
	}
}

func (m *RoleInfo) SetKingSkills(value []*KingSkill) {
	if m != nil {
		m.KingSkills = value
	}
}

func (m *RoleInfo) SetVip(value int64) {
	if m != nil {
		if m.Vip != nil {
			*m.Vip = value
			return
		}
		m.Vip = proto.Int64(value)
	}
}

func (m *RoleInfo) SetChatFreeTime(value int32) {
	if m != nil {
		if m.ChatFreeTime != nil {
			*m.ChatFreeTime = value
			return
		}
		m.ChatFreeTime = proto.Int32(value)
	}
}

func (m *RoleInfo) SetNewPlayerGuideStep(value int32) {
	if m != nil {
		if m.NewPlayerGuideStep != nil {
			*m.NewPlayerGuideStep = value
			return
		}
		m.NewPlayerGuideStep = proto.Int32(value)
	}
}

func (m *RoleInfo) SetAlias(value string) {
	if m != nil {
		if m.Alias != nil {
			*m.Alias = value
			return
		}
		m.Alias = proto.String(value)
	}
}

type MsgRoleInfoUpdateNotify struct {
	AttrType         *RoleAttrType `protobuf:"varint,1,req,name=attr_type,enum=protocol.RoleAttrType" json:"attr_type,omitempty"`
	AttrValue        *int64        `protobuf:"varint,2,req,name=attr_value" json:"attr_value,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *MsgRoleInfoUpdateNotify) Reset()         { *m = MsgRoleInfoUpdateNotify{} }
func (m *MsgRoleInfoUpdateNotify) String() string { return proto.CompactTextString(m) }
func (*MsgRoleInfoUpdateNotify) ProtoMessage()    {}

func (m *MsgRoleInfoUpdateNotify) GetAttrType() RoleAttrType {
	if m != nil && m.AttrType != nil {
		return *m.AttrType
	}
	return RoleAttrType_Lv
}

func (m *MsgRoleInfoUpdateNotify) GetAttrValue() int64 {
	if m != nil && m.AttrValue != nil {
		return *m.AttrValue
	}
	return 0
}

func (m *MsgRoleInfoUpdateNotify) SetAttrType(value RoleAttrType) {
	if m != nil {
		if m.AttrType != nil {
			*m.AttrType = value
			return
		}
	}
}

func (m *MsgRoleInfoUpdateNotify) SetAttrValue(value int64) {
	if m != nil {
		if m.AttrValue != nil {
			*m.AttrValue = value
			return
		}
		m.AttrValue = proto.Int64(value)
	}
}

// 设置昵称
type MsgRoleSetNicknameReq struct {
	Nickname         *string `protobuf:"bytes,1,req,name=nickname" json:"nickname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgRoleSetNicknameReq) Reset()         { *m = MsgRoleSetNicknameReq{} }
func (m *MsgRoleSetNicknameReq) String() string { return proto.CompactTextString(m) }
func (*MsgRoleSetNicknameReq) ProtoMessage()    {}

func (m *MsgRoleSetNicknameReq) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *MsgRoleSetNicknameReq) SetNickname(value string) {
	if m != nil {
		if m.Nickname != nil {
			*m.Nickname = value
			return
		}
		m.Nickname = proto.String(value)
	}
}

type MsgRoleSetNicknameRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgRoleSetNicknameRet) Reset()         { *m = MsgRoleSetNicknameRet{} }
func (m *MsgRoleSetNicknameRet) String() string { return proto.CompactTextString(m) }
func (*MsgRoleSetNicknameRet) ProtoMessage()    {}

func (m *MsgRoleSetNicknameRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgRoleSetNicknameRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

// 新手引导
type MsgRoleNewGuideUpdate struct {
	NewPlayerGuideStep *int32 `protobuf:"varint,1,req,name=new_player_guide_step" json:"new_player_guide_step,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *MsgRoleNewGuideUpdate) Reset()         { *m = MsgRoleNewGuideUpdate{} }
func (m *MsgRoleNewGuideUpdate) String() string { return proto.CompactTextString(m) }
func (*MsgRoleNewGuideUpdate) ProtoMessage()    {}

func (m *MsgRoleNewGuideUpdate) GetNewPlayerGuideStep() int32 {
	if m != nil && m.NewPlayerGuideStep != nil {
		return *m.NewPlayerGuideStep
	}
	return 0
}

func (m *MsgRoleNewGuideUpdate) SetNewPlayerGuideStep(value int32) {
	if m != nil {
		if m.NewPlayerGuideStep != nil {
			*m.NewPlayerGuideStep = value
			return
		}
		m.NewPlayerGuideStep = proto.Int32(value)
	}
}

type KingSkill struct {
	SkillId          *int32 `protobuf:"varint,1,req,name=skill_id" json:"skill_id,omitempty"`
	Lv               *int32 `protobuf:"varint,2,req,name=lv" json:"lv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KingSkill) Reset()         { *m = KingSkill{} }
func (m *KingSkill) String() string { return proto.CompactTextString(m) }
func (*KingSkill) ProtoMessage()    {}

func (m *KingSkill) GetSkillId() int32 {
	if m != nil && m.SkillId != nil {
		return *m.SkillId
	}
	return 0
}

func (m *KingSkill) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *KingSkill) SetSkillId(value int32) {
	if m != nil {
		if m.SkillId != nil {
			*m.SkillId = value
			return
		}
		m.SkillId = proto.Int32(value)
	}
}

func (m *KingSkill) SetLv(value int32) {
	if m != nil {
		if m.Lv != nil {
			*m.Lv = value
			return
		}
		m.Lv = proto.Int32(value)
	}
}

// 开始技能升级
type MsgRoleKingSkillLvUpReq struct {
	SkillId          *int32 `protobuf:"varint,1,req,name=skill_id" json:"skill_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgRoleKingSkillLvUpReq) Reset()         { *m = MsgRoleKingSkillLvUpReq{} }
func (m *MsgRoleKingSkillLvUpReq) String() string { return proto.CompactTextString(m) }
func (*MsgRoleKingSkillLvUpReq) ProtoMessage()    {}

func (m *MsgRoleKingSkillLvUpReq) GetSkillId() int32 {
	if m != nil && m.SkillId != nil {
		return *m.SkillId
	}
	return 0
}

func (m *MsgRoleKingSkillLvUpReq) SetSkillId(value int32) {
	if m != nil {
		if m.SkillId != nil {
			*m.SkillId = value
			return
		}
		m.SkillId = proto.Int32(value)
	}
}

type MsgRoleKingSkillLvUpRet struct {
	RetCode          *int32 `protobuf:"varint,1,req,name=retCode" json:"retCode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgRoleKingSkillLvUpRet) Reset()         { *m = MsgRoleKingSkillLvUpRet{} }
func (m *MsgRoleKingSkillLvUpRet) String() string { return proto.CompactTextString(m) }
func (*MsgRoleKingSkillLvUpRet) ProtoMessage()    {}

func (m *MsgRoleKingSkillLvUpRet) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *MsgRoleKingSkillLvUpRet) SetRetCode(value int32) {
	if m != nil {
		if m.RetCode != nil {
			*m.RetCode = value
			return
		}
		m.RetCode = proto.Int32(value)
	}
}

// 新手引导 掠夺奖励
type MsgGuidePlunderAwardReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgGuidePlunderAwardReq) Reset()         { *m = MsgGuidePlunderAwardReq{} }
func (m *MsgGuidePlunderAwardReq) String() string { return proto.CompactTextString(m) }
func (*MsgGuidePlunderAwardReq) ProtoMessage()    {}

type MsgGuidePlunderAwardRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgGuidePlunderAwardRet) Reset()         { *m = MsgGuidePlunderAwardRet{} }
func (m *MsgGuidePlunderAwardRet) String() string { return proto.CompactTextString(m) }
func (*MsgGuidePlunderAwardRet) ProtoMessage()    {}

func (m *MsgGuidePlunderAwardRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgGuidePlunderAwardRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

// 充值记录
type MsgRechargeReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgRechargeReq) Reset()         { *m = MsgRechargeReq{} }
func (m *MsgRechargeReq) String() string { return proto.CompactTextString(m) }
func (*MsgRechargeReq) ProtoMessage()    {}

type MsgRechargeRet struct {
	Records          map[int32]int32 `protobuf:"bytes,1,rep,name=records" json:"records,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MsgRechargeRet) Reset()         { *m = MsgRechargeRet{} }
func (m *MsgRechargeRet) String() string { return proto.CompactTextString(m) }
func (*MsgRechargeRet) ProtoMessage()    {}

func (m *MsgRechargeRet) GetRecords() map[int32]int32 {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *MsgRechargeRet) SetRecords(value map[int32]int32) {
	if m != nil {
		m.Records = value
	}
}

func init() {
	proto.RegisterEnum("protocol.RoleAttrType", RoleAttrType_name, RoleAttrType_value)
}
