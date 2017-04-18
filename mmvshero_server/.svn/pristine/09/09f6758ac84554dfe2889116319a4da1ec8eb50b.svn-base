package award

import (
	. "Gameserver/logic"
	"common"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"math/rand"
	"time"

	"github.com/golang/protobuf/proto"
)

func Award(award_id int32, role IRole, is_notify bool) ([]*protocol.AwardInfo, common.RetCode) {
	award, has := scheme.Awardmap[award_id]
	if !has {
		return nil, common.RetCode_SchemeData_Error
	}

	//FixReward
	awardinfo := make([]*protocol.AwardInfo, 0)
	ret := common.RetCode_Success
	for _, v := range award.Reward {
		switch v.Type {
		case common.RTYPE_EXP:
			role.AddExp(v.Amount, is_notify, true)
		case common.RTYPE_SOUL:
			role.AddSoul(v.Amount, is_notify, true)
		case common.RTYPE_ORDER:
			role.AddOrder(v.Amount, is_notify, true)
		case common.RTYPE_ARENAPOINT:
			role.AddArenaPoint(v.Amount, is_notify, true)
		case common.RTYPE_GOLD:
			role.AddGold(v.Amount, is_notify, true)
		case common.RTYPE_HERO:
			_, ret = role.HeroObtain(v.Code/10, 1, v.Code%10, is_notify)
		case common.RTYPE_ITEM:
			role.ItemAdd(v.Code, v.Amount, is_notify)
		case common.RTYPE_SOLDIER:
			role.SoldierCreateFree(v.Code, v.Amount)
		case common.RTYPE_BUILDING:
			role.BuildingObtain(v.Code, 1, true)
		default:
			LogError(fmt.Sprintf("FixAward type(%v) error", v.Type))
			continue
		}

		awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(v.Type), Code: proto.Int32(v.Code), Amount: proto.Int32(v.Amount)})
	}

	//RandomReward
	if len(award.RandomReward) > 0 {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, v := range award.RandomReward {
			rate := random.Int31n(10000)
			if rate < v.Rate {
				weight := random.Int31n(v.Data[len(v.Data)-1].Weight)
				for _, data := range v.Data {
					if weight < data.Weight {
						switch data.Type {
						case common.RTYPE_EXP:
							role.AddExp(data.Amount, is_notify, true)
						case common.RTYPE_SOUL:
							role.AddSoul(data.Amount, is_notify, true)
						case common.RTYPE_ORDER:
							role.AddOrder(data.Amount, is_notify, true)
						case common.RTYPE_ARENAPOINT:
							role.AddArenaPoint(data.Amount, is_notify, true)
						case common.RTYPE_GOLD:
							role.AddGold(data.Amount, is_notify, true)
						case common.RTYPE_HERO:
							_, ret = role.HeroObtain(data.Code/10, 1, data.Code%10, is_notify)
						case common.RTYPE_ITEM:
							role.ItemAdd(data.Code, data.Amount, is_notify)
						case common.RTYPE_SOLDIER:
							role.SoldierCreateFree(data.Code, data.Amount)
						case common.RTYPE_BUILDING:
							role.BuildingObtain(data.Code, 1, true)
						default:
							LogError(fmt.Sprintf("RandomAward type(%v) error", data.Type))
							continue
						}
						awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(data.Type), Code: proto.Int32(data.Code), Amount: proto.Int32(data.Amount)})
						break
					}
				}
			}
		}
	}

	//SelfRandomReward
	if len(award.SelfRandomReward) > 0 {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, v := range award.SelfRandomReward {
			rate := random.Int31n(10000)
			if rate < v.Rate {
				weight := random.Int31n(v.Data[len(v.Data)-1].Weight)
				for _, data := range v.Data {
					if weight < data.Weight {
						switch data.Selftype {
						case common.SELFTYPE_HERO_CHIP:
							hero_id := role.HeroRandomType(data.Ban)
							if hero_id != 0 {
								chip_id := scheme.Heromap[hero_id].NeedItemId
								role.ItemAdd(chip_id, data.Amount, true)
								awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(common.RTYPE_ITEM), Code: proto.Int32(chip_id), Amount: proto.Int32(data.Amount)})
							}

						case common.SELFTYPE_HERO_STONE:
							hero_id := role.HeroRandomType(data.Ban)
							if hero_id != 0 {
								stone_id := scheme.Heromap[hero_id].RankNeedItemId
								role.ItemAdd(stone_id, data.Amount, true)
								awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(common.RTYPE_ITEM), Code: proto.Int32(stone_id), Amount: proto.Int32(data.Amount)})
							}
						case common.SELFTYPE_SOLDIER_CHIP:
							soldier_id := role.SoldierRandomType(data.Ban)
							if soldier_id != 0 {
								chip_id := scheme.Soldiermap[soldier_id].NeedItemId
								role.ItemAdd(chip_id, data.Amount, true)
								awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(common.RTYPE_ITEM), Code: proto.Int32(chip_id), Amount: proto.Int32(data.Amount)})
							}
						default:
							LogError(fmt.Sprintf("SelfRandomReward type(%v) error", data.Selftype))
							continue
						}
						break
					}
				}
			}
		}
	}

	return awardinfo, ret
}

