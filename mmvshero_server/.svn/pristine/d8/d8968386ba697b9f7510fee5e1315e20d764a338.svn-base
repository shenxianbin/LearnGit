package hero

import (
	"Gameserver/global"
	. "Gameserver/logic"
	"common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"errors"
	"fmt"
	. "galaxy"
	"math/rand"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	cache_herolist_key_t = "Role:%v:Hero"
	cache_heroobj_key_t  = "Role:%v:Hero:%v"
)

func GenHeroListKey(role_uid int64) string {
	return fmt.Sprintf(cache_herolist_key_t, role_uid)
}

func GenHeroCacheKey(role_uid int64, hero_id int32) string {
	return fmt.Sprintf(cache_heroobj_key_t, role_uid, hero_id)
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
	obj.HeroCache.SetSchemeId(scheme_id)
	obj.HeroCache.SetLv(lv)
	obj.HeroCache.SetStage(scheme_lvup_data.Stage)
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

func (this *Hero) FillHeroInfo() *protocol.HeroInfo {
	msg := new(protocol.HeroInfo)
	msg.SetSchemeId(this.HeroCache.GetSchemeId())
	msg.SetLv(this.HeroCache.GetLv())
	msg.SetStage(this.HeroCache.GetStage())
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

func (this *Hero) AddLvExp(value int32, role IRole) common.RetCode {
	if value <= 0 {
		return common.RetCode_HeroExpArgsError
	}

	max_lv := scheme.Commonmap[define.HeroLvMax].Value
	if this.HeroCache.GetLv() >= max_lv {
		return common.RetCode_HeroLvMax
	}

	if role.GetLv() < this.scheme_lvup_data.LvUpRoleLv {
		return common.RetCode_HeroLvLimitByRole
	}

	if value < this.scheme_lvup_data.NeedExp {
		return common.RetCode_Failed
	}

	scheme_lvup_data := scheme.HeroLvUpGet(this.HeroCache.GetSchemeId(), this.HeroCache.GetLv()+1, this.HeroCache.GetRank())
	if scheme_lvup_data == nil {
		return common.RetCode_SchemeData_Error
	}

	scheme_stage_data := scheme.HeroStageUpGet(this.HeroCache.GetSchemeId(), scheme_lvup_data.Stage, this.HeroCache.GetRank())
	if scheme_stage_data == nil {
		return common.RetCode_SchemeData_Error
	}

	this.scheme_lvup_data = scheme_lvup_data
	this.scheme_stage_data = scheme_stage_data
	this.HeroCache.SetLv(this.HeroCache.GetLv() + 1)
	this.HeroCache.SetStage(scheme_lvup_data.Stage)
	role.AddExp(this.scheme_lvup_data.AddRoleExp, true, true)

	return common.RetCode_Success
}

func (this *Hero) AddSkillLv(skill_id int32, owner IRole) common.RetCode {
	skill, has := this.HeroCache.GetSkillList()[skill_id]
	if !has {
		return common.RetCode_HeroSkillIdError
	}

	if skill.GetSkillLv() >= scheme.Commonmap[define.HeroSkillLvMax].Value || skill.GetSkillLv() >= this.GetLv() {
		return common.RetCode_HeroSkillLvLimit
	}

	skill_scheme := scheme.SkillLvUpGet(skill.GetSkillId(), skill.GetSkillLv())
	if skill_scheme == nil {
		return common.RetCode_SchemeData_Error
	}

	if !owner.IsEnoughSoul(skill_scheme.NeedSoul) {
		return common.RetCode_RoleNotEnoughSoul
	}

	skill.SetSkillLv(skill.GetSkillLv() + 1)
	owner.CostSoul(skill_scheme.NeedSoul, true, true)
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

func (this *Hero) freshSkillGM() {
	skill_list := strings.Split(this.scheme_stage_data.SkillId, ";")
	scheme_skill := make(map[int32]bool)
	for _, skill_data := range skill_list {
		if skill_id, err := strconv.Atoi(skill_data); err == nil {
			scheme_skill[int32(skill_id)] = true
		} else {
			LogError(fmt.Sprintf("NewHero Skill_id(%v) Scheme Error", skill_data))
		}
	}

	for id, _ := range this.HeroCache.SkillList {
		if _, has := scheme_skill[id]; !has {
			delete(this.HeroCache.SkillList, id)
		}
	}

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
		}
	}
}

func (this *Hero) AddRankExp(value int32) common.RetCode {
	if value <= 0 {
		return common.RetCode_HeroExpArgsError
	}

	max_rank := scheme.Commonmap[define.HeroRankLimit].Value
	if this.HeroCache.GetRank() >= max_rank {
		return common.RetCode_HeroRankMax
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

	this.HeroCache.SetLv(lv)
	this.HeroCache.SetStage(scheme_lvup_data.Stage)

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
		return common.RetCode_HeroSkillIdError
	}

	if lv >= scheme.Commonmap[define.HeroSkillLvMax].Value {
		return common.RetCode_HeroSkillLvLimit
	}

	skill_scheme := scheme.SkillLvUpGet(skill.GetSkillId(), lv)
	if skill_scheme == nil {
		return common.RetCode_SchemeData_Error
	}

	skill.SetSkillLv(lv)
	return common.RetCode_Success
}

type HeroSys struct {
	owner          IRole
	hero_list      map[int32]*Hero
	cache_list_key string
}

func (this *HeroSys) Init(owner IRole) {
	this.owner = owner
	this.hero_list = make(map[int32]*Hero)
	this.cache_list_key = fmt.Sprintf(cache_herolist_key_t, this.owner.GetUid())
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
			this.hero_list[hero.GetSchemeId()] = hero
		}
	}

	return nil
}

