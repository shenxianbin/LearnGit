package activity

import (
	"Gameserver/global"
	. "Gameserver/logic"
	"common"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"galaxy/timer"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
)

func InitActivityModule(open_time string) error {
	if act_module == nil {
		act_module = new(activityModule)
		err := act_module.init(open_time)
		if err != nil {
			return err
		}
		init_protocol()
		timer.AddCdTimerEvent(1, timer.CDTimeNoCount, func() {
			act_module.update()
		})
	}
	return nil
}

func ActModule() *activityModule {
	return act_module
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_ActivityListReq), func(sid int64, msg []byte) {
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		buf, err := proto.Marshal(act_module.OnListQuery(role))
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ActivityListRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ActivityQueryReq), func(sid int64, msg []byte) {
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgActivityQueryReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		buf, err := proto.Marshal(act_module.OnQuery(role, req.GetId()))
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ActivityQueryRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ActivityOperateReq), func(sid int64, msg []byte) {
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgActivityOperateReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		buf, err := proto.Marshal(act_module.OnOperate(role, req.GetId(), req.GetIndex()))
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ActivityOperateRet), sid, buf)
	})
}

var act_module *activityModule

type activityListData struct {
	id             int32
	ActId          int32
	ActType        int32
	TimeType       int32
	StartTime      []int32
	EndTime        []int32
	RoleLv         int32
	IsOpen         int32
	IsClean        int32
	IsReceiveClose int32
	StartTimestamp int64
	EndTimestamp   int64
}

type activityModule struct {
	open_server_time time.Time
	act_list_data    map[int32]*activityListData
	act_data         map[int32]interface{}

	now_activity      map[int32]iactivity
	now_activity_type map[int32]map[int32]iactivity
}

func (this *activityModule) init(open_time string) error {
	this.act_list_data = make(map[int32]*activityListData)
	this.act_data = make(map[int32]interface{})
	this.now_activity = make(map[int32]iactivity)
	this.now_activity_type = make(map[int32]map[int32]iactivity)

	err := this.loadActivityListScheme()
	if err != nil {
		LogError(err)
		return err
	}

	err = this.initOpenServerTime(open_time)
	if err != nil {
		LogError(err)
		return err
	}

	this.initTimestamp()

	err = this.loadActivityScheme()
	if err != nil {
		LogError(err)
		return err
	}

	LogDebug("Initialize Success")

	return nil
}

func (this *activityModule) update() {
	now := time.Now()
	if now.Unix() < this.open_server_time.Unix() {
		LogDebug("Activity Update Failed!")
		return
	}

	this.checkActivity(&now)
}

func (this *activityModule) initOpenServerTime(open_time string) error {
	t, err := time.ParseInLocation("2006/1/2 15:04:05", open_time, time.Now().Location())
	if err != nil {
		return err
	}
	this.open_server_time = t
	return nil
}

func (this *activityModule) loadActivityListScheme() error {
	for _, scheme := range scheme.Activitys {
		if scheme.IsOpen == int32(protocol.ActivityOpenFlag_On) {
			data := new(activityListData)
			data.id = scheme.Id
			data.ActId = scheme.ActId
			data.ActType = scheme.ActType
			data.TimeType = scheme.TimeType
			data.RoleLv = scheme.RoleLv
			data.IsClean = scheme.IsClean
			data.IsReceiveClose = scheme.IsReceiveClose
			switch scheme.TimeType {
			case common.ACTIVITY_TIMETYPE_OPENSERVER:
				fallthrough
			case common.ACTIVITY_TIMETYPE_WEEK:
				timeStr := strings.Split(scheme.StartTime, ",")
				if len(timeStr) != 1 {
					return fmt.Errorf("Activity Scheme TimeError : Id %v, TimeType %v, StartTime %v, EndTime %v", scheme.Id, scheme.TimeType, scheme.StartTime, scheme.EndTime)
				}
				data.StartTime = make([]int32, 1)
				time, err := strconv.Atoi(timeStr[0])
				if err != nil {
					return err
				}
				data.StartTime[0] = int32(time)

				timeStr = strings.Split(scheme.EndTime, ",")
				if len(timeStr) != 1 {
					return fmt.Errorf("Activity Scheme TimeError : Id %v, TimeType %v, StartTime %v, EndTime %v", scheme.Id, scheme.TimeType, scheme.StartTime, scheme.EndTime)
				}
				data.EndTime = make([]int32, 1)
				time, err = strconv.Atoi(timeStr[0])
				if err != nil {
					return err
				}
				data.EndTime[0] = int32(time)
			case common.ACTIVITY_TIMETYPE_DATE:
				timeStr := strings.Split(scheme.StartTime, ",")
				if len(timeStr) != 6 {
					return fmt.Errorf("Activity Scheme TimeError : Id %v, TimeType %v, StartTime %v, EndTime %v", scheme.Id, scheme.TimeType, scheme.StartTime, scheme.EndTime)
				}
				data.StartTime = make([]int32, 6)
				for i, v := range timeStr {
					time, err := strconv.Atoi(v)
					if err != nil {
						return err
					}
					data.StartTime[i] = int32(time)
				}

				timeStr = strings.Split(scheme.EndTime, ",")
				if len(timeStr) != 6 {
					return fmt.Errorf("Activity Scheme TimeError : Id %v, TimeType %v, StartTime %v, EndTime %v", scheme.Id, scheme.TimeType, scheme.StartTime, scheme.EndTime)
				}
				data.EndTime = make([]int32, 6)
				for i, v := range timeStr {
					time, err := strconv.Atoi(v)
					if err != nil {
						return err
					}
					data.EndTime[i] = int32(time)
				}
			}
			this.act_list_data[data.id] = data
		}
	}
	return nil
}

