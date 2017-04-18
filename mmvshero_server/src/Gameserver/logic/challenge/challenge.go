package challenge

import (
	. "Gameserver/logic"
	"Gameserver/logic/award"
	"common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"math/rand"
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	MAX_STAGE             = 15
	cache_challenge_key_t = "Role:%v:Challenge"
)

type ChallengeSys struct {
	ChallengeCache

	owner     IRole
	cache_key string
}

func (this *ChallengeSys) Init(owner IRole) {
	this.owner = owner
	this.cache_key = fmt.Sprintf(cache_challenge_key_t, this.owner.GetUid())
	this.ChallengeCache.Layer = make(map[int32]*ChallengeLayer)
}

func (this *ChallengeSys) Load() error {
	resp, err := GxService().Redis().Cmd("GET", this.cache_key)
	if err != nil {
		LogError(err)
		return err
	}

	if buf, _ := resp.Bytes(); buf != nil {
		err := proto.Unmarshal(buf, &this.ChallengeCache)
		if err != nil {
			LogError(err)
			return err
		}
	}

	if this.ChallengeCache.Layer == nil {
		this.ChallengeCache.Layer = make(map[int32]*ChallengeLayer)
	}

	this.checkFreshTime()
	return nil
}

func (this *ChallengeSys) Save() error {
	buf, err := proto.Marshal(&this.ChallengeCache)
	if err != nil {
		LogFatal(err)
		return err
	}

	if _, err := GxService().Redis().Cmd("SET", this.cache_key, buf); err != nil {
		LogFatal(err)
		return err
	}
	return nil
}

func (this *ChallengeSys) fresh() {
	//reset
	this.SetFreshTime(Time())
	this.SetChallengeCount(0)
	this.SetChallengeChance(0)
	this.SetChallengeResetTimes(0)
	this.SetCurLayer(0)

	this.freshStage()

	this.Save()
}

func (this *ChallengeSys) freshStage() {
	//生成关卡
	this.ChallengeCache.Layer = make(map[int32]*ChallengeLayer)
	for i := 1; i <= MAX_STAGE; i++ {
		layer_scheme := scheme.ChallengeStageRandGet(int32(i))
		if layer_scheme == nil {
			LogError("layer_scheme nil : ", i)
			continue
		}

		layer := new(ChallengeLayer)
		layer.SetLayerId(layer_scheme.Id)
		layer.WaveInfo = make(map[int32]int32)
		if layer_scheme.Wave > 0 {
			for w := int32(0); w < layer_scheme.Wave; w++ {
				wave_scheme := scheme.ChallengeTroopRandGet(layer_scheme.Type, w+1, this.owner.GetLv())
				if wave_scheme == nil {
					LogError("layer_scheme nil : ", i, layer_scheme.Type, w+1, this.owner.GetUid(), this.owner.GetLv())
					continue
				}
				layer.WaveInfo[w+1] = wave_scheme.Id
			}
		} else {
			wave_scheme := scheme.ChallengeTroopRandGet(layer_scheme.Type, 1, this.owner.GetLv())
			if wave_scheme == nil {
				LogError("layer_scheme nil : ", i, layer_scheme.Type, 1, this.owner.GetUid(), this.owner.GetLv())
			} else {
				layer.WaveInfo[1] = wave_scheme.Id
			}
		}

		layer.SetBanBuildingType(-1)
		layer.SetBanSoldierId(-1)
		layer.SetBanKingskillId(-1)

		battle_type_list := strings.Split(layer_scheme.BattleType, ";")
		for _, v := range battle_type_list {
			switch v {
			case "1":
				zero_rand := rand.Int31n(100)
				if zero_rand < 20 {
					layer.SetBanBuildingType(0)
				} else {
					layer.SetBanBuildingType(this.owner.BuildingRandomType(nil))
				}
			case "2":
				zero_rand := rand.Int31n(100)
				if zero_rand < 20 {
					layer.SetBanSoldierId(0)
				} else {
					layer.SetBanSoldierId(this.owner.BuildingRandomType(nil))
				}
			case "3":
				zero_rand := rand.Int31n(100)
				if zero_rand < 20 {
					layer.SetBanKingskillId(0)
				} else {
					layer.SetBanKingskillId(this.owner.BuildingRandomType(nil))
				}
			}
		}

		//奖励
		award_list := award.AwardGenEx(layer_scheme.Award, this.owner)
		layer.Awards = make([]*PlunderAwardCache, 0)
		if len(award_list) > 0 {
			for _, v := range award_list {
				a := &PlunderAwardCache{}
				a.SetType(v.GetType())
				a.SetCode(v.GetCode())
				a.SetAmount(v.GetAmount())
				layer.Awards = append(layer.Awards, a)
			}
		}

		this.ChallengeCache.Layer[int32(i)] = layer
	}
}

