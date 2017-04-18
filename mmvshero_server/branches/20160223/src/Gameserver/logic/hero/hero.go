package hero

import (
	. "Gameserver/cache"
	"Gameserver/global"
	. "Gameserver/logic"
	"Gameserver/logic/award"
	"common"
	"common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"errors"
	"fmt"
	. "galaxy"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	HEROCREATE_SECOND = 1
)

const (
	cache_hero_autokey_t   = "Role:%v:HeroAutoKey"
	cache_herolist_key_t   = "Role:%v:Hero"
	cache_heroobj_key_t    = "Role:%v:Hero:%v"
	cache_herocreate_key_t = "Role:%v:HeroCreate"
)

func GenHeroListKey(role_uid int64) string {
	return fmt.Sprintf(cache_herolist_key_t, role_uid)
}

func GenHeroCacheKey(role_uid int64, hero_uid int64) string {
	return fmt.Sprintf(cache_heroobj_key_t, role_uid, hero_uid)
}

func genHeroCreateCacheKey(role_uid int64, hero_create_id int32) string {
	return fmt.Sprintf(cache_herocreate_key_t, role_uid, hero_create_id)
}

func genHeroAutoKey(role_uid int64) string {
	return fmt.Sprintf(cache_hero_autokey_t, role_uid)
}

type HeroCreate struct {
	HeroCreateCache
	scheme_data *scheme.MagicHeroCreate
}

func LoadHeroCreate(buf []byte) (*HeroCreate, error) {
	obj := new(HeroCreate)
	err := proto.Unmarshal(buf, &obj.HeroCreateCache)
	if err != nil {
		return nil, err
	}

	if obj.GetCreateId() != 0 {
		scheme_data, has := scheme.MagicHeroCreatemap[obj.GetCreateId()]
		if !has {
			return nil, errors.New("LoadHeroCreate Scheme Error")
		}
		obj.scheme_data = scheme_data
	}

	return obj, nil
}

func (this *HeroCreate) Start(create_id int32, cost_order_plan_id int32, start_timestamp int64) error {
	if this.HeroCreateCache.GetCd() != 0 && this.HeroCreateCache.GetCd() < time.Now().Unix() {
		return errors.New("HeroCreateStart CoolDown Now")
	}

	if this.HeroCreateCache.GetCreateId() != 0 {
		return errors.New("HeroCreateStart AlreadyStart")
	}

	scheme_data, has := scheme.MagicHeroCreatemap[create_id]
	if !has {
		return errors.New("HeroCreateStart Scheme Error")
	}

	scheme_plan, has := scheme.HeroCreatePlanmap[cost_order_plan_id]
	if !has {
		return errors.New("HeroCreatePlanStart Scheme Error")
	}

	this.HeroCreateCache.SetCd(0)
	this.HeroCreateCache.SetCreateId(create_id)
	this.HeroCreateCache.SetCostOrderPlanId(cost_order_plan_id)
	now := time.Now().Unix()
	this.HeroCreateCache.SetStartTimestamp(now)
	this.HeroCreateCache.SetDeathTimestamp(now + int64(scheme_plan.DeathTime*HEROCREATE_SECOND))
	this.HeroCreateCache.SetFixMagic(0)
	this.scheme_data = scheme_data
	return nil
}

func (this *HeroCreate) Reset() {
	this.HeroCreateCache.SetCd(time.Now().Unix() + int64(scheme.Commonmap[define.MagicHeroCreateCD].Value))
	this.HeroCreateCache.SetCreateId(0)
	this.HeroCreateCache.SetCostOrderPlanId(0)
	this.HeroCreateCache.SetStartTimestamp(0)
	this.HeroCreateCache.SetDeathTimestamp(0)
	this.HeroCreateCache.SetFixMagic(0)
	this.scheme_data = nil
}

func (this *HeroCreate) ReBorn() error {
	if !this.IsDie() {
		return errors.New("HeroCreateReBor Error")
	}

	scheme_plan, has := scheme.HeroCreatePlanmap[this.HeroCreateCache.GetCostOrderPlanId()]
	if !has {
		return errors.New("HeroCreatePlanStart Scheme Error")
	}

	this.HeroCreateCache.SetFixMagic(this.GetHeroCreateCurMagic())
	now := time.Now().Unix()
	this.HeroCreateCache.SetStartTimestamp(now)
	this.HeroCreateCache.SetDeathTimestamp(now + int64(scheme_plan.ReliveDeathTime*HEROCREATE_SECOND))
	return nil
}

func (this *HeroCreate) IsFinish() bool {
	return this.GetHeroCreateCurMagic() >= this.GetHeroCreateNeedMagic()
}

func (this *HeroCreate) IsDie() bool {
	if this.HeroCreateCache.GetCreateId() != 0 && !this.IsFinish() {
		return time.Now().Unix() > this.HeroCreateCache.GetDeathTimestamp()
	}
	return false
}

func (this *HeroCreate) IsCD() bool {
	return time.Now().Unix() < this.HeroCreateCache.GetCd()
}

func (this *HeroCreate) AddHeroCreateOrder(count int32) bool {
	if this.HeroCreateCache.GetCreateId() != 0 && !this.IsDie() {
		this.HeroCreateCache.SetDeathTimestamp(this.HeroCreateCache.GetDeathTimestamp() + int64(count))
		return true
	}
	return false
}

func (this *HeroCreate) AddHeroCreateMagic(magic int32) bool {
	if this.HeroCreateCache.GetCreateId() != 0 && !this.IsDie() {
		this.HeroCreateCache.SetFixMagic(this.HeroCreateCache.GetFixMagic() + magic)
		return true
	}
	return false
}

