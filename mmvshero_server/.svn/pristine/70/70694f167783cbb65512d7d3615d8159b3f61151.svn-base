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

const (
	TOTAL_PAY_INDEX = -1
)

type activityTotalPayAward struct {
	Index     int32
	Condition int32
	Award     int32
}

type activityTotalPayData struct {
	ActId    int32
	Award    []*activityTotalPayAward
	AwardMap map[int32]*activityTotalPayAward
}

func loadActivityTotalPayData(act_data map[int32]interface{}) error {
	for _, v := range scheme.ActivityTotalPays {
		data := new(activityTotalPayData)
		data.ActId = v.Id

		data.Award = make([]*activityTotalPayAward, 0)
		for _, v := range v.Award {
			a := new(activityTotalPayAward)
			a.Index = v.Index
			a.Condition = v.Condition
			a.Award = v.Award
			data.Award = append(data.Award, a)
		}

		data.AwardMap = make(map[int32]*activityTotalPayAward)
		for _, d := range data.Award {
			data.AwardMap[d.Index] = d
		}

		act_data[v.Id] = data
	}
	return nil
}

type activityTotalPay struct {
	activityBase
	data *activityTotalPayData
}

func (this *activityTotalPay) init(list_data *activityListData, data interface{}) {
	this.activityBase.list_data = list_data
	this.data = data.(*activityTotalPayData)
}

func (this *activityTotalPay) onOpen(now *time.Time) {}

func (this *activityTotalPay) onUpdate(now *time.Time) {}

func (this *activityTotalPay) onClose(now *time.Time) {}

func (this *activityTotalPay) onInitUserData(role IRole) {}

func (this *activityTotalPay) onCheckCondition(role IRole, need_notity bool, args ...interface{}) {
	total_pay := role.ActGetCondition(this.getId(), TOTAL_PAY_INDEX, this.list_data.IsClean, this.list_data.StartTimestamp) + args[0].(int64)
	role.ActUpdateCondition(this.getId(), TOTAL_PAY_INDEX, total_pay, this.list_data.StartTimestamp)

	for _, v := range this.data.Award {
		if int32(total_pay) >= v.Condition && role.ActGetStatus(this.getId(), v.Index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanNotGet {
			role.ActUpdateStatus(this.getId(), v.Index, protocol.ActivityStatusFlag_CanGet, this.list_data.StartTimestamp)
			if need_notity {
				sendAwardStatusNotify(role, this.getId(), v.Index, role.ActGetStatus(this.getId(), v.Index, this.list_data.IsClean, this.list_data.StartTimestamp))
			}
		}
	}
}

func (this *activityTotalPay) onQuery(role IRole) (int_value []int32, award_box_info []*protocol.ActivityBoxInfo) {
	int_value = make([]int32, 0)
	int_value = append(int_value, int32(role.ActGetCondition(this.getId(), TOTAL_PAY_INDEX, this.list_data.IsClean, this.list_data.StartTimestamp)))

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

func (this *activityTotalPay) onOperate(role IRole, index int32) (retcode common.RetCode, award_info []*protocol.AwardInfo) {
	retcode = common.RetCode_Failed
	award_info = make([]*protocol.AwardInfo, 0)

	data, has := this.data.AwardMap[index]
	if !has {
		LogWarning("onOperate act_id : ", this.getId(), " roleuid : ", role.GetUid(), " index : ", index, " error")
		return
	}

	if role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanGet {
		award_info, retcode = award.Award(data.Award, role, true)
		if retcode != common.RetCode_Success {
			return
		}
		role.ActUpdateStatus(this.getId(), index, protocol.ActivityStatusFlag_Geted, this.list_data.StartTimestamp)
		sendAwardStatusNotify(role, this.getId(), index, role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp))
		return
	}

	return
}
