package soldier

import (
	"Gameserver/global"
	. "Gameserver/logic"
	. "common/cache"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func (this *Soldier) toCache() *SoldierCache {
	c := &SoldierCache{}
	c.SchemeId = proto.Int32(this.schemeId)
	c.Num = proto.Int32(this.num)
	c.Level = proto.Int32(this.level)
	c.Stage = proto.Int32(this.stage)

	c.SkillLevel = make(map[int32]int32)
	for k, v := range this.skillLevel {
		c.SkillLevel[k] = v
	}

	return c
}

func (this *Soldier) readCache(cache *SoldierCache) {
	this.schemeId = cache.GetSchemeId()
	this.num = cache.GetNum()
	this.level = cache.GetLevel()
	this.stage = cache.GetStage()
	if this.skillLevel == nil {
		this.skillLevel = make(map[int32]int32)
	}
	for k, v := range cache.GetSkillLevel() {
		this.skillLevel[k] = v
	}
}

func (this *Soldier) clone() *Soldier {
	s := &Soldier{}
	s.schemeId = this.schemeId
	s.num = this.num
	s.level = this.level
	s.stage = this.stage
	s.skillLevel = make(map[int32]int32)
	for k, v := range this.skillLevel {
		s.skillLevel[k] = v
	}

	return s
}

func (this *Soldier) toProtocol() *protocol.Soldier {
	c := new(protocol.Soldier)
	c.SchemeId = proto.Int32(this.schemeId)
	c.Num = proto.Int32(this.num)
	c.Level = proto.Int32(this.level)
	c.Stage = proto.Int32(this.stage)

	skillLevels := make([]*protocol.SkillLevel, 0)
	var i int32 = 0
	for k, v := range this.GetSkillLevel() {
		skillLevel := &protocol.SkillLevel{}
		skillLevel.SkillId = proto.Int32(k)
		skillLevel.SkillLevel = proto.Int32(v)
		skillLevels = append(skillLevels, skillLevel)
		i++
	}
	c.SkillLevel = skillLevels

	return c
}

func (this *Soldier) save(user IRole) bool {
	buff, err := proto.Marshal(this.toCache())
	if err != nil {
		galaxy.LogError(err)
		return false
	}

	if _, err = RedisCmd("SET", this.CacheKey, buff); err != nil {
		galaxy.LogError(err)
		return false
	}

	if _, err := RedisCmd("SADD", GenSoldierListKey(user.GetUid()), this.CacheKey); err != nil {
		galaxy.LogError(err)
		return false
	}

	this.notify(user)
	return true
}

func (this *Soldier) notify(user IRole) {
	ret := &protocol.MsgSoldierNotify{}
	ret.SetSoldier(this.toProtocol())

	buf, err := proto.Marshal(ret)
	if err != nil {
		galaxy.LogError(err)
		return
	}

	global.SendMsg(int32(protocol.MsgCode_SoldierNotify), user.GetSid(), buf)
}

func (this *Soldier) GetSchemeId() int32 {
	return this.schemeId
}
func (this *Soldier) GetNum() int32 {
	return this.num
}
func (this *Soldier) GetLevel() int32 {
	return this.level
}
func (this *Soldier) GetStage() int32 {
	return this.stage
}
func (this *Soldier) GetSkillLevel() map[int32]int32 {
	return this.skillLevel
}
