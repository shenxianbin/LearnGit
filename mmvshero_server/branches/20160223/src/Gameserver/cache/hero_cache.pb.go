// Code generated by protoc-gen-go.
// source: hero_cache.proto
// DO NOT EDIT!

package cache

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HeroCreateCache struct {
	Cd               *int64 `protobuf:"varint,1,req,name=cd" json:"cd,omitempty"`
	CreateId         *int32 `protobuf:"varint,2,req,name=create_id" json:"create_id,omitempty"`
	CostOrderPlanId  *int32 `protobuf:"varint,3,req,name=cost_order_plan_id" json:"cost_order_plan_id,omitempty"`
	StartTimestamp   *int64 `protobuf:"varint,4,req,name=start_timestamp" json:"start_timestamp,omitempty"`
	DeathTimestamp   *int64 `protobuf:"varint,5,req,name=death_timestamp" json:"death_timestamp,omitempty"`
	FixMagic         *int32 `protobuf:"varint,6,req,name=fix_magic" json:"fix_magic,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeroCreateCache) Reset()         { *m = HeroCreateCache{} }
func (m *HeroCreateCache) String() string { return proto.CompactTextString(m) }
func (*HeroCreateCache) ProtoMessage()    {}

func (m *HeroCreateCache) GetCd() int64 {
	if m != nil && m.Cd != nil {
		return *m.Cd
	}
	return 0
}

func (m *HeroCreateCache) GetCreateId() int32 {
	if m != nil && m.CreateId != nil {
		return *m.CreateId
	}
	return 0
}

func (m *HeroCreateCache) GetCostOrderPlanId() int32 {
	if m != nil && m.CostOrderPlanId != nil {
		return *m.CostOrderPlanId
	}
	return 0
}

func (m *HeroCreateCache) GetStartTimestamp() int64 {
	if m != nil && m.StartTimestamp != nil {
		return *m.StartTimestamp
	}
	return 0
}

func (m *HeroCreateCache) GetDeathTimestamp() int64 {
	if m != nil && m.DeathTimestamp != nil {
		return *m.DeathTimestamp
	}
	return 0
}

func (m *HeroCreateCache) GetFixMagic() int32 {
	if m != nil && m.FixMagic != nil {
		return *m.FixMagic
	}
	return 0
}

func (m *HeroCreateCache) SetCd(value int64) {
	if m != nil {
		if m.Cd != nil {
			*m.Cd = value
			return
		}
		m.Cd = proto.Int64(value)
	}
}

func (m *HeroCreateCache) SetCreateId(value int32) {
	if m != nil {
		if m.CreateId != nil {
			*m.CreateId = value
			return
		}
		m.CreateId = proto.Int32(value)
	}
}

func (m *HeroCreateCache) SetCostOrderPlanId(value int32) {
	if m != nil {
		if m.CostOrderPlanId != nil {
			*m.CostOrderPlanId = value
			return
		}
		m.CostOrderPlanId = proto.Int32(value)
	}
}

func (m *HeroCreateCache) SetStartTimestamp(value int64) {
	if m != nil {
		if m.StartTimestamp != nil {
			*m.StartTimestamp = value
			return
		}
		m.StartTimestamp = proto.Int64(value)
	}
}

func (m *HeroCreateCache) SetDeathTimestamp(value int64) {
	if m != nil {
		if m.DeathTimestamp != nil {
			*m.DeathTimestamp = value
			return
		}
		m.DeathTimestamp = proto.Int64(value)
	}
}

func (m *HeroCreateCache) SetFixMagic(value int32) {
	if m != nil {
		if m.FixMagic != nil {
			*m.FixMagic = value
			return
		}
		m.FixMagic = proto.Int32(value)
	}
}

type HeroSkillCache struct {
	SkillId          *int32 `protobuf:"varint,1,req,name=skill_id" json:"skill_id,omitempty"`
	SkillLv          *int32 `protobuf:"varint,2,req,name=skill_lv" json:"skill_lv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeroSkillCache) Reset()         { *m = HeroSkillCache{} }
func (m *HeroSkillCache) String() string { return proto.CompactTextString(m) }
func (*HeroSkillCache) ProtoMessage()    {}

func (m *HeroSkillCache) GetSkillId() int32 {
	if m != nil && m.SkillId != nil {
		return *m.SkillId
	}
	return 0
}

func (m *HeroSkillCache) GetSkillLv() int32 {
	if m != nil && m.SkillLv != nil {
		return *m.SkillLv
	}
	return 0
}

func (m *HeroSkillCache) SetSkillId(value int32) {
	if m != nil {
		if m.SkillId != nil {
			*m.SkillId = value
			return
		}
		m.SkillId = proto.Int32(value)
	}
}

func (m *HeroSkillCache) SetSkillLv(value int32) {
	if m != nil {
		if m.SkillLv != nil {
			*m.SkillLv = value
			return
		}
		m.SkillLv = proto.Int32(value)
	}
}

type HeroCache struct {
	Uid              *int64                    `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	SchemeId         *int32                    `protobuf:"varint,2,req,name=scheme_id" json:"scheme_id,omitempty"`
	Lv               *int32                    `protobuf:"varint,3,req,name=lv" json:"lv,omitempty"`
	LvExp            *int32                    `protobuf:"varint,4,req,name=lv_exp" json:"lv_exp,omitempty"`
	Stage            *int32                    `protobuf:"varint,5,req,name=stage" json:"stage,omitempty"`
	StageTimestamp   *int64                    `protobuf:"varint,6,req,name=stage_timestamp" json:"stage_timestamp,omitempty"`
	StageSpeedup     *int64                    `protobuf:"varint,7,req,name=stage_speedup" json:"stage_speedup,omitempty"`
	Rank             *int32                    `protobuf:"varint,8,req,name=rank" json:"rank,omitempty"`
	RankExp          *int32                    `protobuf:"varint,9,req,name=rank_exp" json:"rank_exp,omitempty"`
	SkillList        map[int32]*HeroSkillCache `protobuf:"bytes,10,rep,name=skill_list" json:"skill_list,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *HeroCache) Reset()         { *m = HeroCache{} }
func (m *HeroCache) String() string { return proto.CompactTextString(m) }
func (*HeroCache) ProtoMessage()    {}

func (m *HeroCache) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *HeroCache) GetSchemeId() int32 {
	if m != nil && m.SchemeId != nil {
		return *m.SchemeId
	}
	return 0
}

func (m *HeroCache) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *HeroCache) GetLvExp() int32 {
	if m != nil && m.LvExp != nil {
		return *m.LvExp
	}
	return 0
}

