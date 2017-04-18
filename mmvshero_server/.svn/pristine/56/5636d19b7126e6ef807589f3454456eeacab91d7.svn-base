package activity

import (
	. "Gameserver/logic"
	"Gameserver/logic/award"
	"common"
	"common/protocol"
	"common/scheme"
	. "galaxy"
	"time"
)

type activityGrowFundBuy struct {
	Index int32
	Cost  int32
}

type activityGrowFundAward struct {
	Index     int32
	Condition int32
	Award     int32
}

type activityGrowFundData struct {
	ActId    int32
	Buy      []*activityGrowFundBuy
	Award    []*activityGrowFundAward
	AwardMap map[int32]*activityGrowFundAward
}

func loadActivityGrowFundData(act_data map[int32]interface{}) error {
	for _, v := range scheme.ActivityGrowFunds {
		data := new(activityGrowFundData)
		data.ActId = v.Id

		data.Buy = make([]*activityGrowFundBuy, 0)
		for _, v := range v.Buy {
			b := new(activityGrowFundBuy)
			b.Index = v.Index
			b.Cost = v.Cost
			data.Buy = append(data.Buy, b)
		}

		data.Award = make([]*activityGrowFundAward, 0)
		for _, v := range v.Award {
			a := new(activityGrowFundAward)
			a.Index = v.Index
			a.Condition = v.Condition
			a.Award = v.Award
			data.Award = append(data.Award, a)
		}

		data.AwardMap = make(map[int32]*activityGrowFundAward)
		for _, d := range data.Award {
			data.AwardMap[d.Index] = d
		}

		act_data[v.Id] = data
	}
	return nil
}

type activityGrowFund struct {
	activityBase
	data *activityGrowFundData
}

func (this *activityGrowFund) init(list_data *activityListData, data interface{}) {
	this.activityBase.list_data = list_data
	this.data = data.(*activityGrowFundData)
}

func (this *activityGrowFund) onOpen(now *time.Time) {}

func (this *activityGrowFund) onUpdate(now *time.Time) {}

func (this *activityGrowFund) onClose(now *time.Time) {}

func (this *activityGrowFund) onInitUserData(role IRole) {}

func (this *activityGrowFund) onCheckCondition(role IRole, need_notity bool, args ...interface{}) {
	for _, v := range this.data.Award {
		if role.GetLv() >= v.Condition && role.ActGetStatus(this.getId(), this.data.Buy[0].Index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanGet &&
			role.ActGetStatus(this.getId(), v.Index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanNotGet {
			role.ActUpdateStatus(this.getId(), v.Index, protocol.ActivityStatusFlag_CanGet, this.list_data.StartTimestamp)
			if need_notity {
				sendAwardStatusNotify(role, this.getId(), v.Index, role.ActGetStatus(this.getId(), v.Index, this.list_data.IsClean, this.list_data.StartTimestamp))
			}
		}
	}
}

func (this *activityGrowFund) onQuery(role IRole) (int_value []int32, award_box_info []*protocol.ActivityBoxInfo) {
	int_value = make([]int32, 0)
	award_box_info = make([]*protocol.ActivityBoxInfo, 0)
	for _, v := range this.data.Award {
		box := new(protocol.ActivityBoxInfo)
		box.SetIndex(v.Index)
		box.SetStatus(int32(role.ActGetStatus(this.getId(), v.Index, this.list_data.IsClean, this.list_data.StartTimestamp)))
		box.SetCondition(0)
		award_box_info = append(award_box_info, box)
	}
	return
}

func (this *activityGrowFund) onOperate(role IRole, index int32) (retcode common.RetCode, award_info []*protocol.AwardInfo) {
	retcode = common.RetCode_Failed
	award_info = make([]*protocol.AwardInfo, 0)

	data, has := this.data.AwardMap[index]
	if !has && index != this.data.Buy[0].Index {
		LogWarning("onOperate act_id : ", this.getId(), " roleuid : ", role.GetUid(), " index : ", index, " error")
		return
	}

	if index == this.data.Buy[0].Index {
		if role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanNotGet {
			if !role.IsEnoughGold(this.data.Buy[0].Cost) {
				retcode = common.RetCode_RoleNotEnoughGold
				return
			}
			role.CostGold(this.data.Buy[0].Cost, true, true)
			role.ActUpdateStatus(this.getId(), index, protocol.ActivityStatusFlag_CanGet, this.list_data.StartTimestamp)
			sendAwardStatusNotify(role, this.getId(), index, role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp))
			this.onCheckCondition(role, true)
		}
	} else {
		if role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanGet {
			award_info, retcode = award.Award(data.Award, role, true)
			if retcode != common.RetCode_Success {
				return
			}
			role.ActUpdateStatus(this.getId(), index, protocol.ActivityStatusFlag_Geted, this.list_data.StartTimestamp)
			sendAwardStatusNotify(role, this.getId(), index, role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp))
			return
		}
	}
	return
}
