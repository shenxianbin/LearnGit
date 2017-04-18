// Code generated by protoc-gen-go.
// source: chat.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MessageType int32

const (
	MessageType_Normal      MessageType = 1
	MessageType_Battlefield MessageType = 2
)

var MessageType_name = map[int32]string{
	1: "Normal",
	2: "Battlefield",
}
var MessageType_value = map[string]int32{
	"Normal":      1,
	"Battlefield": 2,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}

type ChatType int32

const (
	ChatType_Private ChatType = 1
	ChatType_World   ChatType = 2
)

var ChatType_name = map[int32]string{
	1: "Private",
	2: "World",
}
var ChatType_value = map[string]int32{
	"Private": 1,
	"World":   2,
}

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}
func (x ChatType) String() string {
	return proto.EnumName(ChatType_name, int32(x))
}
func (x *ChatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChatType_value, data, "ChatType")
	if err != nil {
		return err
	}
	*x = ChatType(value)
	return nil
}

type ChatInfo struct {
	TimeStamp        *int64 `protobuf:"varint,1,req,name=time_stamp" json:"time_stamp,omitempty"`
	Content          []byte `protobuf:"bytes,2,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChatInfo) Reset()         { *m = ChatInfo{} }
func (m *ChatInfo) String() string { return proto.CompactTextString(m) }
func (*ChatInfo) ProtoMessage()    {}

func (m *ChatInfo) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *ChatInfo) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ChatInfo) SetTimeStamp(value int64) {
	if m != nil {
		if m.TimeStamp != nil {
			*m.TimeStamp = value
			return
		}
		m.TimeStamp = proto.Int64(value)
	}
}

func (m *ChatInfo) SetContent(value []byte) {
	if m != nil {
		if m.Content != nil {
			m.Content = value
			return
		}
		m.Content = value
	}
}

type ChatContent struct {
	MessageType      *int32 `protobuf:"varint,1,req,name=message_type" json:"message_type,omitempty"`
	Content          []byte `protobuf:"bytes,2,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChatContent) Reset()         { *m = ChatContent{} }
func (m *ChatContent) String() string { return proto.CompactTextString(m) }
func (*ChatContent) ProtoMessage()    {}

func (m *ChatContent) GetMessageType() int32 {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return 0
}

func (m *ChatContent) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ChatContent) SetMessageType(value int32) {
	if m != nil {
		if m.MessageType != nil {
			*m.MessageType = value
			return
		}
		m.MessageType = proto.Int32(value)
	}
}

func (m *ChatContent) SetContent(value []byte) {
	if m != nil {
		if m.Content != nil {
			m.Content = value
			return
		}
		m.Content = value
	}
}

type ChatMessageInfo struct {
	RoleUid          *int64 `protobuf:"varint,1,req,name=role_uid" json:"role_uid,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Nickname         []byte `protobuf:"bytes,3,req,name=nickname" json:"nickname,omitempty"`
	Content          []byte `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	TargetRoleUid    *int64 `protobuf:"varint,5,opt,name=target_role_uid" json:"target_role_uid,omitempty"`
	TargetLevel      *int32 `protobuf:"varint,6,opt,name=target_level" json:"target_level,omitempty"`
	TargetNickname   []byte `protobuf:"bytes,7,opt,name=target_nickname" json:"target_nickname,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChatMessageInfo) Reset()         { *m = ChatMessageInfo{} }
func (m *ChatMessageInfo) String() string { return proto.CompactTextString(m) }
func (*ChatMessageInfo) ProtoMessage()    {}

func (m *ChatMessageInfo) GetRoleUid() int64 {
	if m != nil && m.RoleUid != nil {
		return *m.RoleUid
	}
	return 0
}

func (m *ChatMessageInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *ChatMessageInfo) GetNickname() []byte {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *ChatMessageInfo) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ChatMessageInfo) GetTargetRoleUid() int64 {
	if m != nil && m.TargetRoleUid != nil {
		return *m.TargetRoleUid
	}
	return 0
}

func (m *ChatMessageInfo) GetTargetLevel() int32 {
	if m != nil && m.TargetLevel != nil {
		return *m.TargetLevel
	}
	return 0
}

func (m *ChatMessageInfo) GetTargetNickname() []byte {
	if m != nil {
		return m.TargetNickname
	}
	return nil
}

func (m *ChatMessageInfo) SetRoleUid(value int64) {
	if m != nil {
		if m.RoleUid != nil {
			*m.RoleUid = value
			return
		}
		m.RoleUid = proto.Int64(value)
	}
}

func (m *ChatMessageInfo) SetLevel(value int32) {
	if m != nil {
		if m.Level != nil {
			*m.Level = value
			return
		}
		m.Level = proto.Int32(value)
	}
}

func (m *ChatMessageInfo) SetNickname(value []byte) {
	if m != nil {
		if m.Nickname != nil {
			m.Nickname = value
			return
		}
		m.Nickname = value
	}
}

func (m *ChatMessageInfo) SetContent(value []byte) {
	if m != nil {
		if m.Content != nil {
			m.Content = value
			return
		}
		m.Content = value
	}
}

func (m *ChatMessageInfo) SetTargetRoleUid(value int64) {
	if m != nil {
		if m.TargetRoleUid != nil {
			*m.TargetRoleUid = value
			return
		}
		m.TargetRoleUid = proto.Int64(value)
	}
}

func (m *ChatMessageInfo) SetTargetLevel(value int32) {
	if m != nil {
		if m.TargetLevel != nil {
			*m.TargetLevel = value
			return
		}
		m.TargetLevel = proto.Int32(value)
	}
}

func (m *ChatMessageInfo) SetTargetNickname(value []byte) {
	if m != nil {
		if m.TargetNickname != nil {
			m.TargetNickname = value
			return
		}
		m.TargetNickname = value
	}
}

type MsgChatQueryReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgChatQueryReq) Reset()         { *m = MsgChatQueryReq{} }
func (m *MsgChatQueryReq) String() string { return proto.CompactTextString(m) }
func (*MsgChatQueryReq) ProtoMessage()    {}

type MsgChatQueryRet struct {
	Retcode          *int32      `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Infos            []*ChatInfo `protobuf:"bytes,3,rep,name=infos" json:"infos,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MsgChatQueryRet) Reset()         { *m = MsgChatQueryRet{} }
func (m *MsgChatQueryRet) String() string { return proto.CompactTextString(m) }
func (*MsgChatQueryRet) ProtoMessage()    {}

func (m *MsgChatQueryRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgChatQueryRet) GetInfos() []*ChatInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (m *MsgChatQueryRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgChatQueryRet) SetInfos(value []*ChatInfo) {
	if m != nil {
		m.Infos = value
	}
}

type MsgChatReq struct {
	ChatType         *int32 `protobuf:"varint,1,req,name=chat_type" json:"chat_type,omitempty"`
	RoleUid          *int64 `protobuf:"varint,2,req,name=role_uid" json:"role_uid,omitempty"`
	Content          []byte `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgChatReq) Reset()         { *m = MsgChatReq{} }