func (this *HeroCreate) GetHeroCreateCurMagic() int32 {
	return int32(this.HeroCreateCache.GetDeathTimestamp()-this.HeroCreateCache.GetStartTimestamp()) + this.HeroCreateCache.GetFixMagic()
}

func (this *HeroCreate) GetHeroCreateNeedMagic() int32 {
	if this.HeroCreateCache.GetCreateId() != 0 {
		return this.scheme_data.MagicValue
	}
	return 0
}

func (this *HeroCreate) GetHeroCreateDropId() int32 {
	if this.HeroCreateCache.GetCreateId() != 0 {
		return this.scheme_data.Drop
	}
	return 0
}

func (this *HeroCreate) FillHeroCreateInfo() *protocol.HeroCreateInfo {
	msg := new(protocol.HeroCreateInfo)
	msg.SetCd(this.HeroCreateCache.GetCd())
	msg.SetCreateId(this.HeroCreateCache.GetCreateId())
	msg.SetCostOrderPlanId(this.HeroCreateCache.GetCostOrderPlanId())
	msg.SetStartTimestamp(this.HeroCreateCache.GetStartTimestamp())
	msg.SetDeathTimestamp(this.HeroCreateCache.GetDeathTimestamp())
	msg.SetFixMagic(this.HeroCreateCache.GetFixMagic())
	return msg
}

type Hero struct {
	HeroCache
	scheme_base_data  *scheme.Hero
	scheme_lvup_data  *scheme.HeroLvUp
	scheme_stage_data *scheme.HeroStageUp
	scheme_rank_data  *scheme.HeroRank
}

func NewHero(scheme_id int32, lv int32, rank int32, role_uid int64) (*Hero, error) {
	scheme_base_data, has := scheme.Heromap[scheme_id]
	if !has {
		return nil, errors.New(fmt.Sprintf("NewHero Base Scheme Error (scheme_id : %v)", scheme_id))
	}
	scheme_lvup_data := scheme.HeroLvUpGet(scheme_id, lv, rank)
	if scheme_lvup_data == nil {
		return nil, errors.New(fmt.Sprintf("NewHero LvUp Scheme Error (scheme_id : %v) (lv : %v) (rank : %v)", scheme_id, lv, rank))
	}
	scheme_stage_data := scheme.HeroStageUpGet(scheme_id, scheme_lvup_data.Stage, rank)
	if scheme_stage_data == nil {
		return nil, errors.New(fmt.Sprintf("NewHero StageUp Scheme Error (scheme_id : %v) (stage : %v) (rank : %v)", scheme_id, scheme_lvup_data.Stage, rank))
	}
	scheme_rank_data := scheme.HeroRankGet(scheme_base_data.MagicHeroRankId, rank)
	if scheme_rank_data == nil {
		return nil, errors.New(fmt.Sprintf("NewHero Rank Scheme Error (MagicHeroRankId : %v) (rank : %v)", scheme_base_data.MagicHeroRankId, rank))
	}

	obj := new(Hero)
	resp, err := GxService().Redis().Cmd("INCR", genHeroAutoKey(role_uid))
	if err != nil {
		return nil, err
	}

	uid, _ := resp.Int64()
	obj.HeroCache.SetUid(uid)
	obj.HeroCache.SetSchemeId(scheme_id)
	obj.HeroCache.SetLv(lv)
	obj.HeroCache.SetLvExp(0)
	obj.HeroCache.SetStage(scheme_lvup_data.Stage)
	obj.HeroCache.SetStageTimestamp(0)
	obj.HeroCache.SetStageSpeedup(0)
	obj.HeroCache.SetRank(rank)
	obj.HeroCache.SetRankExp(0)
	skill_map := make(map[int32]*HeroSkillCache)
	skill_list := strings.Split(scheme_stage_data.SkillId, ";")
	for _, skill_data := range skill_list {
		if skill_id, err := strconv.Atoi(skill_data); err == nil {
			if skill_id != -1 {
				skill_cache := new(HeroSkillCache)
				skill_cache.SetSkillId(int32(skill_id))
				skill_cache.SetSkillLv(1)
				skill_map[int32(skill_id)] = skill_cache
			}
		} else {
			return nil, errors.New("NewHero Skill_id Scheme Error")
		}
	}
	obj.HeroCache.SetSkillList(skill_map)

	obj.scheme_base_data = scheme_base_data
	obj.scheme_lvup_data = scheme_lvup_data
	obj.scheme_stage_data = scheme_stage_data
	obj.scheme_rank_data = scheme_rank_data

	return obj, nil
}

func LoadHero(buf []byte) (*Hero, error) {
	obj := new(Hero)
	err := proto.Unmarshal(buf, &obj.HeroCache)
	if err != nil {
		return nil, err
	}

	scheme_base_data, has := scheme.Heromap[obj.GetSchemeId()]
	if !has {
		return nil, errors.New(fmt.Sprintf("LoadHero Base Scheme Error (scheme_id : %v)", obj.GetSchemeId()))
	}
	scheme_lvup_data := scheme.HeroLvUpGet(obj.GetSchemeId(), obj.GetLv(), obj.GetRank())
	if scheme_lvup_data == nil {
		return nil, errors.New(fmt.Sprintf("LoadHero LvUp Scheme Error (scheme_id : %v) (lv : %v) (stage : %v) (rank : %v)", obj.GetSchemeId(), obj.GetLv(), obj.GetStage(), obj.GetRank()))
	}
	scheme_stage_data := scheme.HeroStageUpGet(obj.GetSchemeId(), obj.GetStage(), obj.GetRank())
	if scheme_stage_data == nil {
		return nil, errors.New(fmt.Sprintf("LoadHero StageUp Scheme Error (scheme_id : %v) (stage : %v) (rank : %v)", obj.GetSchemeId(), obj.GetStage(), obj.GetRank()))
	}
	scheme_rank_data := scheme.HeroRankGet(scheme_base_data.MagicHeroRankId, obj.GetRank())
	if scheme_rank_data == nil {
		return nil, errors.New(fmt.Sprintf("LoadHero Rank Scheme Error (MagicHeroRankId : %v) (rank : %v)", scheme_base_data.MagicHeroRankId, obj.GetRank()))
	}

	obj.scheme_base_data = scheme_base_data
	obj.scheme_lvup_data = scheme_lvup_data
	obj.scheme_stage_data = scheme_stage_data
	obj.scheme_rank_data = scheme_rank_data
	return obj, nil
}

