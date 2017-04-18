package mall

import (
	. "Gameserver/cache"
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	"common"
	"common/define"
	"common/protocol"
	"common/scheme"
	"common/static"
	"fmt"
	. "galaxy"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	cache_mall_key_t = "Role:%v:Mall"
)

func genMallCacheKey(role_uid int64) string {
	return fmt.Sprintf(cache_mall_key_t, role_uid)
}

type MallSys struct {
	owner IRole
	MallCache
	cache_key string
}

func (this *MallSys) Init(owner IRole) {
	this.owner = owner
	this.MallCache.MallList = make(map[int32]int64)
	this.cache_key = genMallCacheKey(this.owner.GetUid())
}

func (this *MallSys) Load() error {
	resp, err := GxService().Redis().Cmd("GET", this.cache_key)
	if err != nil {
		return err
	}

	if buf, _ := resp.Bytes(); buf != nil {
		err = proto.Unmarshal(buf, &this.MallCache)
		if err != nil {
			LogError(err)
			return err
		}
	}

	return nil
}

func (this *MallSys) Save() {
	buf, err := proto.Marshal(&this.MallCache)
	if err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SET", this.cache_key, buf); err != nil {
		LogFatal(err)
		return
	}
}

func (this *MallSys) FillMallInfo() []*protocol.MallInfo {
	infos := make([]*protocol.MallInfo, len(this.MallCache.GetMallList()))
	index := 0
	for id, value := range this.MallCache.GetMallList() {
		infos[index] = &protocol.MallInfo{
			MallId: proto.Int32(id),
			Args:   proto.Int64(value),
		}
		index++
	}
	return infos
}

func (this *MallSys) MallBuy(id int32) (common.RetCode, int32, int64) {
	mall_scheme, has := scheme.Mallmap[id]
	if !has {
		LogError("Mall Scheme (", id, ") error")
		return common.RetCode_SchemeData_Error, 0, 0
	}

	//check
	need := mall_scheme.Price
	switch mall_scheme.MallType {
	case common.MALL_TYPE_NONE:
	case common.MALL_TYPE_CD:
		if value, has := this.MallCache.MallList[id]; has {
			if time.Now().Unix() < value {
				return common.RetCode_CoolDown, id, 0
			}
		}
	case common.MALL_TYPE_ORDER:
		var value int32
		if temp, has := this.MallCache.MallList[id]; has {
			value = int32(temp)
		} else {
			value = 1
		}

		switch {
		case value < scheme.Commonmap[define.TimeNum1].Value:
			need = scheme.Commonmap[define.OrderCost1].Value
		case value >= scheme.Commonmap[define.TimeNum1].Value && value < scheme.Commonmap[define.TimeNum2].Value:
			need = scheme.Commonmap[define.OrderCost2].Value
		case value >= scheme.Commonmap[define.TimeNum2].Value && value < scheme.Commonmap[define.TimeNum3].Value:
			need = scheme.Commonmap[define.OrderCost3].Value
		case value >= scheme.Commonmap[define.TimeNum3].Value && value < scheme.Commonmap[define.TimeNum4].Value:
			need = scheme.Commonmap[define.OrderCost4].Value
		}
	case common.MALL_TYPE_SPIRIT:
		var value int32
		if temp, has := this.MallCache.MallList[id]; has {
			value = int32(temp)
		} else {
			value = 1
		}

		switch {
		case value < scheme.Commonmap[define.TimeNum1].Value:
			need = scheme.Commonmap[define.SpiritCost1].Value
		case value >= scheme.Commonmap[define.TimeNum1].Value && value < scheme.Commonmap[define.TimeNum2].Value:
			need = scheme.Commonmap[define.SpiritCost2].Value
		case value >= scheme.Commonmap[define.TimeNum2].Value && value < scheme.Commonmap[define.TimeNum3].Value:
			need = scheme.Commonmap[define.SpiritCost3].Value
		case value >= scheme.Commonmap[define.TimeNum3].Value && value < scheme.Commonmap[define.TimeNum4].Value:
			need = scheme.Commonmap[define.SpiritCost4].Value
		}
	default:
		return common.RetCode_SchemeData_Error, id, 0
	}

	switch mall_scheme.ResourceType {
	case common.RTYPE_BLOOD:
		if !this.owner.IsEnoughBlood(need) {
			return common.RetCode_Unable, id, 0
		}
	case common.RTYPE_SOUL:
		if !this.owner.IsEnoughSoul(need) {
			return common.RetCode_Unable, id, 0
		}
	case common.RTYPE_GOLD:
		if !this.owner.IsEnoughGold(need) {
			return common.RetCode_Unable, id, 0
		}
	case common.RTYPE_STONE:
		if !this.owner.IsEnoughStone(need) {
			return common.RetCode_Unable, id, 0
		}
	default:
		return common.RetCode_SchemeData_Error, id, 0
	}

	_, retcode := Award(mall_scheme.AwardId, this.owner, true)
	if retcode != common.RetCode_Success {
		return retcode, id, 0
	}

	switch mall_scheme.ResourceType {
	case common.RTYPE_BLOOD:
		this.owner.CostBlood(need, true, true)
	case common.RTYPE_SOUL:
		this.owner.CostSoul(need, true, true)
	case common.RTYPE_GOLD:
		this.owner.CostGold(need, true, true)
		this.owner.StaticPayLog(int32(static.PayType_item), id, need)
	case common.RTYPE_STONE:
		this.owner.CostStone(need, true, true)
	}

	switch mall_scheme.MallType {
	case common.MALL_TYPE_NONE:
		return retcode, id, 0
	case common.MALL_TYPE_CD:
		this.MallCache.MallList[id] = time.Now().Unix() + int64(mall_scheme.Args1)
		this.Save()
		return retcode, id, this.MallCache.MallList[id]
	case common.MALL_TYPE_ORDER, common.MALL_TYPE_SPIRIT:
		if value, has := this.MallCache.MallList[id]; has {
			this.MallCache.MallList[id] = value + 1
		} else {
			this.MallCache.MallList[id] = 1
		}
		this.Save()
		return retcode, id, this.MallCache.MallList[id]
	}

	return common.RetCode_SchemeData_Error, id, 0
}