func (m *MsgChatReq) String() string { return proto.CompactTextString(m) }
func (*MsgChatReq) ProtoMessage()    {}

func (m *MsgChatReq) GetChatType() int32 {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return 0
}

func (m *MsgChatReq) GetRoleUid() int64 {
	if m != nil && m.RoleUid != nil {
		return *m.RoleUid
	}
	return 0
}

func (m *MsgChatReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MsgChatReq) SetChatType(value int32) {
	if m != nil {
		if m.ChatType != nil {
			*m.ChatType = value
			return
		}
		m.ChatType = proto.Int32(value)
	}
}

func (m *MsgChatReq) SetRoleUid(value int64) {
	if m != nil {
		if m.RoleUid != nil {
			*m.RoleUid = value
			return
		}
		m.RoleUid = proto.Int64(value)
	}
}

func (m *MsgChatReq) SetContent(value []byte) {
	if m != nil {
		if m.Content != nil {
			m.Content = value
			return
		}
		m.Content = value
	}
}

type MsgChatRet struct {
	Retcode          *int32    `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	ChatType         *int32    `protobuf:"varint,2,req,name=chat_type" json:"chat_type,omitempty"`
	Info             *ChatInfo `protobuf:"bytes,3,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MsgChatRet) Reset()         { *m = MsgChatRet{} }
func (m *MsgChatRet) String() string { return proto.CompactTextString(m) }
func (*MsgChatRet) ProtoMessage()    {}

func (m *MsgChatRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgChatRet) GetChatType() int32 {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return 0
}

func (m *MsgChatRet) GetInfo() *ChatInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgChatRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgChatRet) SetChatType(value int32) {
	if m != nil {
		if m.ChatType != nil {
			*m.ChatType = value
			return
		}
		m.ChatType = proto.Int32(value)
	}
}

func (m *MsgChatRet) SetInfo(value *ChatInfo) {
	if m != nil {
		m.Info = value
	}
}

type MsgChatNotify struct {
	ChatType         *int32    `protobuf:"varint,1,req,name=chat_type" json:"chat_type,omitempty"`
	Info             *ChatInfo `protobuf:"bytes,2,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MsgChatNotify) Reset()         { *m = MsgChatNotify{} }
func (m *MsgChatNotify) String() string { return proto.CompactTextString(m) }
func (*MsgChatNotify) ProtoMessage()    {}

func (m *MsgChatNotify) GetChatType() int32 {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return 0
}

func (m *MsgChatNotify) GetInfo() *ChatInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgChatNotify) SetChatType(value int32) {
	if m != nil {
		if m.ChatType != nil {
			*m.ChatType = value
			return
		}
		m.ChatType = proto.Int32(value)
	}
}

func (m *MsgChatNotify) SetInfo(value *ChatInfo) {
	if m != nil {
		m.Info = value
	}
}

func init() {
	proto.RegisterEnum("protocol.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("protocol.ChatType", ChatType_name, ChatType_value)
}