// Code generated by protoc-gen-go.
// source: arena.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ArenaRankInfo struct {
	Rank             *int32  `protobuf:"varint,1,req,name=rank" json:"rank,omitempty"`
	RoleUid          *int64  `protobuf:"varint,2,req,name=role_uid" json:"role_uid,omitempty"`
	Nickname         *string `protobuf:"bytes,3,req,name=nickname" json:"nickname,omitempty"`
	RoleLv           *int32  `protobuf:"varint,4,req,name=role_lv" json:"role_lv,omitempty"`
	Score            *int32  `protobuf:"varint,5,req,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ArenaRankInfo) Reset()         { *m = ArenaRankInfo{} }
func (m *ArenaRankInfo) String() string { return proto.CompactTextString(m) }
func (*ArenaRankInfo) ProtoMessage()    {}

func (m *ArenaRankInfo) GetRank() int32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *ArenaRankInfo) GetRoleUid() int64 {
	if m != nil && m.RoleUid != nil {
		return *m.RoleUid
	}
	return 0
}

func (m *ArenaRankInfo) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *ArenaRankInfo) GetRoleLv() int32 {
	if m != nil && m.RoleLv != nil {
		return *m.RoleLv
	}
	return 0
}

func (m *ArenaRankInfo) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *ArenaRankInfo) SetRank(value int32) {
	if m != nil {
		if m.Rank != nil {
			*m.Rank = value
			return
		}
		m.Rank = proto.Int32(value)
	}
}

func (m *ArenaRankInfo) SetRoleUid(value int64) {
	if m != nil {
		if m.RoleUid != nil {
			*m.RoleUid = value
			return
		}
		m.RoleUid = proto.Int64(value)
	}
}

func (m *ArenaRankInfo) SetNickname(value string) {
	if m != nil {
		if m.Nickname != nil {
			*m.Nickname = value
			return
		}
		m.Nickname = proto.String(value)
	}
}

func (m *ArenaRankInfo) SetRoleLv(value int32) {
	if m != nil {
		if m.RoleLv != nil {
			*m.RoleLv = value
			return
		}
		m.RoleLv = proto.Int32(value)
	}
}

func (m *ArenaRankInfo) SetScore(value int32) {
	if m != nil {
		if m.Score != nil {
			*m.Score = value
			return
		}
		m.Score = proto.Int32(value)
	}
}

type MsgArenaQueryReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgArenaQueryReq) Reset()         { *m = MsgArenaQueryReq{} }
func (m *MsgArenaQueryReq) String() string { return proto.CompactTextString(m) }
func (*MsgArenaQueryReq) ProtoMessage()    {}

type MsgArenaQueryRet struct {
	Retcode          *int32           `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	MyPoint          *int32           `protobuf:"varint,2,req,name=my_point" json:"my_point,omitempty"`
	BestScore        *int32           `protobuf:"varint,3,req,name=best_score" json:"best_score,omitempty"`
	MyRank           *int32           `protobuf:"varint,4,req,name=my_rank" json:"my_rank,omitempty"`
	BossVersion      *int64           `protobuf:"varint,5,req,name=boss_version" json:"boss_version,omitempty"`
	BossId           *int32           `protobuf:"varint,6,req,name=boss_id" json:"boss_id,omitempty"`
	BossAward1       []*AwardInfo     `protobuf:"bytes,7,rep,name=boss_award1" json:"boss_award1,omitempty"`
	BossAward2       []*AwardInfo     `protobuf:"bytes,8,rep,name=boss_award2" json:"boss_award2,omitempty"`
	BossAward3       []*AwardInfo     `protobuf:"bytes,9,rep,name=boss_award3" json:"boss_award3,omitempty"`
	Infos            []*ArenaRankInfo `protobuf:"bytes,10,rep,name=infos" json:"infos,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MsgArenaQueryRet) Reset()         { *m = MsgArenaQueryRet{} }
func (m *MsgArenaQueryRet) String() string { return proto.CompactTextString(m) }
func (*MsgArenaQueryRet) ProtoMessage()    {}

func (m *MsgArenaQueryRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgArenaQueryRet) GetMyPoint() int32 {
	if m != nil && m.MyPoint != nil {
		return *m.MyPoint
	}
	return 0
}

func (m *MsgArenaQueryRet) GetBestScore() int32 {
	if m != nil && m.BestScore != nil {
		return *m.BestScore
	}
	return 0
}

func (m *MsgArenaQueryRet) GetMyRank() int32 {
	if m != nil && m.MyRank != nil {
		return *m.MyRank
	}
	return 0
}

func (m *MsgArenaQueryRet) GetBossVersion() int64 {
	if m != nil && m.BossVersion != nil {
		return *m.BossVersion
	}
	return 0
}

func (m *MsgArenaQueryRet) GetBossId() int32 {
	if m != nil && m.BossId != nil {
		return *m.BossId
	}
	return 0
}

func (m *MsgArenaQueryRet) GetBossAward1() []*AwardInfo {
	if m != nil {
		return m.BossAward1
	}
	return nil
}

func (m *MsgArenaQueryRet) GetBossAward2() []*AwardInfo {
	if m != nil {
		return m.BossAward2
	}
	return nil
}

func (m *MsgArenaQueryRet) GetBossAward3() []*AwardInfo {
	if m != nil {
		return m.BossAward3
	}
	return nil
}

func (m *MsgArenaQueryRet) GetInfos() []*ArenaRankInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (m *MsgArenaQueryRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgArenaQueryRet) SetMyPoint(value int32) {
	if m != nil {
		if m.MyPoint != nil {
			*m.MyPoint = value
			return
		}
		m.MyPoint = proto.Int32(value)
	}
}

func (m *MsgArenaQueryRet) SetBestScore(value int32) {
	if m != nil {
		if m.BestScore != nil {
			*m.BestScore = value
			return
		}
		m.BestScore = proto.Int32(value)
	}
}

func (m *MsgArenaQueryRet) SetMyRank(value int32) {
	if m != nil {
		if m.MyRank != nil {
			*m.MyRank = value
			return
		}
		m.MyRank = proto.Int32(value)
	}
}

func (m *MsgArenaQueryRet) SetBossVersion(value int64) {
	if m != nil {
		if m.BossVersion != nil {
			*m.BossVersion = value
			return
		}
		m.BossVersion = proto.Int64(value)
	}
}

func (m *MsgArenaQueryRet) SetBossId(value int32) {
	if m != nil {
		if m.BossId != nil {
			*m.BossId = value
			return
		}
		m.BossId = proto.Int32(value)
	}
}

func (m *MsgArenaQueryRet) SetBossAward1(value []*AwardInfo) {
	if m != nil {
		m.BossAward1 = value
	}
}

func (m *MsgArenaQueryRet) SetBossAward2(value []*AwardInfo) {
	if m != nil {
		m.BossAward2 = value
	}
}

func (m *MsgArenaQueryRet) SetBossAward3(value []*AwardInfo) {
	if m != nil {
		m.BossAward3 = value
	}
}

func (m *MsgArenaQueryRet) SetInfos(value []*ArenaRankInfo) {
	if m != nil {
		m.Infos = value
	}
}

type MsgArenaFightReq struct {
	Score            *int32 `protobuf:"varint,1,req,name=score" json:"score,omitempty"`
	IsCostOrder      *bool  `protobuf:"varint,2,req,name=is_cost_order" json:"is_cost_order,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgArenaFightReq) Reset()         { *m = MsgArenaFightReq{} }
