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

	awardinfo := make([]*protocol.AwardInfo, 0)
	ret := common.RetCode_Success
	for _, v := range award.Reward {
		switch v.Type {
		case common.RTYPE_EXP:
			role.AddExp(v.Amount, is_notify, true)
		case common.RTYPE_BLOOD:
			role.AddBlood(v.Amount, is_notify, true)
		case common.RTYPE_SOUL:
			role.AddSoul(v.Amount, is_notify, true)
		case common.RTYPE_ORDER:
			role.AddOrder(v.Amount, is_notify, true)
		case common.RTYPE_SPIRIT:
			role.AddSpirit(v.Amount, is_notify, true)
		case common.RTYPE_STONE:
			role.AddStone(v.Amount, is_notify, true)
		case common.RTYPE_GOLD:
			role.AddGold(v.Amount, is_notify, true)
		case common.RTYPE_MAGIC_HERO:
			if role.HeroObtain(v.Code/10, 1, v.Code%10, is_notify) == common.UID_FAILED {
				ret = common.RetCode_Fail
			}
		case common.RTYPE_ITEM:
			role.ItemAdd(v.Code, v.Amount, is_notify)
		case common.RTYPE_MAGIC_SOLDIER:
			role.SoldierCreateFree(v.Code, v.Amount)
		case common.RTYPE_SHIELD:
			role.AddShield(int64(v.Amount), is_notify, true)
		case common.RTYPE_DECORATION:
			role.DecorationObtain(v.Code, v.Amount)
		case common.RTYPE_BUILDING:
			role.BuildingObtain(v.Code, 1, true)
		default:
			LogError(fmt.Sprintf("FixAward type(%v) error", v.Type))
			continue
		}

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
						switch data.Type {
						case common.RTYPE_EXP:
							role.AddExp(data.Amount, is_notify, true)
						case common.RTYPE_BLOOD:
							role.AddBlood(data.Amount, is_notify, true)
						case common.RTYPE_SOUL:
							role.AddSoul(data.Amount, is_notify, true)
						case common.RTYPE_ORDER:
							role.AddOrder(data.Amount, is_notify, true)
						case common.RTYPE_SPIRIT:
							role.AddSpirit(data.Amount, is_notify, true)
						case common.RTYPE_STONE:
							role.AddStone(data.Amount, is_notify, true)
						case common.RTYPE_GOLD:
							role.AddGold(data.Amount, is_notify, true)
						case common.RTYPE_MAGIC_HERO:
							if role.HeroObtain(data.Code/10, 1, data.Code%10, is_notify) == common.UID_FAILED {
								ret = common.RetCode_Fail
							}
						case common.RTYPE_ITEM:
							role.ItemAdd(data.Code, data.Amount, is_notify)
						case common.RTYPE_MAGIC_SOLDIER:
							role.SoldierCreateFree(data.Code, data.Amount)
						case common.RTYPE_SHIELD:
							role.AddShield(int64(data.Amount), is_notify, true)
						case common.RTYPE_DECORATION:
							role.DecorationObtain(data.Code, data.Amount)
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

	return awardinfo, ret
}