func (this *Hero) GetPopulation() int32 {
	return this.scheme_base_data.Population
}

func (this *Hero) GetRankPoint() int32 {
	return this.scheme_rank_data.BaseRankPoint
}

func (this *Hero) GetAllExp() int32 {
	return this.scheme_lvup_data.ExpCount + this.HeroCache.GetLvExp()
}

func (this *Hero) FillHeroInfo() *protocol.HeroInfo {
	msg := new(protocol.HeroInfo)
	msg.SetUid(this.HeroCache.GetUid())
	msg.SetSchemeId(this.HeroCache.GetSchemeId())
	msg.SetLv(this.HeroCache.GetLv())
	msg.SetLvExp(this.HeroCache.GetLvExp())
	msg.SetStage(this.HeroCache.GetStage())
	msg.SetStageTimestamp(this.HeroCache.GetStageTimestamp())
	msg.SetStageSpeedup(this.HeroCache.GetStageSpeedup())
	msg.SetRank(this.HeroCache.GetRank())
	msg.SetRankExp(this.HeroCache.GetRankExp())
	msg.SkillList = make([]*protocol.HeroSkillInfo, len(this.HeroCache.SkillList))
	index := 0
	for _, v := range this.HeroCache.SkillList {
		element := &protocol.HeroSkillInfo{
			SkillId: proto.Int32(v.GetSkillId()),
			SkillLv: proto.Int32(v.GetSkillLv()),
		}
		msg.SkillList[index] = element
		index++
	}

	return msg
}

func (this *Hero) AddLvExp(value int32, kinglv int32) common.RetCode {
	if value <= 0 {
		return common.RetCode_Fail
	}

	max_lv := scheme.Commonmap[define.HeroLvMax].Value
	if this.HeroCache.GetLv() >= max_lv {
		return common.RetCode_Unable
	}

	if this.GetLv() > this.scheme_stage_data.LvLimit {
		return common.RetCode_Unable
	}

	if this.GetLv() == this.scheme_stage_data.LvLimit && this.GetLvExp() >= this.scheme_lvup_data.NeedExp {
		return common.RetCode_Unable
	}

	if kinglv < this.scheme_lvup_data.LvUpKingLv && this.GetLvExp() >= this.scheme_lvup_data.NeedExp {
		return common.RetCode_Unable
	}

	total_exp := this.GetLvExp() + value
	for total_exp > 0 {
		total_exp -= this.scheme_lvup_data.NeedExp
		if total_exp >= 0 {
			if kinglv < this.scheme_lvup_data.LvUpKingLv {
				this.HeroCache.SetLvExp(this.scheme_lvup_data.NeedExp)
				break
			}

			if this.GetLv() == this.scheme_stage_data.LvLimit {
				this.HeroCache.SetLvExp(this.scheme_lvup_data.NeedExp)
				break
			}

			scheme_lvup_data := scheme.HeroLvUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetLv()+1, this.HeroCache.GetRank())
			if scheme_lvup_data == nil {
				return common.RetCode_SchemeData_Error
			}
			this.scheme_lvup_data = scheme_lvup_data
			this.HeroCache.SetLv(this.HeroCache.GetLv() + 1)

			if this.HeroCache.GetLv() >= max_lv {
				this.HeroCache.SetLvExp(0)
				break
			}

			if total_exp == 0 {
				this.HeroCache.SetLvExp(total_exp)
			}
		} else {
			this.HeroCache.SetLvExp(total_exp + this.scheme_lvup_data.NeedExp)
		}
	}

	return common.RetCode_Success
}

func (this *Hero) freshLv(kingLv int32) common.RetCode {
	if kingLv >= this.scheme_lvup_data.LvUpKingLv && this.HeroCache.GetLvExp() >= this.scheme_lvup_data.NeedExp {
		if this.GetLv() == this.scheme_stage_data.LvLimit || this.GetLv() >= scheme.Commonmap[define.HeroLvMax].Value {
			return common.RetCode_Unable
		}

		scheme_lvup_data := scheme.HeroLvUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetLv()+1, this.HeroCache.GetRank())
		if scheme_lvup_data == nil {
			return common.RetCode_SchemeData_Error
		}
		this.scheme_lvup_data = scheme_lvup_data
		this.HeroCache.SetLv(this.HeroCache.GetLv() + 1)
		this.HeroCache.SetLvExp(0)
		return common.RetCode_Success
	}
	return common.RetCode_Unable
}

