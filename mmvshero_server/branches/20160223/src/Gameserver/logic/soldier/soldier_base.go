package soldier

import (
	. "Gameserver/cache"
	"Gameserver/global"
	. "Gameserver/logic"
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

	c.Exp = proto.Int32(this.exp)
	c.Timestamp = proto.Int64(this.timestamp)
	c.EvoSpeedTimeStamp = proto.Int64(this.evoSpeedTimeStamp)
	c.Active = proto.Int32(this.active)
	c.AutoId = proto.Int64(this.autoId)
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

	this.exp = cache.GetExp()
	this.timestamp = cache.GetTimestamp()
	this.evoSpeedTimeStamp = cache.GetEvoSpeedTimeStamp()
	this.active = cache.GetActive()
	this.autoId = cache.GetAutoId()
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

	s.exp = this.exp
	s.timestamp = this.timestamp
	s.evoSpeedTimeStamp = this.evoSpeedTimeStamp
	s.active = this.active
	s.autoId = this.autoId
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

	c.Exp = proto.Int32(this.exp)
	c.Timestamp = proto.Int64(this.timestamp)
	c.EvoSpeedTimeStamp = proto.Int64(this.evoSpeedTimeStamp)
	c.Active = proto.Int32(this.active)
	c.AutoId = proto.Int64(this.autoId)
	return c
}

func (this *Soldier) save(user IRole) bool {
	buff, err := proto.Marshal(this.toCache())
	if err != nil {
		galaxy.LogError(err)
		return false
	}

	if this.autoId == 0 {
		if _, err = RedisCmd("SET", this.CacheKey, buff); err != nil {
			galaxy.LogError(err)
			return false
		}

		if _, err := RedisCmd("SADD", GenSoldierListKey(user.GetUid()), this.CacheKey); err != nil {
			galaxy.LogError(err)
			return false
		}
	} else {
		if _, err = RedisCmd("SET", GenSoldierInMapKey(user.GetUid(), this.GetAutoId()), buff); err != nil {
			galaxy.LogError(err)
			return false
		}

		if _, err := RedisCmd("SADD", GenSoldierInMapListKey(user.GetUid()), this.CacheKey); err != nil {
			galaxy.LogError(err)
			return false
		}
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
func (this *Soldier) GetExp() int32 {
	return this.exp
}
func (this *Soldier) GetTimestamp() int64 {
	return this.timestamp
}
func (this *Soldier) GetEvoSpeedTimeStamp() int64 {
	return this.evoSpeedTimeStamp
}
func (this *Soldier) GetActive() int32 {
	return this.active
}
func (this *Soldier) GetAutoId() int64 {
	return this.autoId
}