//该接口不支持SelfRandomReward字段
func AwardGen(award_id int32) []*protocol.AwardInfo {
	award, has := scheme.Awardmap[award_id]
	if !has {
		return nil
	}

	awardinfo := make([]*protocol.AwardInfo, 0)
	for _, v := range award.Reward {
		awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(v.Type), Code: proto.Int32(v.Code), Amount: proto.Int32(v.Amount)})
	}

	if len(award.RandomReward) > 0 {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, v := range award.RandomReward {
			rate := random.Int31n(10000)
			if rate < v.Rate {
				weight := random.Int31n(v.Data[len(v.Data)-1].Weight)
				for _, data := range v.Data {
					if weight < data.Weight {
						awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(data.Type), Code: proto.Int32(data.Code), Amount: proto.Int32(data.Amount)})
						break
					}
				}
			}
		}
	}

	return awardinfo
}

func AwardGenEx(award_id int32, role IRole) []*protocol.AwardInfo {
	award, has := scheme.Awardmap[award_id]
	if !has {
		return nil
	}
	//LogDebug("Award : ", award)

	awardinfo := make([]*protocol.AwardInfo, 0)
	for _, v := range award.Reward {
		awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(v.Type), Code: proto.Int32(v.Code), Amount: proto.Int32(v.Amount)})
	}

	//LogDebug("Award Fix : ", awardinfo)

	if len(award.RandomReward) > 0 {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, v := range award.RandomReward {
			rate := random.Int31n(10000)
			if rate < v.Rate {
				weight := random.Int31n(v.Data[len(v.Data)-1].Weight)
				for _, data := range v.Data {
					if weight < data.Weight {
						awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(data.Type), Code: proto.Int32(data.Code), Amount: proto.Int32(data.Amount)})
						break
					}
				}
			}
		}
	}

	//LogDebug("Award Random : ", awardinfo)

	//SelfRandomReward
	if len(award.SelfRandomReward) > 0 {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, v := range award.SelfRandomReward {
			rate := random.Int31n(10000)
			if rate < v.Rate {
				weight := random.Int31n(v.Data[len(v.Data)-1].Weight)
				for _, data := range v.Data {
					if weight < data.Weight {
						switch data.Selftype {
						case common.SELFTYPE_HERO_CHIP:
							hero_id := role.HeroRandomType(data.Ban)
							if hero_id != 0 {
								chip_id := scheme.Heromap[hero_id].NeedItemId
								awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(common.RTYPE_ITEM), Code: proto.Int32(chip_id), Amount: proto.Int32(data.Amount)})
							}

						case common.SELFTYPE_HERO_STONE:
							hero_id := role.HeroRandomType(data.Ban)
							if hero_id != 0 {
								stone_id := scheme.Heromap[hero_id].RankNeedItemId
								awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(common.RTYPE_ITEM), Code: proto.Int32(stone_id), Amount: proto.Int32(data.Amount)})
							}
						case common.SELFTYPE_SOLDIER_CHIP:
							soldier_id := role.SoldierRandomType(data.Ban)
							if soldier_id != 0 {
								chip_id := scheme.Soldiermap[soldier_id].NeedItemId
								awardinfo = append(awardinfo, &protocol.AwardInfo{Type: proto.Int32(common.RTYPE_ITEM), Code: proto.Int32(chip_id), Amount: proto.Int32(data.Amount)})
							}
						default:
							LogError(fmt.Sprintf("SelfRandomReward type(%v) error", data.Selftype))
							continue
						}
						break
					}
				}
			}
		}
	}

	//LogDebug("Award SelfRandom : ", awardinfo)

	return awardinfo
}

func AwardByInfo(award []*protocol.AwardInfo, role IRole, is_notify bool) {
	if award == nil {
		return
	}

	for _, v := range award {
		switch v.GetType() {
		case common.RTYPE_EXP:
			role.AddExp(v.GetAmount(), is_notify, true)
		case common.RTYPE_SOUL:
			role.AddSoul(v.GetAmount(), is_notify, true)
		case common.RTYPE_ORDER:
			role.AddOrder(v.GetAmount(), is_notify, true)
		case common.RTYPE_ARENAPOINT:
			role.AddArenaPoint(v.GetAmount(), is_notify, true)
		case common.RTYPE_GOLD:
			role.AddGold(v.GetAmount(), is_notify, true)
		case common.RTYPE_HERO:
			role.HeroObtain(v.GetCode()/10, 1, v.GetCode()%10, is_notify)
		case common.RTYPE_ITEM:
			role.ItemAdd(v.GetCode(), v.GetAmount(), is_notify)
		case common.RTYPE_SOLDIER:
			role.SoldierCreateFree(v.GetCode(), v.GetAmount())
		case common.RTYPE_BUILDING:
			role.BuildingObtain(v.GetCode(), 1, true)
		default:
			LogError(fmt.Sprintf("FixAward type(%v) error", v.Type))
			continue
		}
	}
}