func (this *Hero) AddSkillLv(skill_id int32, owner IRole) common.RetCode {
	skill, has := this.HeroCache.GetSkillList()[skill_id]
	if !has {
		return common.RetCode_Fail
	}

	if skill.GetSkillLv() >= scheme.Commonmap[define.HeroSkillLvMax].Value || skill.GetSkillLv() >= this.GetLv() {
		return common.RetCode_Unable
	}

	skill_scheme := scheme.SkillLvUpGet(skill.GetSkillId(), skill.GetSkillLv())
	if skill_scheme == nil {
		return common.RetCode_SchemeData_Error
	}

	if !owner.ItemIsEnough(skill_scheme.NeedItemId, skill_scheme.NeedItemNum) {
		return common.RetCode_Unable
	}

	skill.SetSkillLv(skill.GetSkillLv() + 1)
	owner.ItemCost(skill_scheme.NeedItemId, skill_scheme.NeedItemNum, true)
	return common.RetCode_Success
}

func (this *Hero) freshSkill() {
	skill_list := strings.Split(this.scheme_stage_data.SkillId, ";")
	for _, skill_data := range skill_list {
		if skill_id, err := strconv.Atoi(skill_data); err == nil {
			if skill_id != -1 {
				if _, has := this.HeroCache.SkillList[int32(skill_id)]; !has {
					skill_cache := new(HeroSkillCache)
					skill_cache.SetSkillId(int32(skill_id))
					skill_cache.SetSkillLv(1)
					this.HeroCache.SkillList[int32(skill_id)] = skill_cache
				}
			}
		} else {
			LogError(fmt.Sprintf("NewHero Skill_id(%v) Scheme Error", skill_data))
		}
	}
}

func (this *Hero) AddRankExp(value int32) common.RetCode {
	if value <= 0 {
		return common.RetCode_Unable
	}

	max_rank := scheme.Commonmap[define.HeroRankLimit].Value
	if this.HeroCache.GetRank() >= max_rank {
		return common.RetCode_Success
	}

	total_exp := this.HeroCache.GetRankExp() + value
	for total_exp > 0 {
		total_exp -= this.scheme_rank_data.NeedRankPoint
		if total_exp >= 0 {
			scheme_lv_data := scheme.HeroLvUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetLv(), this.HeroCache.GetRank()+1)
			if scheme_lv_data == nil {
				return common.RetCode_SchemeData_Error
			}

			scheme_stage_data := scheme.HeroStageUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetStage(), this.HeroCache.GetRank()+1)
			if scheme_stage_data == nil {
				return common.RetCode_SchemeData_Error
			}

			scheme_rank_data := scheme.HeroRankGet(this.scheme_base_data.MagicHeroRankId, this.HeroCache.GetRank()+1)
			if scheme_rank_data == nil {
				return common.RetCode_SchemeData_Error
			}

			this.scheme_lvup_data = scheme_lv_data
			this.scheme_stage_data = scheme_stage_data
			this.scheme_rank_data = scheme_rank_data

			this.HeroCache.SetRank(this.HeroCache.GetRank() + 1)
			if this.HeroCache.GetRank() >= max_rank {
				return common.RetCode_Success
			}
			if total_exp == 0 {
				this.HeroCache.SetRankExp(total_exp)
			}
		} else {
			this.HeroCache.SetRankExp(total_exp + this.scheme_rank_data.NeedRankPoint)
		}
	}

	return common.RetCode_Success
}

func (this *Hero) EditLv(lv int32) common.RetCode {
	scheme_lvup_data := scheme.HeroLvUpGet(this.HeroCache.GetSchemeId(), lv, this.HeroCache.GetRank())
	if scheme_lvup_data == nil {
		return common.RetCode_SchemeData_Error
	}
	scheme_stage_data := scheme.HeroStageUpGet(this.HeroCache.GetSchemeId(), scheme_lvup_data.Stage, this.HeroCache.GetRank())
	if scheme_stage_data == nil {
		return common.RetCode_SchemeData_Error
	}

	this.scheme_lvup_data = scheme_lvup_data
	this.scheme_stage_data = scheme_stage_data

	if this.HeroCache.GetLvExp() > this.scheme_lvup_data.NeedExp {
		this.HeroCache.SetLvExp(this.HeroCache.GetLvExp() - 1)
	}

	if this.HeroCache.GetLv() >= scheme.Commonmap[define.HeroLvMax].Value {
		this.HeroCache.SetLvExp(0)
	}

	this.HeroCache.SetLv(lv)
	this.HeroCache.SetStage(scheme_lvup_data.Stage)
	this.HeroCache.SetStageTimestamp(0)
	this.HeroCache.SetStageSpeedup(0)

	return common.RetCode_Success
}

func (this *Hero) EditRank(rank int32) common.RetCode {
	scheme_lv_data := scheme.HeroLvUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetLv(), rank)
	if scheme_lv_data == nil {
		return common.RetCode_SchemeData_Error
	}

	scheme_stage_data := scheme.HeroStageUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetStage(), rank)
	if scheme_stage_data == nil {
		return common.RetCode_SchemeData_Error
	}

	scheme_rank_data := scheme.HeroRankGet(this.scheme_base_data.MagicHeroRankId, rank)
	if scheme_rank_data == nil {
		return common.RetCode_SchemeData_Error
	}

	this.scheme_lvup_data = scheme_lv_data
	this.scheme_stage_data = scheme_stage_data
	this.scheme_rank_data = scheme_rank_data

	this.HeroCache.SetRank(rank)
	if this.HeroCache.GetRank() >= scheme.Commonmap[define.HeroRankLimit].Value {
		this.HeroCache.SetRankExp(0)
	}

	return common.RetCode_Success
}