func (this *activityModule) loadActivityScheme() error {
	err := loadActivityOrderData(this.act_data)
	if err != nil {
		return err
	}
	err = loadActivityLevelBagData(this.act_data)
	if err != nil {
		return err
	}
	err = loadActivityGrowFundData(this.act_data)
	if err != nil {
		return err
	}
	err = loadActivityTotalPayData(this.act_data)
	if err != nil {
		return err
	}
	return nil
}

func (this *activityModule) initTimestamp() {
	for k, v := range this.act_list_data {
		switch v.TimeType {
		case common.ACTIVITY_TIMETYPE_OPENSERVER:
			this.act_list_data[k].StartTimestamp = int64((v.StartTime[0]-1)*86400) + this.open_server_time.Unix()
			this.act_list_data[k].EndTimestamp = int64((v.EndTime[0])*86400) + this.open_server_time.Unix()
		case common.ACTIVITY_TIMETYPE_DATE:
			this.act_list_data[k].StartTimestamp = time.Date(int(v.StartTime[0]), time.Month(v.StartTime[1]), int(v.StartTime[2]), int(v.StartTime[3]), int(v.StartTime[4]), int(v.StartTime[5]), 0, time.Now().Location()).Unix()
			this.act_list_data[k].EndTimestamp = time.Date(int(v.EndTime[0]), time.Month(v.EndTime[1]), int(v.EndTime[2]), int(v.EndTime[3]), int(v.EndTime[4]), int(v.EndTime[5]), 0, time.Now().Location()).Unix()
		case common.ACTIVITY_TIMETYPE_FOREVER:
			this.act_list_data[k].StartTimestamp = this.open_server_time.Unix()
		}
	}
}

func (this *activityModule) checkActivity(now *time.Time) {
	update_info := make([]*protocol.ActivityOpenInfo, 0)
	for k, _ := range this.act_list_data {
		this.checkAct(this.act_list_data[k], now, update_info)
	}

	this.onListUpdateBroadCast(update_info)
}

func (this *activityModule) checkAct(list_data *activityListData, now *time.Time, infos []*protocol.ActivityOpenInfo) {
	switch list_data.TimeType {
	case common.ACTIVITY_TIMETYPE_OPENSERVER:
		this.checkOpenServerTimeTypeAct(list_data, now, infos)
	case common.ACTIVITY_TIMETYPE_WEEK:
		this.checkWeekTimeTypeAct(list_data, now, infos)
	case common.ACTIVITY_TIMETYPE_DATE:
		this.checkDateTimeTypeAct(list_data, now, infos)
	case common.ACTIVITY_TIMETYPE_FOREVER:
		this.setActStatusFlagOpen(list_data, now, infos)
	}
}

