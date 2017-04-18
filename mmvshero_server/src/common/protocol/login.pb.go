// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MsgLoginAuthReq struct {
	TokenKey         *string `protobuf:"bytes,1,req,name=token_key" json:"token_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgLoginAuthReq) Reset()         { *m = MsgLoginAuthReq{} }
func (m *MsgLoginAuthReq) String() string { return proto.CompactTextString(m) }
func (*MsgLoginAuthReq) ProtoMessage()    {}

func (m *MsgLoginAuthReq) GetTokenKey() string {
	if m != nil && m.TokenKey != nil {
		return *m.TokenKey
	}
	return ""
}

func (m *MsgLoginAuthReq) SetTokenKey(value string) {
	if m != nil {
		if m.TokenKey != nil {
			*m.TokenKey = value
			return
		}
		m.TokenKey = proto.String(value)
	}
}

type MsgLoginAuthRet struct {
	RetCode          *int32 `protobuf:"varint,1,req,name=ret_code" json:"ret_code,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgLoginAuthRet) Reset()         { *m = MsgLoginAuthRet{} }
func (m *MsgLoginAuthRet) String() string { return proto.CompactTextString(m) }
func (*MsgLoginAuthRet) ProtoMessage()    {}

func (m *MsgLoginAuthRet) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *MsgLoginAuthRet) SetRetCode(value int32) {
	if m != nil {
		if m.RetCode != nil {
			*m.RetCode = value
			return
		}
		m.RetCode = proto.Int32(value)
	}
}

type MsgLoginInReq struct {
	TokenKey         *string `protobuf:"bytes,1,req,name=token_key" json:"token_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgLoginInReq) Reset()         { *m = MsgLoginInReq{} }
func (m *MsgLoginInReq) String() string { return proto.CompactTextString(m) }
func (*MsgLoginInReq) ProtoMessage()    {}

func (m *MsgLoginInReq) GetTokenKey() string {
	if m != nil && m.TokenKey != nil {
		return *m.TokenKey
	}
	return ""
}

func (m *MsgLoginInReq) SetTokenKey(value string) {
	if m != nil {
		if m.TokenKey != nil {
			*m.TokenKey = value
			return
		}
		m.TokenKey = proto.String(value)
	}
}

type MsgLoginInRet struct {
	SystemTime       *int64            `protobuf:"varint,1,req,name=system_time" json:"system_time,omitempty"`
	OpenServer       *int64            `protobuf:"varint,2,req,name=open_server" json:"open_server,omitempty"`
	RoleInfo         *RoleInfo         `protobuf:"bytes,3,req,name=role_info" json:"role_info,omitempty"`
	ItemListInfo     *ItemListInfo     `protobuf:"bytes,4,req,name=item_list_info" json:"item_list_info,omitempty"`
	HeroListInfo     *HeroListInfo     `protobuf:"bytes,5,req,name=hero_list_info" json:"hero_list_info,omitempty"`
	AllSoldiers      *AllSoldiers      `protobuf:"bytes,6,req,name=all_soldiers" json:"all_soldiers,omitempty"`
	BuildingInfo     *BuildingListInfo `protobuf:"bytes,7,req,name=building_info" json:"building_info,omitempty"`
	MapInfo          *MapInfo          `protobuf:"bytes,8,req,name=map_info" json:"map_info,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MsgLoginInRet) Reset()         { *m = MsgLoginInRet{} }
func (m *MsgLoginInRet) String() string { return proto.CompactTextString(m) }
func (*MsgLoginInRet) ProtoMessage()    {}

func (m *MsgLoginInRet) GetSystemTime() int64 {
	if m != nil && m.SystemTime != nil {
		return *m.SystemTime
	}
	return 0
}

func (m *MsgLoginInRet) GetOpenServer() int64 {
	if m != nil && m.OpenServer != nil {
		return *m.OpenServer
	}
	return 0
}

func (m *MsgLoginInRet) GetRoleInfo() *RoleInfo {
	if m != nil {
		return m.RoleInfo
	}
	return nil
}

func (m *MsgLoginInRet) GetItemListInfo() *ItemListInfo {
	if m != nil {
		return m.ItemListInfo
	}
	return nil
}

func (m *MsgLoginInRet) GetHeroListInfo() *HeroListInfo {
	if m != nil {
		return m.HeroListInfo
	}
	return nil
}

func (m *MsgLoginInRet) GetAllSoldiers() *AllSoldiers {
	if m != nil {
		return m.AllSoldiers
	}
	return nil
}

func (m *MsgLoginInRet) GetBuildingInfo() *BuildingListInfo {
	if m != nil {
		return m.BuildingInfo
	}
	return nil
}

func (m *MsgLoginInRet) GetMapInfo() *MapInfo {
	if m != nil {
		return m.MapInfo
	}
	return nil
}

func (m *MsgLoginInRet) SetSystemTime(value int64) {
	if m != nil {
		if m.SystemTime != nil {
			*m.SystemTime = value
			return
		}
		m.SystemTime = proto.Int64(value)
	}
}

func (m *MsgLoginInRet) SetOpenServer(value int64) {
	if m != nil {
		if m.OpenServer != nil {
			*m.OpenServer = value
			return
		}
		m.OpenServer = proto.Int64(value)
	}
}

func (m *MsgLoginInRet) SetRoleInfo(value *RoleInfo) {
	if m != nil {
		m.RoleInfo = value
	}
}

func (m *MsgLoginInRet) SetItemListInfo(value *ItemListInfo) {
	if m != nil {
		m.ItemListInfo = value
	}
}

func (m *MsgLoginInRet) SetHeroListInfo(value *HeroListInfo) {
	if m != nil {
		m.HeroListInfo = value
	}
}

func (m *MsgLoginInRet) SetAllSoldiers(value *AllSoldiers) {
	if m != nil {
		m.AllSoldiers = value
	}
}

func (m *MsgLoginInRet) SetBuildingInfo(value *BuildingListInfo) {
	if m != nil {
		m.BuildingInfo = value
	}
}

func (m *MsgLoginInRet) SetMapInfo(value *MapInfo) {
	if m != nil {
		m.MapInfo = value
	}
}

func init() {
}
