package activity

import (
	"Gameserver/global"
	. "Gameserver/logic"
	"common"
	. "common/cache"
	"common/protocol"
	"fmt"
	. "galaxy"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	ActivityTypeOrder = iota + 1
	ActivityTypeLevelBag
	ActivityTypeGrowFund
	ActivityTypeTotalPay
)

type iactivity interface {
	init(list_data *activityListData, data interface{})
	initRelativeTime(now *time.Time)
	onOpen(now *time.Time)
	onUpdate(now *time.Time)
	onClose(now *time.Time)
	onInitUserData(role IRole)
	onCheckCondition(role IRole, need_notity bool, args ...interface{})
	onQuery(role IRole) (int_value []int32, award_box_info []*protocol.ActivityBoxInfo)
	onOperate(role IRole, index int32) (retcode common.RetCode, award_info []*protocol.AwardInfo)
	getId() int32
	getActId() int32
	getActType() int32
	getListData() *activityListData
}

type activityBase struct {
	list_data *activityListData
}

func (this *activityBase) getId() int32 {
	return this.list_data.id
}

func (this *activityBase) getActId() int32 {
	return this.list_data.ActId
}

func (this *activityBase) getActType() int32 {
	return this.list_data.ActType
}

func (this *activityBase) getListData() *activityListData {
	return this.list_data
}

func (this *activityBase) initRelativeTime(now *time.Time) {
	switch this.list_data.TimeType {
	case common.ACTIVITY_TIMETYPE_WEEK:
		this.calWeekTimeTypeStartTime(now)
		LogDebug("calWeekTimeTypeStartTime ID: ", this.list_data.ActId, " ActivityStartTime:", this.list_data.StartTimestamp)
	}
}

func (this *activityBase) calWeekTimeTypeStartTime(now *time.Time) {
	var delta int32
	if int32(now.Weekday()) >= this.list_data.StartTime[0] {
		delta = int32(now.Weekday()) - this.list_data.StartTime[0]
	} else {
		delta = 7 - this.list_data.StartTime[0] + int32(now.Weekday())
	}

	this.list_data.StartTimestamp = time.Date(now.Year(), now.Month(), now.Day()-int(delta), 0, 0, 0, 0, now.Location()).Unix()
}

const (
	cache_activitylist_key_t    = "Role:%v:Activity"
	cache_activityobj_key_t     = "Role:%v:Activity:%v:Index:%v"
	cache_activityversion_key_t = "Role:%v:ActivityVersion"
)

func genActivityCacheKey(role_uid int64, id int32, index int32) string {
	return fmt.Sprintf(cache_activityobj_key_t, role_uid, id, index)
}

type ActivitySys struct {
	owner          IRole
	activity_list  map[int32]map[int32]*ActivityRoleCache
	cache_list_key string
}

func (this *ActivitySys) Init(owner IRole) {
	this.owner = owner
	this.activity_list = make(map[int32]map[int32]*ActivityRoleCache)
	this.cache_list_key = fmt.Sprintf(cache_activitylist_key_t, this.owner.GetUid())
}

func (this *ActivitySys) Check() {
	act_module.OnInitUserData(this.owner, ActivityTypeOrder)
}

func (this *ActivitySys) Load() error {
	resp, err := GxService().Redis().Cmd("SMEMBERS", this.cache_list_key)
	if err != nil {
		return err
	}

	cacheKeys, err := resp.List()
	if err != nil {
		return err
	}

	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", this.cache_list_key, key)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			cache := new(ActivityRoleCache)
			err = proto.Unmarshal(buf, cache)
			if err != nil {
				LogFatal(err)
				continue
			}

			if _, has := this.activity_list[cache.GetId()]; has {
				this.activity_list[cache.GetId()][cache.GetIndex()] = cache
			} else {
				this.activity_list[cache.GetId()] = make(map[int32]*ActivityRoleCache)
				this.activity_list[cache.GetId()][cache.GetIndex()] = cache
			}

		}
	}

	return nil
}

