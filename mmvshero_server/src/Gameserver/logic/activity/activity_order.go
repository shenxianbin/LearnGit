package activity

import (
	. "Gameserver/logic"
	"Gameserver/logic/award"
	"common"
	"common/define"
	"common/protocol"
	"common/scheme"
	. "galaxy"
	"time"
)

const (
	SHARE_INDEX       = 2
	BATTLESHARE_INDEX = 3
	FRESH_INDEX       = -1
)

type activityOrderAward struct {
	Index     int32
	Condition []int32
	Award     int32
}

type activityOrderData struct {
	ActId    int32
	Award    []*activityOrderAward
	AwardMap map[int32]*activityOrderAward
}

func loadActivityOrderData(act_data map[int32]interface{}) error {
	for _, v := range scheme.ActivityOrders {
		data := new(activityOrderData)
		data.ActId = v.Id

		data.Award = make([]*activityOrderAward, 0)
		for _, v := range v.Award {
			a := new(activityOrderAward)
			a.Index = v.Index
			a.Condition = v.Condition
			a.Award = v.Award
			data.Award = append(data.Award, a)
		}

		data.Award = append(data.Award, &activityOrderAward{
			Index:     v.BattleShareAward[0].Index,
			Condition: v.BattleShareAward[0].Condition,
			Award:     v.BattleShareAward[0].Award,
		})

		data.AwardMap = make(map[int32]*activityOrderAward)
		for _, d := range data.Award {
			data.AwardMap[d.Index] = d
		}

		act_data[v.Id] = data
	}
	return nil
}

type activityOrder struct {
	activityBase
	data *activityOrderData
}

func (this *activityOrder) init(list_data *activityListData, data interface{}) {
	this.activityBase.list_data = list_data
	this.data = data.(*activityOrderData)
}

func (this *activityOrder) onOpen(now *time.Time) {}

func (this *activityOrder) onUpdate(now *time.Time) {}

func (this *activityOrder) onClose(now *time.Time) {}

func (this *activityOrder) onInitUserData(role IRole) {
	this.fresh(role)
	LogDebug("onInitUserData RoleUid : ", role.GetUid())
}

func (this *activityOrder) onCheckCondition(role IRole, need_notity bool, args ...interface{}) {}

func (this *activityOrder) onQuery(role IRole) (int_value []int32, award_box_info []*protocol.ActivityBoxInfo) {
	this.fresh(role)
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

func (this *activityOrder) onOperate(role IRole, index int32) (retcode common.RetCode, award_info []*protocol.AwardInfo) {
	this.fresh(role)
	retcode = common.RetCode_Failed
	award_info = make([]*protocol.AwardInfo, 0)

	data, has := this.data.AwardMap[index]
	if !has {
		LogWarning("onOperate act_id : ", this.getId(), " roleuid : ", role.GetUid(), " index : ", index, " error")
		return
	}

	LogDebug("onOperate act_id : ", this.getId(), " roleuid : ", role.GetUid(), " index : ", index)
	if index != SHARE_INDEX && index != BATTLESHARE_INDEX {
		LogDebug("onOperate order act_id : ", this.getId(), " roleuid : ", role.GetUid(), " index : ", index)
		now := time.Now()
		if int32(now.Hour()) >= data.Condition[0] && int32(now.Hour()) < data.Condition[1] && role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanGet {
			award_info, retcode = award.Award(data.Award, role, true)
			if retcode != common.RetCode_Success {
				return
			}
			role.ActUpdateStatus(this.getId(), index, protocol.ActivityStatusFlag_Geted, this.list_data.StartTimestamp)
			sendAwardStatusNotify(role, this.getId(), index, role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp))
			return

		}
	} else {
		LogDebug("onOperate share act_id : ", this.getId(), " roleuid : ", role.GetUid(), " index : ", index)
		if role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp) == protocol.ActivityStatusFlag_CanNotGet {
			role.ActUpdateStatus(this.getId(), index, protocol.ActivityStatusFlag_CanGet, this.list_data.StartTimestamp)
			sendAwardStatusNotify(role, this.getId(), index, role.ActGetStatus(this.getId(), index, this.list_data.IsClean, this.list_data.StartTimestamp))
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
	}
	return
}

func (this *activityOrder) fresh(role IRole) {
	timestamp := role.ActGetCondition(this.getId(), FRESH_INDEX, this.list_data.IsClean, this.list_data.StartTimestamp)
	if timestamp < RefreshTime(scheme.Commonmap[define.SysResetTime].Value) {
		for _, v := range this.data.Award {
			if v.Index != SHARE_INDEX {
				role.ActUpdateStatus(this.getId(), v.Index, protocol.ActivityStatusFlag_CanGet, this.list_data.StartTimestamp)
			} else if v.Index == SHARE_INDEX {
				role.ActUpdateStatus(this.getId(), v.Index, protocol.ActivityStatusFlag_CanNotGet, this.list_data.StartTimestamp)
			}
		}

		role.ActUpdateStatus(this.getId(), BATTLESHARE_INDEX, protocol.ActivityStatusFlag_CanGet, this.list_data.StartTimestamp)

		role.ActUpdateCondition(this.getId(), FRESH_INDEX, time.Now().Unix(), this.list_data.StartTimestamp)
	}
}