func (this *ChallengeSys) checkFreshTime() {
	if this.owner.GetLv() < scheme.ModuleOpenmap[common.MODULEOPEN_CHALLENGE].NeedRoleLv {
		return
	}

	if this.GetFreshTime() < RefreshTime(scheme.Commonmap[define.SysResetTime].Value) {
		LogDebug("checkFreshTime need fresh owner_id : ", this.owner.GetUid(), " lv : ", this.owner.GetLv())
		this.fresh()
	}
}

func (this *ChallengeSys) checkIsFresh() bool {
	if this.owner.GetLv() < scheme.ModuleOpenmap[common.MODULEOPEN_CHALLENGE].NeedRoleLv {
		return false
	}

	if this.GetFreshTime() < RefreshTime(scheme.Commonmap[define.SysResetTime].Value) {
		return true
	}
	return false
}

func (this *ChallengeSys) ChallengeQuery() *protocol.MsgChallengeQueryRet {
	LogDebug("EnterChallengeQuery owner_id : ", this.owner.GetUid(), " lv : ", this.owner.GetLv())
	this.checkFreshTime()

	ret := new(protocol.MsgChallengeQueryRet)
	ret.SetChallengeCount(this.GetChallengeCount())
	ret.SetChallengeChance(this.GetChallengeChance())
	ret.SetCurLayer(this.GetCurLayer())
	ret.SetChallengeResetTime(this.GetChallengeResetTimes())

	layers := make([]*protocol.ChallengeLayerInfo, 0)
	for k, v := range this.GetLayer() {
		layer := new(protocol.ChallengeLayerInfo)
		layer.SetLayer(k)
		layer.SetLayerId(v.GetLayerId())
		awards := make([]*protocol.AwardInfo, 0)

		for _, v1 := range v.GetAwards() {
			award := new(protocol.AwardInfo)
			award.SetAmount(v1.GetAmount())
			award.SetCode(v1.GetCode())
			award.SetType(v1.GetType())
			awards = append(awards, award)
		}
		layer.SetAwards(awards)

		waveInfos := make([]*protocol.ChallengeWaveInfo, 0)
		for id, wave := range v.GetWaveInfo() {
			waveInfo := new(protocol.ChallengeWaveInfo)
			waveInfo.SetId(id)
			waveInfo.SetWave(wave)

			waveInfos = append(waveInfos, waveInfo)
		}
		layer.SetWaves(waveInfos)
		layer.SetBanBuildingType(v.GetBanBuildingType())
		layer.SetBanKingskillId(v.GetBanKingskillId())
		layer.SetBanSoldierId(v.GetBanSoldierId())

		layers = append(layers, layer)
	}
	ret.SetLayers(layers)

	LogDebug("EndChallengeQuery owner_id : ", this.owner.GetUid(), " lv : ", this.owner.GetLv())

	return ret
}

func (this *ChallengeSys) ChallengeStartFight(layerNum int32) (ret *protocol.MsgChallengeStartFightRet) {
	ret = new(protocol.MsgChallengeStartFightRet)
	if layerNum != this.GetCurLayer()+1 || layerNum > MAX_STAGE || this.checkIsFresh() {
		ret.SetRetcode(int32(common.RetCode_Failed))
		return
	}

	if _, has := this.Layer[layerNum]; has {
		//检查挑战次数
		if this.GetChallengeCount() >= scheme.Commonmap[define.ChallengeLimit].Value {
			ret.SetRetcode(int32(common.RetCode_CD))
			return
		}

		//检查挑战机会
		if this.GetChallengeChance() >= scheme.Commonmap[define.ChallengeChance].Value {
			ret.SetRetcode(int32(common.RetCode_CD))
			return
		}

		if layerNum == 1 {
			this.SetChallengeCount(this.GetChallengeCount() + 1)
			this.Save()
		}

		ret.SetRetcode(int32(common.RetCode_Success))
		return
	}

	ret.SetRetcode(int32(common.RetCode_Failed))
	return
}