func (this *Hero) EditSkillLv(skill_id int32, lv int32) common.RetCode {
	skill, has := this.HeroCache.GetSkillList()[skill_id]
	if !has {
		return common.RetCode_Fail
	}

	if lv >= scheme.Commonmap[define.HeroSkillLvMax].Value {
		return common.RetCode_Unable
	}

	skill_scheme := scheme.SkillLvUpGet(skill.GetSkillId(), lv)
	if skill_scheme == nil {
		return common.RetCode_SchemeData_Error
	}

	skill.SetSkillLv(lv)
	return common.RetCode_Success
}

type HeroSys struct {
	owner                 IRole
	hero_list             map[int64]*Hero
	hero_create_slot      *HeroCreate
	cache_list_key        string
	cache_create_slot_key string
}

func (this *HeroSys) Init(owner IRole) {
	this.owner = owner
	this.hero_list = make(map[int64]*Hero)
	this.hero_create_slot = new(HeroCreate)
	this.cache_list_key = fmt.Sprintf(cache_herolist_key_t, this.owner.GetUid())
	this.cache_create_slot_key = fmt.Sprintf(cache_herocreate_key_t, this.owner.GetUid())
}

func (this *HeroSys) Load() error {
	resp, err := GxService().Redis().Cmd("SMEMBERS", this.cache_list_key)
	if err != nil {
		return err
	}

	cacheKeys, err := resp.List()
	if err != nil {
		return err
	}

	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", this.cache_list_key, key)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			hero, err := LoadHero(buf)
			if err != nil {
				LogFatal(err)
				return err
			}
			this.hero_list[hero.GetUid()] = hero
		}
	}

	resp, err = GxService().Redis().Cmd("GET", this.cache_create_slot_key)
	if err != nil {
		return err
	}

	if buf, _ := resp.Bytes(); buf != nil {
		hero_create, err := LoadHeroCreate(buf)
		if err != nil {
			return err
		}
		this.hero_create_slot = hero_create
	}

	return nil
}

func (this *HeroSys) Save(hero *Hero) {
	buf, err := proto.Marshal(&hero.HeroCache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := GenHeroCacheKey(this.owner.GetUid(), hero.HeroCache.GetUid())
	if _, err := GxService().Redis().Cmd("SET", key, buf); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SADD", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
}

func (this *HeroSys) SaveHeroCreate() {
	buf, err := proto.Marshal(&this.hero_create_slot.HeroCreateCache)
	if err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SET", this.cache_create_slot_key, buf); err != nil {
		LogFatal(err)
		return
	}
}

func (this *HeroSys) Del(hero_uid int64) {
	key := GenHeroCacheKey(this.owner.GetUid(), hero_uid)
	if _, err := GxService().Redis().Cmd("DEL", key); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SREM", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
	delete(this.hero_list, hero_uid)
}

func (this *HeroSys) FillHeroCreateInfo() *protocol.HeroCreateInfo {
	return this.hero_create_slot.FillHeroCreateInfo()
}

func (this *HeroSys) FillHeroListInfo() *protocol.HeroListInfo {
	msg := new(protocol.HeroListInfo)
	msg.HeroList = make([]*protocol.HeroInfo, len(this.hero_list))
	index := 0
	for _, v := range this.hero_list {
		msg.HeroList[index] = v.FillHeroInfo()
		index++
	}
	return msg
}

func (this *HeroSys) FillHeroFightInfo(hero_uid int64) *protocol.HeroFightInfo {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return nil
	}

	msg := new(protocol.HeroFightInfo)
	msg.SetUid(hero.GetUid())
	msg.SetSchemeId(hero.GetSchemeId())
	msg.SetLv(hero.GetLv())
	msg.SetStage(hero.GetStage())
	msg.SetRank(hero.GetRank())
	msg.SkillList = make([]*protocol.HeroSkillInfo, len(hero.HeroCache.SkillList))
	index := 0
	for _, v := range hero.HeroCache.SkillList {
		element := &protocol.HeroSkillInfo{
			SkillId: proto.Int32(v.GetSkillId()),
			SkillLv: proto.Int32(v.GetSkillLv()),
		}
		msg.SkillList[index] = element
		index++
	}
	return msg
}

func (this *HeroSys) HeroFreshLv(kingLv int32) {
	for _, hero := range this.hero_list {
		if hero.freshLv(kingLv) == common.RetCode_Success {
			this.static_hero(hero)
			this.Save(hero)
			this.send_info_notify(hero)
		}
	}
}

func (this *HeroSys) HeroGet(hero_uid int64) IHero {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return nil
	}
	return hero
}

func (this *HeroSys) HeroCost(hero_uid int64, is_notify bool) {
	this.Del(hero_uid)
	this.static_hero_del(hero_uid)
	if is_notify {
		this.send_lose_notify(hero_uid)
	}
}

func (this *HeroSys) HeroCreateStart(create_id int32, is_notify bool) common.RetCode {
	if this.hero_create_slot.IsCD() {
		return common.RetCode_Unable
	}

	plan_id := scheme.HeroCreateRandPlanId()
	err := this.hero_create_slot.Start(create_id, plan_id, time.Now().Unix())
	if err != nil {
		return common.RetCode_Fail
	}
	this.SaveHeroCreate()
	if is_notify {
		this.send_create_notify()
	}

	return common.RetCode_Success
}

func (this *HeroSys) HeroCreateFinish(is_notify bool) common.RetCode {
	if this.hero_create_slot.GetCreateId() == 0 {
		return common.RetCode_Unable
	}

	if this.hero_create_slot.GetHeroCreateDropId() == 0 {
		return common.RetCode_Fail
	}

	if this.hero_create_slot.IsDie() {
		return common.RetCode_Unable
	}

	if !this.hero_create_slot.IsFinish() {
		return common.RetCode_Unable
	}

	//通过Award 获得
	_, ret := award.Award(this.hero_create_slot.GetHeroCreateDropId(), this.owner, true)
	if ret != common.RetCode_Success {
		return ret
	}

	this.hero_create_slot.Reset()
	this.SaveHeroCreate()
	if is_notify {
		this.send_create_notify()
	}

	return common.RetCode_Success
}