func (m *HeroCache) GetStage() int32 {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return 0
}

func (m *HeroCache) GetStageTimestamp() int64 {
	if m != nil && m.StageTimestamp != nil {
		return *m.StageTimestamp
	}
	return 0
}

func (m *HeroCache) GetStageSpeedup() int64 {
	if m != nil && m.StageSpeedup != nil {
		return *m.StageSpeedup
	}
	return 0
}

func (m *HeroCache) GetRank() int32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *HeroCache) GetRankExp() int32 {
	if m != nil && m.RankExp != nil {
		return *m.RankExp
	}
	return 0
}

func (m *HeroCache) GetSkillList() map[int32]*HeroSkillCache {
	if m != nil {
		return m.SkillList
	}
	return nil
}

func (m *HeroCache) SetUid(value int64) {
	if m != nil {
		if m.Uid != nil {
			*m.Uid = value
			return
		}
		m.Uid = proto.Int64(value)
	}
}

func (m *HeroCache) SetSchemeId(value int32) {
	if m != nil {
		if m.SchemeId != nil {
			*m.SchemeId = value
			return
		}
		m.SchemeId = proto.Int32(value)
	}
}

func (m *HeroCache) SetLv(value int32) {
	if m != nil {
		if m.Lv != nil {
			*m.Lv = value
			return
		}
		m.Lv = proto.Int32(value)
	}
}

func (m *HeroCache) SetLvExp(value int32) {
	if m != nil {
		if m.LvExp != nil {
			*m.LvExp = value
			return
		}
		m.LvExp = proto.Int32(value)
	}
}

func (m *HeroCache) SetStage(value int32) {
	if m != nil {
		if m.Stage != nil {
			*m.Stage = value
			return
		}
		m.Stage = proto.Int32(value)
	}
}

func (m *HeroCache) SetStageTimestamp(value int64) {
	if m != nil {
		if m.StageTimestamp != nil {
			*m.StageTimestamp = value
			return
		}
		m.StageTimestamp = proto.Int64(value)
	}
}

func (m *HeroCache) SetStageSpeedup(value int64) {
	if m != nil {
		if m.StageSpeedup != nil {
			*m.StageSpeedup = value
			return
		}
		m.StageSpeedup = proto.Int64(value)
	}
}

func (m *HeroCache) SetRank(value int32) {
	if m != nil {
		if m.Rank != nil {
			*m.Rank = value
			return
		}
		m.Rank = proto.Int32(value)
	}
}

func (m *HeroCache) SetRankExp(value int32) {
	if m != nil {
		if m.RankExp != nil {
			*m.RankExp = value
			return
		}
		m.RankExp = proto.Int32(value)
	}
}

func (m *HeroCache) SetSkillList(value map[int32]*HeroSkillCache) {
	if m != nil {
		m.SkillList = value
	}
}

func init() {
}