func (this *ActivitySys) Save(cache *ActivityRoleCache) {
	buf, err := proto.Marshal(cache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := genActivityCacheKey(this.owner.GetUid(), cache.GetId(), cache.GetIndex())
	if _, err := GxService().Redis().Cmd("SET", key, buf); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SADD", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
}

func (this *ActivitySys) Del(id int32, index int32) {
	key := genActivityCacheKey(this.owner.GetUid(), id, index)
	if _, err := GxService().Redis().Cmd("DEL", key); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SREM", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}

	if _, has := this.activity_list[id]; has {
		delete(this.activity_list[id], index)
		if len(this.activity_list[id]) <= 0 {
			delete(this.activity_list, id)
		}
	}
}

func (this *ActivitySys) ActGetStatus(id int32, index int32, is_clean int32, version int64) protocol.ActivityStatusFlag {
	if v, has := this.activity_list[id]; has {
		if d, has := v[index]; has {
			if is_clean != 0 && version != 0 && d.GetTimestamp() != version {
				this.activity_list[id][index].SetStatus(int32(protocol.ActivityStatusFlag_CanNotGet))
				this.activity_list[id][index].SetCondition(0)
				this.activity_list[id][index].SetTimestamp(0)
				this.activity_list[id][index].SetVersion(version)
				this.Save(this.activity_list[id][index])
			}
			return protocol.ActivityStatusFlag(this.activity_list[id][index].GetStatus())
		}
	}
	return protocol.ActivityStatusFlag_CanNotGet
}

func (this *ActivitySys) ActGetCondition(id int32, index int32, is_clean int32, version int64) int64 {
	if v, has := this.activity_list[id]; has {
		if d, has := v[index]; has {
			if is_clean != 0 && version != 0 && d.GetTimestamp() != version {
				LogDebug("ActGetCondition Enter Reset Id : ", id, " index : ", index, " version : ", version)
				this.activity_list[id][index].SetStatus(int32(protocol.ActivityStatusFlag_CanNotGet))
				this.activity_list[id][index].SetCondition(0)
				this.activity_list[id][index].SetTimestamp(0)
				this.activity_list[id][index].SetVersion(version)
				this.Save(this.activity_list[id][index])
			}
			return d.GetCondition()
		}
	}
	return 0
}

func (this *ActivitySys) ActGetTimestamp(id int32, index int32, is_clean int32, version int64) int64 {
	if v, has := this.activity_list[id]; has {
		if d, has := v[index]; has {
			if is_clean != 0 && version != 0 && d.GetTimestamp() != version {
				LogDebug("ActGetTimestamp Enter Reset Id : ", id, " index : ", index, " version : ", version)
				this.activity_list[id][index].SetStatus(int32(protocol.ActivityStatusFlag_CanNotGet))
				this.activity_list[id][index].SetCondition(0)
				this.activity_list[id][index].SetTimestamp(0)
				this.activity_list[id][index].SetVersion(version)
				this.Save(this.activity_list[id][index])
			}
			return d.GetTimestamp()
		}
	}
	return 0
}

func (this *ActivitySys) ActUpdateStatus(id int32, index int32, status protocol.ActivityStatusFlag, version int64) {
	now := time.Now().Unix()
	if v, has1 := this.activity_list[id]; has1 {
		if _, has2 := v[index]; has2 {
			this.activity_list[id][index].SetStatus(int32(status))
			this.activity_list[id][index].SetTimestamp(now)
			this.Save(this.activity_list[id][index])
		} else {
			cache := new(ActivityRoleCache)
			cache.SetId(id)
			cache.SetIndex(index)
			cache.SetStatus(int32(status))
			cache.SetCondition(0)
			cache.SetTimestamp(now)
			cache.SetVersion(version)
			this.activity_list[id][index] = cache
			this.Save(this.activity_list[id][index])
		}
	} else {
		this.activity_list[id] = make(map[int32]*ActivityRoleCache)
		cache := new(ActivityRoleCache)
		cache.SetId(id)
		cache.SetIndex(index)
		cache.SetStatus(int32(status))
		cache.SetCondition(0)
		cache.SetTimestamp(now)
		cache.SetVersion(version)
		this.activity_list[id][index] = cache
		this.Save(this.activity_list[id][index])
	}
}

func (this *ActivitySys) ActUpdateCondition(id int32, index int32, condition int64, version int64) {
	if v, has1 := this.activity_list[id]; has1 {
		if _, has2 := v[index]; has2 {
			this.activity_list[id][index].SetCondition(condition)
			this.Save(this.activity_list[id][index])
		} else {
			cache := new(ActivityRoleCache)
			cache.SetId(id)
			cache.SetIndex(index)
			cache.SetStatus(int32(protocol.ActivityStatusFlag_CanNotGet))
			cache.SetCondition(condition)
			cache.SetTimestamp(0)
			cache.SetVersion(version)
			this.activity_list[id][index] = cache
			this.Save(this.activity_list[id][index])
		}
	} else {
		this.activity_list[id] = make(map[int32]*ActivityRoleCache)
		cache := new(ActivityRoleCache)
		cache.SetId(id)
		cache.SetIndex(index)
		cache.SetStatus(int32(protocol.ActivityStatusFlag_CanNotGet))
		cache.SetCondition(condition)
		cache.SetTimestamp(0)
		cache.SetVersion(version)
		this.activity_list[id][index] = cache
		this.Save(this.activity_list[id][index])
	}
}

func sendAwardStatusNotify(role IRole, id int32, index int32, status protocol.ActivityStatusFlag) {
	msg := new(protocol.ActivityStatusInfo)
	msg.SetId(id)
	msg.SetIndex(index)
	msg.SetStatus(int32(status))
	msg.SetCondition(0)

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
	}
	global.SendMsg(int32(protocol.MsgCode_ActivityAwardStatusChangeNotify), role.GetSid(), buf)
}

func sendAwardStatusNotifyWithCondition(role IRole, id int32, index int32, status protocol.ActivityStatusFlag, condition int64) {
	msg := new(protocol.ActivityStatusInfo)
	msg.SetId(id)
	msg.SetIndex(index)
	msg.SetStatus(int32(status))
	msg.SetCondition(condition)

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
	}
	global.SendMsg(int32(protocol.MsgCode_ActivityAwardStatusChangeNotify), role.GetSid(), buf)
}
