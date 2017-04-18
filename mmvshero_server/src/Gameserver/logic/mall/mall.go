package mall

import (
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	"common"
	. "common/cache"
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
	this.MallCache.MallList = make(map[int32]*MallInfo)
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
			MallId:     proto.Int32(id),
			LimitCount: proto.Int32(value.GetLimitCount()),
			Cd:         proto.Int64(value.GetCd()),
			ConArgs:    proto.Int32(value.GetConArgs()),
		}
		index++
	}
	return infos
}

func (this *MallSys) MallBuy(id int32) (retcode common.RetCode, retinfo *protocol.MallInfo) {
	retcode = common.RetCode_Success
	retinfo = &protocol.MallInfo{
		MallId:     proto.Int32(id),
		LimitCount: proto.Int32(0),
		Cd:         proto.Int64(0),
		ConArgs:    proto.Int32(0),
	}

	mall_scheme, has := scheme.Mallmap[id]
	if !has {
		LogError("Mall Scheme (", id, ") error")
		retcode = common.RetCode_SchemeData_Error
		return
	}

	//check condition
	if mall_scheme.ConArgsType == common.MALL_CON_TYPE_LV && this.owner.GetLv() < mall_scheme.ConArgsValue {
		retcode = common.RetCode_MallLvNotEnough
		return
	}

	if mall_scheme.LastId != 0 {
		if _, has := this.MallCache.MallList[mall_scheme.LastId]; !has {
			retcode = common.RetCode_MallArgsError
			return
		}
	}

	if mall, has := this.MallCache.MallList[id]; has {
		if mall_scheme.LimitCount > 0 && mall.GetLimitCount() >= mall_scheme.LimitCount {
			retcode = common.RetCode_MallLimitCountFull
			return
		}

		if mall_scheme.CD != 0 && mall.GetCd() >= time.Now().Unix() {
			retcode = common.RetCode_CD
			return
		}
	}

	//cost
	need := mall_scheme.CostValue
	switch mall_scheme.MallType {
	case common.MALL_TYPE_NONE:
	case common.MALL_TYPE_ORDER:
		var value int32
		if temp, has := this.MallCache.MallList[id]; has {
			value = temp.GetLimitCount()
		} else {
			value = 1
		}

		switch {
		case value < scheme.Commonmap[define.OrderBuyTimeNum1].Value:
			need = scheme.Commonmap[define.OrderBuyCost1].Value
		case value >= scheme.Commonmap[define.OrderBuyTimeNum1].Value && value < scheme.Commonmap[define.OrderBuyTimeNum2].Value:
			need = scheme.Commonmap[define.OrderBuyCost2].Value
		case value >= scheme.Commonmap[define.OrderBuyTimeNum2].Value && value < scheme.Commonmap[define.OrderBuyTimeNum3].Value:
			need = scheme.Commonmap[define.OrderBuyCost3].Value
		case value >= scheme.Commonmap[define.OrderBuyTimeNum3].Value && value < scheme.Commonmap[define.OrderBuyTimeNum4].Value:
			need = scheme.Commonmap[define.OrderBuyCost4].Value
		case value >= scheme.Commonmap[define.OrderBuyTimeNum4].Value:
			need = scheme.Commonmap[define.OrderBuyCost5].Value
		}
	default:
		LogError("MallType : ", mall_scheme.MallType, " Error Id : ", id)
		retcode = common.RetCode_SchemeData_Error
		return
	}

	switch mall_scheme.CostType {
	case common.RTYPE_SOUL:
		if !this.owner.IsEnoughSoul(need) {
			retcode = common.RetCode_RoleNotEnoughSoul
			return
		}
	case common.RTYPE_GOLD:
		if !this.owner.IsEnoughGold(need) {
			retcode = common.RetCode_RoleNotEnoughGold
			return
		}
	default:
		LogError("CostType : ", mall_scheme.CostType, " Error Id : ", id)
		retcode = common.RetCode_SchemeData_Error
		return
	}

	_, ret := Award(mall_scheme.AwardId, this.owner, true)
	if ret != common.RetCode_Success {
		retcode = ret
		return
	}

	switch mall_scheme.CostType {
	case common.RTYPE_SOUL:
		this.owner.CostSoul(need, true, true)
	case common.RTYPE_GOLD:
		this.owner.CostGold(need, true, true)
		this.owner.StaticPayLog(int32(static.PayType_item), id, need)
	}

	//update
	if _, has := this.MallCache.MallList[id]; !has {
		this.MallCache.MallList[id] = &MallInfo{}
	}

	this.MallCache.MallList[id].SetLimitCount(this.MallCache.MallList[id].GetLimitCount() + 1)
	if mall_scheme.CD != 0 {
		this.MallCache.MallList[id].SetCd(time.Now().Unix() + int64(mall_scheme.CD))
	} else {
		this.MallCache.MallList[id].SetCd(0)
	}
	this.MallCache.MallList[id].SetConArgs(0)

	this.Save()

	retinfo.SetLimitCount(this.MallCache.MallList[id].GetLimitCount())
	retinfo.SetCd(this.MallCache.MallList[id].GetCd())
	retinfo.SetConArgs(this.MallCache.MallList[id].GetConArgs())

	return
}

func (this *MallSys) MallGoldFill(soul int32) common.RetCode {
	if soul < 0 {
		return common.RetCode_MallArgsError
	}

	gold := ResourceToCoin(common.RTYPE_SOUL, soul)
	if gold == 0 {
		return common.RetCode_Failed
	}

	if !this.owner.IsEnoughGold(gold) {
		return common.RetCode_RoleNotEnoughGold
	}

	this.owner.CostGold(gold, true, false)
	this.owner.AddSoul(soul, true, true)

	return common.RetCode_Success
}