func (this *HeroSys) HeroCreateAddOrder(order int32, is_notify bool) bool {
	if this.hero_create_slot.AddHeroCreateOrder(order) {
		this.SaveHeroCreate()
		if is_notify {
			this.send_create_notify()
		}
		return true
	}
	return false
}

func (this *HeroSys) HeroCreateAddMagic(magic int32, is_notify bool) bool {
	if this.hero_create_slot.AddHeroCreateMagic(magic) {
		this.SaveHeroCreate()
		if is_notify {
			this.send_create_notify()
		}
		return true
	}
	return false
}

func (this *HeroSys) HeroCreateShock(is_notify bool) common.RetCode {
	if !this.owner.IsEnoughGold(scheme.Commonmap[define.MagicHeroCreateShockGold].Value) {
		return common.RetCode_Unable
	}

	if err := this.hero_create_slot.ReBorn(); err == nil {
		this.owner.CostGold(scheme.Commonmap[define.MagicHeroCreateShockGold].Value, true, true)
		this.SaveHeroCreate()
		if is_notify {
			this.send_create_notify()
		}
		return common.RetCode_Success
	} else {
		LogError(err)
	}
	return common.RetCode_Fail
}

func (this *HeroSys) HeroCreateGiveUp(is_notify bool) common.RetCode {
	if this.hero_create_slot.GetCreateId() == 0 && !this.hero_create_slot.IsDie() {
		return common.RetCode_Unable
	}

	this.hero_create_slot.Reset()
	this.SaveHeroCreate()
	if is_notify {
		this.send_create_notify()
	}
	return common.RetCode_Success
}

func (this *HeroSys) HeroObtain(hero_scheme_id int32, hero_lv int32, hero_rank int32, is_notify bool) int64 {
	LogDebug("HeroObtain Enter : hero_scheme_id(", hero_scheme_id, ") hero_lv(", hero_lv, ") hero_rank(", hero_rank, ")")
	hero, err := NewHero(hero_scheme_id, hero_lv, hero_rank, this.owner.GetUid())
	if err != nil {
		LogError(err)
		return common.UID_FAILED
	}
	this.hero_list[hero.GetUid()] = hero
	this.static_hero(hero)
	this.Save(hero)
	LogDebug("HeroObtain Success : role_uid(", this.owner.GetUid(), ") hero_scheme_id(", hero_scheme_id, ")")

	if is_notify {
		this.send_info_notify(hero)
	}

	//A =4   15
	//s =5   16
	if hero.scheme_rank_data.Rank == 4 {
		this.owner.AchievementAddNum(15, 1, false)
	} else if hero.scheme_rank_data.Rank == 5 {
		this.owner.AchievementAddNum(16, 1, false)
	}

	//添加成就
	this.owner.AchievementAddNum(13, 1, false)

	return hero.GetUid()
}

func (this *HeroSys) HeroLvUp(hero_uid int64, add_exp int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	if hero.HeroCache.GetStageTimestamp() != 0 {
		return common.RetCode_Unable
	}

	old_lv := hero.GetLv()
	ret := hero.AddLvExp(add_exp, this.owner.GetKingLv())
	if ret == common.RetCode_Success {
		if hero.GetLv() > old_lv {
			this.static_hero(hero)
		}

		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		this.owner.AchievementAddNum(14, hero.GetLv(), true)
	}
	return ret
}

func (this *HeroSys) HeroSkillLvUp(hero_uid int64, skill_id int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	if hero.HeroCache.GetStageTimestamp() != 0 {
		return common.RetCode_Unable
	}

	ret := hero.AddSkillLv(skill_id, this.owner)
	if ret == common.RetCode_Success {
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}
	}
	return ret
}