func (this *ChallengeSys) ChallengeFightResult(layerNum int32, isSuccess bool) (ret *protocol.MsgChallengeFightResultRet) {
	ret = new(protocol.MsgChallengeFightResultRet)
	ret.SetChallengeChance(this.GetChallengeChance())
	if layerNum != this.GetCurLayer()+1 || layerNum > MAX_STAGE {
		ret.SetRetcode(int32(common.RetCode_Failed))
		return
	}

	if layer, has := this.Layer[layerNum]; has {
		if isSuccess {
			if len(layer.Awards) > 0 {
				awards := make([]*protocol.AwardInfo, 0)
				for _, v := range layer.Awards {
					a := &protocol.AwardInfo{}
					a.SetType(v.GetType())
					a.SetCode(v.GetCode())
					a.SetAmount(v.GetAmount())
					awards = append(awards, a)
				}
				award.AwardByInfo(awards, this.owner, true)
			}

			if this.GetCurLayer() != int32(len(this.GetLayer())) {
				this.SetCurLayer(this.GetCurLayer() + 1)
			}

			//添加成就
			this.owner.AchievementAddNum(16, layerNum, true)
		} else {
			this.SetChallengeChance(this.GetChallengeChance() + 1)
		}

		this.Save()
		ret.SetRetcode(int32(common.RetCode_Success))
		ret.SetChallengeChance(this.GetChallengeChance())
		return
	}

	ret.SetRetcode(int32(common.RetCode_Failed))
	return
}

func (this *ChallengeSys) ChallengeReset() (ret *protocol.MsgChallengeResetRet) {
	ret = new(protocol.MsgChallengeResetRet)
	ret.SetRetcode(int32(common.RetCode_Success))
	ret.SetChallengeChance(this.GetChallengeChance())
	ret.SetChallengeCount(this.GetChallengeCount())
	ret.SetChallengeResetTime(this.GetChallengeResetTimes())
	if this.GetChallengeResetTimes() >= scheme.Commonmap[define.ChallengeResetTimes].Value || this.checkIsFresh() {
		ret.SetRetcode(int32(common.RetCode_Failed))
		return
	}

	if !this.owner.IsEnoughGold(scheme.Commonmap[define.ChallengeResetCost].Value) {
		ret.SetRetcode(int32(common.RetCode_RoleNotEnoughGold))
		return
	}
	this.owner.CostGold(scheme.Commonmap[define.ChallengeResetCost].Value, true, true)

	this.SetChallengeCount(0)
	this.SetChallengeChance(0)
	this.SetChallengeResetTimes(this.GetChallengeResetTimes() + 1)
	this.SetCurLayer(0)
	this.freshStage()

	layers := make([]*protocol.ChallengeLayerInfo, 0)
	for k, v := range this.GetLayer() {
		layer := new(protocol.ChallengeLayerInfo)
		layer.SetLayer(k)
		layer.SetLayerId(v.GetLayerId())
		awards := make([]*protocol.AwardInfo, 0)

		for _, v1 := range v.GetAwards() {
			award := new(protocol.AwardInfo)
			award.SetAmount(v1.GetAmount())
			award.SetCode(v1.GetCode())
			award.SetType(v1.GetType())
			awards = append(awards, award)
		}
		layer.SetAwards(awards)

		waveInfos := make([]*protocol.ChallengeWaveInfo, 0)
		for id, wave := range v.GetWaveInfo() {
			waveInfo := new(protocol.ChallengeWaveInfo)
			waveInfo.SetId(id)
			waveInfo.SetWave(wave)

			waveInfos = append(waveInfos, waveInfo)
		}
		layer.SetWaves(waveInfos)
		layer.SetBanBuildingType(v.GetBanBuildingType())
		layer.SetBanKingskillId(v.GetBanKingskillId())
		layer.SetBanSoldierId(v.GetBanSoldierId())

		layers = append(layers, layer)
	}
	ret.SetLayers(layers)

	this.Save()

	return
}