func (this *activityModule) checkOpenServerTimeTypeAct(list_data *activityListData, now *time.Time, infos []*protocol.ActivityOpenInfo) {
	if now.Unix() >= list_data.StartTimestamp && now.Unix() < list_data.EndTimestamp {
		this.setActStatusFlagOpen(list_data, now, infos)
	} else {
		this.setActStatusFlagClose(list_data, now, infos)
	}
}

func (this *activityModule) checkWeekTimeTypeAct(list_data *activityListData, now *time.Time, infos []*protocol.ActivityOpenInfo) {
	//不跨周
	if list_data.StartTime[0] <= list_data.EndTime[0] {
		if int32(now.Weekday()) >= list_data.StartTime[0] && int32(now.Weekday()) <= list_data.EndTime[0] {
			this.setActStatusFlagOpen(list_data, now, infos)
		} else {
			this.setActStatusFlagClose(list_data, now, infos)
		}
	} else {
		//跨周
		if int32(now.Weekday()) >= list_data.StartTime[0] || int32(now.Weekday()) <= list_data.EndTime[0] {
			this.setActStatusFlagOpen(list_data, now, infos)
		} else {
			this.setActStatusFlagClose(list_data, now, infos)
		}
	}
}

func (this *activityModule) checkDateTimeTypeAct(list_data *activityListData, now *time.Time, infos []*protocol.ActivityOpenInfo) {
	if now.Unix() >= list_data.StartTimestamp && now.Unix() < list_data.EndTimestamp {
		this.setActStatusFlagOpen(list_data, now, infos)
	} else {
		this.setActStatusFlagClose(list_data, now, infos)
	}
}

func (this *activityModule) checkPlayerCondition(list_data *activityListData, role_lv int32) bool {
	if list_data.IsOpen == int32(protocol.ActivityOpenFlag_On) && role_lv >= list_data.RoleLv {
		return true
	}
	return false
}

func (this *activityModule) setActStatusFlagOpen(list_data *activityListData, now *time.Time, infos []*protocol.ActivityOpenInfo) {
	if list_data.IsOpen == int32(protocol.ActivityOpenFlag_Off) {
		list_data.IsOpen = int32(protocol.ActivityOpenFlag_On)
		activityData := this.act_data[list_data.ActId]
		if activityData != nil {
			if this.openActivity(now, list_data, activityData) {
				info := new(protocol.ActivityOpenInfo)
				info.SetIsOpen(list_data.IsOpen)
				info.SetId(list_data.id)
				infos = append(infos, info)
				LogDebug("SetActStatusFlagOpen ActivityID : ", list_data.id, " ON")
			} else {
				LogDebug("SetActStatusFlagOpen ActivityID : ", list_data.id, " TypeError")
			}
		}
	}
}

func (this *activityModule) setActStatusFlagClose(list_data *activityListData, now *time.Time, infos []*protocol.ActivityOpenInfo) {
	if list_data.IsOpen == int32(protocol.ActivityOpenFlag_On) {
		list_data.IsOpen = int32(protocol.ActivityOpenFlag_Off)
		if this.closeActivity(list_data.id, now) {
			info := new(protocol.ActivityOpenInfo)
			info.SetIsOpen(list_data.IsOpen)
			info.SetId(list_data.id)
			infos = append(infos, info)
			LogDebug("setActStatusFlagClose ActivityID : ", list_data.id, " OFF")
		} else {
			LogDebug("setActStatusFlagClose ActivityID : ", list_data.id, " TypeError")
		}
	}
}

func (this *activityModule) insertNowActTypeList(activity iactivity) {
	if _, has := this.now_activity_type[activity.getActType()]; has {
		this.now_activity_type[activity.getActType()][activity.getId()] = activity
	} else {
		this.now_activity_type[activity.getActType()] = make(map[int32]iactivity)
		this.now_activity_type[activity.getActType()][activity.getId()] = activity
	}
}

func (this *activityModule) deleteNowActTypeList(act_type int32, id int32) {
	if _, has := this.now_activity_type[act_type]; has {
		delete(this.now_activity_type[act_type], id)
		if len(this.now_activity_type[act_type]) <= 0 {
			delete(this.now_activity_type, act_type)
		}
	}
}