func (this *HeroSys) HeroEvoStart(hero_uid int64, need_hero_uids []int64, use_money bool, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	if !use_money {
		if _, id := this.owner.GetMagicQueue(); id != 0 {
			return common.RetCode_Unable
		}
	}

	if hero.HeroCache.GetLv() < hero.scheme_stage_data.LvLimit-1 && hero.HeroCache.GetLvExp() >= hero.scheme_lvup_data.NeedExp-1 {
		return common.RetCode_Unable
	}

	if hero.scheme_stage_data.NextStageId == -1 {
		return common.RetCode_Unable
	}

	if hero.HeroCache.GetStageTimestamp() != 0 {
		return common.RetCode_Unable
	}

	scheme_lvup_data := scheme.HeroLvUpGet(hero.GetSchemeId(), hero.GetLv()+1, hero.GetRank())
	if scheme_lvup_data == nil {
		return common.RetCode_SchemeData_Error
	}

	scheme_stage_data := scheme.HeroStageUpGet(hero.GetSchemeId(), hero.GetStage()+1, hero.GetRank())
	if scheme_stage_data == nil {
		return common.RetCode_SchemeData_Error
	}

	need_item_id := strings.Split(hero.scheme_stage_data.EvoNeedItemId, ";")
	need_item_num := strings.Split(hero.scheme_stage_data.EvoNeedItemNum, ";")
	if len(need_item_id) != len(need_item_num) {
		return common.RetCode_SchemeData_Error
	}

	list := make(map[int32]int32)
	for i, v := range need_item_id {
		item_id, _ := strconv.Atoi(v)
		item_num, _ := strconv.Atoi(need_item_num[i])
		if !this.owner.ItemIsEnough(int32(item_id), int32(item_num)) {
			return common.RetCode_Unable
		}
		list[int32(item_id)] = int32(item_num)
	}

	if need_hero_uids == nil || int32(len(need_hero_uids)) != hero.scheme_stage_data.EvoNeedMagicHeroNum {
		return common.RetCode_Unable
	}

	for _, uid := range need_hero_uids {
		if hero.GetUid() == uid {
			return common.RetCode_Unable
		}

		need_hero, has := this.hero_list[uid]
		if !has {
			return common.RetCode_Unable
		}

		if need_hero.GetSchemeId() != hero.scheme_stage_data.EvoNeedMagicHeroId || need_hero.GetRank() < hero.scheme_stage_data.EvoNeedMagicHeroRank || need_hero.GetLv() < hero.scheme_stage_data.EvoNeedMagicHeroLv {
			return common.RetCode_Unable
		}
	}

	for id, num := range list {
		this.owner.ItemCost(id, num, true)
	}

	for _, uid := range need_hero_uids {
		this.HeroCost(uid, true)
	}

	if use_money {
		gold := ResourceToCoin(common.RTYPE_TIME, hero.scheme_stage_data.EvoNeedTime)
		if !this.owner.IsEnoughGold(gold) {
			return common.RetCode_Unable
		}
		this.owner.CostGold(gold, true, true)
		hero.HeroCache.SetStage(hero.GetStage() + 1)
		hero.HeroCache.SetLv(hero.HeroCache.GetLv() + 1)
		hero.HeroCache.SetLvExp(0)
		this.static_hero(hero)
		this.owner.StaticPayLog(int32(static.PayType_evolutionOnekey), 0, gold)
		hero.scheme_lvup_data = scheme_lvup_data
		hero.scheme_stage_data = scheme_stage_data
	} else {
		hero.HeroCache.SetStageTimestamp(time.Now().Unix() + int64(hero.scheme_stage_data.EvoNeedTime))
		hero.HeroCache.SetStageSpeedup(time.Now().Unix())
		this.owner.SetMagicQueue(common.RTYPE_MAGIC_HERO, hero.GetUid(), true)
	}
	this.Save(hero)
	if is_notify {
		this.send_info_notify(hero)
	}
	return common.RetCode_Success
}

func (this *HeroSys) HeroEvoFinish(use_money bool, is_notify bool) common.RetCode {
	rtype, id := this.owner.GetMagicQueue()
	if rtype != common.RTYPE_MAGIC_HERO {
		return common.RetCode_Unable
	}

	hero, has := this.hero_list[id]
	if !has {
		return common.RetCode_Unable
	}

	if hero.HeroCache.GetStageTimestamp() == 0 {
		return common.RetCode_Unable
	}

	scheme_lvup_data := scheme.HeroLvUpGet(hero.GetSchemeId(), hero.GetLv()+1, hero.GetRank())
	if scheme_lvup_data == nil {
		return common.RetCode_SchemeData_Error
	}

	scheme_stage_data := scheme.HeroStageUpGet(hero.GetSchemeId(), hero.GetStage()+1, hero.GetRank())
	if scheme_stage_data == nil {
		return common.RetCode_SchemeData_Error
	}

	if time.Now().Unix() < hero.GetStageTimestamp() && !use_money {
		return common.RetCode_Unable
	}

	if use_money {
		delta := hero.HeroCache.GetStageTimestamp() - time.Now().Unix()
		gold := ResourceToCoin(common.RTYPE_TIME, int32(delta))
		if !this.owner.IsEnoughGold(gold) {
			return common.RetCode_Unable
		}
		this.owner.CostGold(gold, true, true)
		this.owner.StaticPayLog(int32(static.PayType_evolutionSpeedup), 0, gold)
	}

	hero.HeroCache.SetStage(hero.GetStage() + 1)
	hero.HeroCache.SetLv(hero.HeroCache.GetLv() + 1)
	hero.HeroCache.SetLvExp(0)
	this.static_hero(hero)
	hero.scheme_lvup_data = scheme_lvup_data
	hero.scheme_stage_data = scheme_stage_data

	hero.HeroCache.SetStageTimestamp(0)
	hero.HeroCache.SetStageSpeedup(0)
	this.owner.ResetMagicQueue(true)
	this.Save(hero)
	if is_notify {
		this.send_info_notify(hero)
	}
	return common.RetCode_Success
}

func (this *HeroSys) HeroEvoSpeedUp(is_notify bool) common.RetCode {
	rtype, id := this.owner.GetMagicQueue()
	if rtype != common.RTYPE_MAGIC_HERO {
		return common.RetCode_Unable
	}

	hero, has := this.hero_list[id]
	if !has {
		return common.RetCode_Unable
	}

	if hero.HeroCache.GetStageTimestamp() == 0 {
		return common.RetCode_Unable
	}

	now := time.Now().Unix()
	count := (now - hero.HeroCache.GetStageSpeedup()) / int64(scheme.Commonmap[define.EvoSpeedRecover].Value)
	if count > int64(this.owner.GetEvoSpeedLimit()) {
		count = int64(this.owner.GetEvoSpeedLimit())
		hero.HeroCache.SetStageSpeedup(now - int64(scheme.Commonmap[define.EvoSpeedRecover].Value)*count)
	}
	if count >= 1 {
		fix := hero.HeroCache.GetStageSpeedup() + int64(scheme.Commonmap[define.EvoSpeedRecover].Value)
		after := hero.HeroCache.GetStageTimestamp() - int64(scheme.Commonmap[define.EvoSpeedValue].Value)
		hero.HeroCache.SetStageTimestamp(after)
		hero.HeroCache.SetStageSpeedup(fix)
	} else {
		return common.RetCode_CoolDown
	}

	this.Save(hero)
	if is_notify {
		this.send_info_notify(hero)
	}
	return common.RetCode_Success
}