func (m *MsgArenaFightReq) String() string { return proto.CompactTextString(m) }
func (*MsgArenaFightReq) ProtoMessage()    {}

func (m *MsgArenaFightReq) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *MsgArenaFightReq) GetIsCostOrder() bool {
	if m != nil && m.IsCostOrder != nil {
		return *m.IsCostOrder
	}
	return false
}

func (m *MsgArenaFightReq) SetScore(value int32) {
	if m != nil {
		if m.Score != nil {
			*m.Score = value
			return
		}
		m.Score = proto.Int32(value)
	}
}

func (m *MsgArenaFightReq) SetIsCostOrder(value bool) {
	if m != nil {
		if m.IsCostOrder != nil {
			*m.IsCostOrder = value
			return
		}
		m.IsCostOrder = proto.Bool(value)
	}
}

type MsgArenaFightRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	NewRank          *int32 `protobuf:"varint,2,req,name=new_rank" json:"new_rank,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgArenaFightRet) Reset()         { *m = MsgArenaFightRet{} }
func (m *MsgArenaFightRet) String() string { return proto.CompactTextString(m) }
func (*MsgArenaFightRet) ProtoMessage()    {}

func (m *MsgArenaFightRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgArenaFightRet) GetNewRank() int32 {
	if m != nil && m.NewRank != nil {
		return *m.NewRank
	}
	return 0
}

func (m *MsgArenaFightRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgArenaFightRet) SetNewRank(value int32) {
	if m != nil {
		if m.NewRank != nil {
			*m.NewRank = value
			return
		}
		m.NewRank = proto.Int32(value)
	}
}

type MsgArenaShopQueryReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgArenaShopQueryReq) Reset()         { *m = MsgArenaShopQueryReq{} }
func (m *MsgArenaShopQueryReq) String() string { return proto.CompactTextString(m) }
func (*MsgArenaShopQueryReq) ProtoMessage()    {}

type MsgArenaShopQueryRet struct {
	Retcode          *int32       `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Timestamp        *int64       `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	Panel            *int32       `protobuf:"varint,3,req,name=panel" json:"panel,omitempty"`
	Info             []*AwardInfo `protobuf:"bytes,4,rep,name=info" json:"info,omitempty"`
	ShopRecord       []int32      `protobuf:"varint,5,rep,name=shop_record" json:"shop_record,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MsgArenaShopQueryRet) Reset()         { *m = MsgArenaShopQueryRet{} }