//新增加活动
func (this *activityModule) openActivity(now *time.Time, list_data *activityListData, data interface{}) bool {
	LogDebug("openActivity enter ActivityID : ", list_data.id)
	var activity iactivity
	switch list_data.ActType {
	case ActivityTypeOrder:
		activity = new(activityOrder)
		activity.init(list_data, data)
	case ActivityTypeLevelBag:
		activity = new(activityLevelBag)
		activity.init(list_data, data)
	case ActivityTypeGrowFund:
		activity = new(activityGrowFund)
		activity.init(list_data, data)
	case ActivityTypeTotalPay:
		activity = new(activityTotalPay)
		activity.init(list_data, data)
	}

	if activity != nil {
		this.now_activity[list_data.id] = activity
		this.insertNowActTypeList(activity)
		activity.initRelativeTime(now)
		activity.onOpen(now)
		activity.onUpdate(now)
		return true
	}

	return false
}

//删除活动
func (this *activityModule) closeActivity(id int32, now *time.Time) bool {
	activity, has := this.now_activity[id]
	if has {
		activity.onClose(now)
		delete(this.now_activity, id)
		this.deleteNowActTypeList(activity.getActType(), activity.getId())
		return true
	}

	return false
}

func (this *activityModule) isCanEnterActivity(activity iactivity, role IRole) bool {
	return this.checkPlayerCondition(activity.getListData(), role.GetLv())
}

func (this *activityModule) onListUpdateBroadCast(infos []*protocol.ActivityOpenInfo) {
	if len(infos) > 0 {
		msg := new(protocol.MsgActivityListNotify)
		msg.SetInfos(infos)

		buf, err := proto.Marshal(msg)
		if err != nil {
			return
		}

		global.SendBroadCast(int32(protocol.MsgCode_ActivityListNotify), 0, buf)
	}
}

func (this *activityModule) OnListQuery(role IRole) (ret *protocol.MsgActivityListRet) {
	ret = &protocol.MsgActivityListRet{}
	ret.Infos = make([]*protocol.ActivityOpenInfo, 0)

	for _, v := range this.now_activity {
		info := new(protocol.ActivityOpenInfo)
		info.SetId(v.getId())
		info.SetIsOpen(int32(protocol.ActivityOpenFlag_On))
		ret.Infos = append(ret.Infos, info)
	}
	LogDebug(" OnListQuery RoleId : ", role.GetUid(), " Infos: ", ret.Infos)
	return
}

func (this *activityModule) GetOpenServerTime() int64 {
	return this.open_server_time.Unix()
}

func (this *activityModule) OnQuery(role IRole, id int32) (ret *protocol.MsgActivityQueryRet) {
	ret = new(protocol.MsgActivityQueryRet)
	ret.SetId(id)
	ret.SetIntValues(make([]int32, 0))
	ret.SetInfos(make([]*protocol.ActivityBoxInfo, 0))

	activity, has := this.now_activity[id]
	if has {
		if this.isCanEnterActivity(activity, role) {
			int_value, box_info := activity.onQuery(role)
			ret.SetIntValues(int_value)
			ret.SetInfos(box_info)
			return
		}
	} else {
		LogDebug(" OnQuery Activity Null ID : ", id, " RoleUid : ", role.GetUid())
	}
	return
}

func (this *activityModule) OnOperate(role IRole, id int32, index int32) (ret *protocol.MsgActivityOperateRet) {
	ret = new(protocol.MsgActivityOperateRet)
	ret.SetRetcode(int32(common.RetCode_Failed))
	ret.SetId(id)
	ret.SetInfos(make([]*protocol.AwardInfo, 0))

	activity, has := this.now_activity[id]
	if has {
		if this.isCanEnterActivity(activity, role) {
			retcode, award_info := activity.onOperate(role, index)
			ret.SetRetcode(int32(retcode))
			ret.SetInfos(award_info)
			return
		}
	} else {
		LogDebug(" OnOperate Activity Null ID : ", id, " RoleUid : ", role.GetUid())
	}
	return
}

//活动玩家初始化方法
func (this *activityModule) OnInitUserData(role IRole, act_type int32) {
	if activitys, has := this.now_activity_type[act_type]; has {
		for k, _ := range activitys {
			activitys[k].onInitUserData(role)
		}
	}
}

func (this *activityModule) OnCheckCondition(role IRole, act_type int32, need_notify bool, args ...interface{}) {
	if activitys, has := this.now_activity_type[act_type]; has {
		for k, _ := range activitys {
			activitys[k].onCheckCondition(role, need_notify, args...)
		}
	}
}