func (this *HeroSys) Save(hero *Hero) {
	buf, err := proto.Marshal(&hero.HeroCache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := GenHeroCacheKey(this.owner.GetUid(), hero.HeroCache.GetSchemeId())
	if _, err := GxService().Redis().Cmd("SET", key, buf); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SADD", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
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

func (this *HeroSys) HeroGet(hero_id int32) IHero {
	hero, has := this.hero_list[hero_id]
	if !has {
		return nil
	}
	return hero
}

func (this *HeroSys) HeroRandomType(ban_id []int32) int32 {
	ban := make(map[int32]bool)
	if ban_id != nil {
		for _, v := range ban_id {
			ban[v] = true
		}
	}

	list := make([]int32, 0)
	for _, v := range this.hero_list {
		if _, has := ban[v.GetSchemeId()]; !has {
			list = append(list, v.GetSchemeId())
		}
	}

	if len(list) > 0 {
		r := rand.Intn(len(list))
		return list[r]
	}

	return 0
}

func (this *HeroSys) HeroObtain(hero_scheme_id int32, hero_lv int32, hero_rank int32, is_notify bool) (int32, common.RetCode) {
	LogDebug("HeroObtain Enter : hero_scheme_id(", hero_scheme_id, ") hero_lv(", hero_lv, ") hero_rank(", hero_rank, ")")
	if _, has := this.hero_list[hero_scheme_id]; has {
		return 0, common.RetCode_HeroUidError
	}

	hero, err := NewHero(hero_scheme_id, hero_lv, hero_rank, this.owner.GetUid())
	if err != nil {
		LogError(err)
		return 0, common.RetCode_SchemeData_Error
	}
	this.hero_list[hero.GetSchemeId()] = hero
	this.static_hero(hero)
	this.Save(hero)
	LogDebug("HeroObtain Success : role_uid(", this.owner.GetUid(), ") hero_scheme_id(", hero_scheme_id, ")")

	if is_notify {
		this.send_info_notify(hero)
	}

	if hero.scheme_rank_data.Rank == 2 {
		this.owner.AchievementAddNum(7, 1, false)
	}
	if hero.scheme_rank_data.Rank == 4 {
		this.owner.AchievementAddNum(8, 1, false)
	}

	return hero.GetSchemeId(), common.RetCode_Success
}

func (this *HeroSys) HeroLvUp(hero_id int32, item_scheme_id int32, num int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_id]
	if !has {
		return common.RetCode_HeroUidError
	}

	if item_scheme_id != hero.scheme_base_data.NeedItemId {
		return common.RetCode_HeroExpArgsError
	}

	old_lv := hero.GetLv()
	ret := hero.AddLvExp(num, this.owner)
	if ret == common.RetCode_Success {
		if hero.GetLv() > old_lv {
			this.static_hero(hero)
		}

		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		this.owner.AchievementAddNum(4, hero.GetLv(), true)
	}
	return ret
}

func (this *HeroSys) HeroSkillLvUp(hero_id int32, skill_id int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_id]
	if !has {
		return common.RetCode_HeroUidError
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

func (this *HeroSys) HeroSize() int32 {
	return int32(len(this.hero_list))
}

func (this *HeroSys) HeroFind(hero_id int32) bool {
	_, has := this.hero_list[hero_id]
	return has
}

func (this *HeroSys) HeroPopulation(hero_id int32) int32 {
	hero, has := this.hero_list[hero_id]
	if has {
		return hero.GetPopulation()
	}
	return 0
}

func (this *HeroSys) HeroAddRank(hero_id, value int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_id]
	if !has {
		return common.RetCode_HeroUidError
	}

	ret := hero.AddRankExp(value)
	if ret == common.RetCode_Success {
		hero.freshSkillGM()
		this.static_hero(hero)
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		if hero.scheme_rank_data.Rank == 3 {
			this.owner.AchievementAddNum(7, 1, false)
		}
	}
	return ret
}

func (this *HeroSys) HeroEditLv(hero_id int32, lv int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_id]
	if !has {
		return common.RetCode_HeroUidError
	}

	ret := hero.EditLv(lv)
	if ret == common.RetCode_Success {
		this.static_hero(hero)
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		this.owner.AchievementAddNum(4, hero.GetLv(), true)
	}
	return ret
}

func (this *HeroSys) HeroEditRank(hero_id int32, rank int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_id]
	if !has {
		return common.RetCode_HeroUidError
	}

	ret := hero.EditRank(rank)
	if ret == common.RetCode_Success {
		hero.freshSkillGM()
		this.static_hero(hero)
		this.Save(hero)
		if is_notify {
			this.send_info_notify(hero)
		}

		//完成成就
		if hero.scheme_rank_data.Rank == 3 {
			this.owner.AchievementAddNum(7, 1, false)
		}
	}
	return ret
}

func (this *HeroSys) HeroEditSkillLv(hero_id int32, skill_id int32, lv int32, is_notify bool) common.RetCode {
	hero, has := this.hero_list[hero_id]
	if !has {
		return common.RetCode_HeroUidError
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

func (this *HeroSys) static_hero(hero *Hero) {
	msg := &static.MsgStaticHero{}
	msg.SetRoleUid(this.owner.GetUid())
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