func (m *MsgArenaShopQueryRet) String() string { return proto.CompactTextString(m) }
func (*MsgArenaShopQueryRet) ProtoMessage()    {}

func (m *MsgArenaShopQueryRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgArenaShopQueryRet) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *MsgArenaShopQueryRet) GetPanel() int32 {
	if m != nil && m.Panel != nil {
		return *m.Panel
	}
	return 0
}

func (m *MsgArenaShopQueryRet) GetInfo() []*AwardInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MsgArenaShopQueryRet) GetShopRecord() []int32 {
	if m != nil {
		return m.ShopRecord
	}
	return nil
}

func (m *MsgArenaShopQueryRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgArenaShopQueryRet) SetTimestamp(value int64) {
	if m != nil {
		if m.Timestamp != nil {
			*m.Timestamp = value
			return
		}
		m.Timestamp = proto.Int64(value)
	}
}

func (m *MsgArenaShopQueryRet) SetPanel(value int32) {
	if m != nil {
		if m.Panel != nil {
			*m.Panel = value
			return
		}
		m.Panel = proto.Int32(value)
	}
}

func (m *MsgArenaShopQueryRet) SetInfo(value []*AwardInfo) {
	if m != nil {
		m.Info = value
	}
}

func (m *MsgArenaShopQueryRet) SetShopRecord(value []int32) {
	if m != nil {
		m.ShopRecord = value
	}
}

type MsgArenaShopBuyReq struct {
	Pos              *int32 `protobuf:"varint,1,req,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgArenaShopBuyReq) Reset()         { *m = MsgArenaShopBuyReq{} }
func (m *MsgArenaShopBuyReq) String() string { return proto.CompactTextString(m) }
func (*MsgArenaShopBuyReq) ProtoMessage()    {}

func (m *MsgArenaShopBuyReq) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *MsgArenaShopBuyReq) SetPos(value int32) {
	if m != nil {
		if m.Pos != nil {
			*m.Pos = value
			return
		}
		m.Pos = proto.Int32(value)
	}
}

type MsgArenaShopBuyRet struct {
	Retcode          *int32 `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Pos              *int32 `protobuf:"varint,2,req,name=pos" json:"pos,omitempty"`
	Count            *int32 `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	RestPoint        *int32 `protobuf:"varint,4,req,name=rest_point" json:"rest_point,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgArenaShopBuyRet) Reset()         { *m = MsgArenaShopBuyRet{} }
func (m *MsgArenaShopBuyRet) String() string { return proto.CompactTextString(m) }
func (*MsgArenaShopBuyRet) ProtoMessage()    {}

func (m *MsgArenaShopBuyRet) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *MsgArenaShopBuyRet) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *MsgArenaShopBuyRet) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *MsgArenaShopBuyRet) GetRestPoint() int32 {
	if m != nil && m.RestPoint != nil {
		return *m.RestPoint
	}
	return 0
}

func (m *MsgArenaShopBuyRet) SetRetcode(value int32) {
	if m != nil {
		if m.Retcode != nil {
			*m.Retcode = value
			return
		}
		m.Retcode = proto.Int32(value)
	}
}

func (m *MsgArenaShopBuyRet) SetPos(value int32) {
	if m != nil {
		if m.Pos != nil {
			*m.Pos = value
			return
		}
		m.Pos = proto.Int32(value)
	}
}

func (m *MsgArenaShopBuyRet) SetCount(value int32) {
	if m != nil {
		if m.Count != nil {
			*m.Count = value
			return
		}
		m.Count = proto.Int32(value)
	}
}

func (m *MsgArenaShopBuyRet) SetRestPoint(value int32) {
	if m != nil {
		if m.RestPoint != nil {
			*m.RestPoint = value
			return
		}
		m.RestPoint = proto.Int32(value)
	}
}

func init() {
}