func (this *HeroSys) HeroMix(hero_uid int64, hero_uids []int64, is_notify bool) common.RetCode {
	if hero_uids == nil {
		return common.RetCode_Fail
	}

	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	if hero.GetStageTimestamp() != 0 {
		return common.RetCode_Unable
	}

	heros := make([]*Hero, len(hero_uids))
	for index, v := range hero_uids {
		if hero.GetUid() == v {
			return common.RetCode_Unable
		}

		if this.owner.MapFindHero(v) {
			return common.RetCode_Unable
		}

		to_hero, has := this.hero_list[v]
		if !has {
			return common.RetCode_Unable
		}

		if hero.HeroCache.GetSchemeId() != to_hero.HeroCache.GetSchemeId() || to_hero.HeroCache.GetStageTimestamp() != 0 {
			return common.RetCode_Unable
		}
		heros[index] = to_hero
	}

	for _, v := range heros {
		hero.AddLvExp(v.GetAllExp(), this.owner.GetKingLv())
		ret := hero.AddRankExp(v.GetRankPoint())
		if ret != common.RetCode_Success {
			return ret
		}

		hero.freshSkill()
		hero_skill := hero.GetSkillList()
		to_hero_skill := v.GetSkillList()
		for _, skill := range hero_skill {
			if temp, has := to_hero_skill[skill.GetSkillId()]; has {
				if temp.GetSkillLv() > skill.GetSkillLv() {
					skill.SetSkillLv(temp.GetSkillLv())
				}
			}
		}
		this.HeroCost(v.HeroCache.GetUid(), true)
	}
	this.static_hero(hero)
	this.Save(hero)
	if is_notify {
		this.send_info_notify(hero)
	}

	//A =4   15
	//s =5   16
	if hero.scheme_rank_data.Rank == 4 {
		this.owner.AchievementAddNum(15, 1, false)
	} else if hero.scheme_rank_data.Rank == 5 {
		this.owner.AchievementAddNum(16, 1, false)
	}

	return common.RetCode_Success
}

func (this *HeroSys) HeroSize() int32 {
	return int32(len(this.hero_list))
}

func (this *HeroSys) HeroFind(hero_uid int64) bool {
	_, has := this.hero_list[hero_uid]
	return has
}

func (this *HeroSys) HeroPopulation(hero_uid int64) int32 {
	hero, has := this.hero_list[hero_uid]
	if has {
		return hero.GetPopulation()
	}
	return 0
}

func (this *HeroSys) HeroEditLv(hero_uid int64, lv int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	ret := hero.EditLv(lv)
	if ret == common.RetCode_Success {
		if magic_type, magic_id := this.owner.GetMagicQueue(); magic_type == common.RTYPE_MAGIC_HERO && magic_id == hero_uid {
			this.owner.ResetMagicQueue(true)
		}

		this.static_hero(hero)
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		this.owner.AchievementAddNum(14, hero.GetLv(), true)
	}
	return ret
}

func (this *HeroSys) HeroEditRank(hero_uid int64, rank int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	ret := hero.EditRank(rank)
	if ret == common.RetCode_Success {
		hero.freshSkill()
		this.static_hero(hero)
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		if hero.scheme_rank_data.Rank == 4 {
			this.owner.AchievementAddNum(15, 1, false)
		} else if hero.scheme_rank_data.Rank == 5 {
			this.owner.AchievementAddNum(16, 1, false)
		}
	}
	return ret
}

func (this *HeroSys) HeroEditSkillLv(hero_uid int64, skill_id int32, lv int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_uid]
	if !has {
		return common.RetCode_Unable
	}

	ret := hero.EditSkillLv(skill_id, lv)
	if ret == common.RetCode_Success {
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}
	}
	return ret
}

func (this *HeroSys) send_info_notify(hero *Hero) {
	msg := &protocol.MsgHeroInfoNotify{}
	msg.Hero = hero.FillHeroInfo()

	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	global.SendMsg(int32(protocol.MsgCode_HeroInfoNotify), this.owner.GetSid(), buf)
}

func (this *HeroSys) send_lose_notify(uid int64) {
	msg := &protocol.MsgHeroLoseNotify{
		Uid: proto.Int64(uid),
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	global.SendMsg(int32(protocol.MsgCode_HeroLoseNotify), this.owner.GetSid(), buf)
}

func (this *HeroSys) send_create_notify() {
	msg := new(protocol.MsgHeroCreateNotify)
	msg.Infos = this.hero_create_slot.FillHeroCreateInfo()
	fmt.Println(msg)
	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	global.SendMsg(int32(protocol.MsgCode_HeroCreateNotify), this.owner.GetSid(), buf)
}

func (this *HeroSys) static_hero(hero *Hero) {
	msg := &static.MsgStaticHero{}
	msg.SetRoleUid(this.owner.GetUid())
	msg.SetUid(hero.GetUid())
	msg.SetSchemeId(hero.GetSchemeId())
	msg.SetLv(hero.GetLv())
	msg.SetStage(hero.GetStage())
	msg.SetRank(hero.GetRank())

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_Hero), buf)
}

func (this *HeroSys) static_hero_del(hero_uid int64) {
	msg := &static.MsgStaticHeroDel{}
	msg.SetRoleUid(this.owner.GetUid())
	msg.SetUid(hero_uid)

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_HeroDel), buf)
}
