package soldier

import (
	. "Gameserver/cache"
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	d "common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	"galaxy"
	"github.com/golang/protobuf/proto"
	"sort"
)

//地图事件
type SoldierEvent struct {
	eventId    int32 //自增id
	startTime  int64 //发生的时间
	endTime    int64 //结束的时间
	duration   int64 //动作持续时间,秒 A
	eventType  int32 //事件类型， 1恶斗，2生蛋，3生病
	remote     int32 //A soldier id in map
	local      int32 //B 原地
	levelTotal int32 //A和B的等级和，生成事件的时候计算
}

//单种事件列表
type SoldierEvents struct {
	eventList map[int32]*SoldierEvent //int32 事件类型 从1开始 1恶斗，2生蛋，3生病
}

const (
	eventKey       = "Role:%v:SoldierEvent:%v"
	eventKeys      = "SoldierEventKeys:%v"
	eventKeyAutoId = "Role:%v:SoldierEventAutoId"
)

func (this *UserSoldiers) getEventkey(roleId int64, eventId int32) string {
	return fmt.Sprintf(eventKey, roleId, eventId)
}

func (this *UserSoldiers) getEventkeys(roleId int64) string {
	return fmt.Sprintf(eventKeys, roleId)
}

func (this *UserSoldiers) getEventAutoId() (int32, error) {
	resp, err := galaxy.GxService().Redis().Cmd("INCR", fmt.Sprintf(eventKeyAutoId, this.user.GetUid()))
	if err != nil {
		galaxy.LogError(err)
		return 0, err
	}

	temp, err := resp.Int64()
	if err != nil {
		galaxy.LogError(err)
		return 0, err
	}
	return int32(temp), err
}

func (this *UserSoldiers) loadEventsFromDb() error {
	roleId := this.user.GetUid()
	//获取所有keys
	resp, err := galaxy.GxService().Redis().Cmd("SMEMBERS", this.getEventkeys(roleId))
	if err != nil {
		galaxy.LogError(err)
		return err //没有
	}

	keys, _ := resp.List()

	for _, key := range keys {
		//获取数据
		resp, err := galaxy.GxService().Redis().Cmd("GET", key)
		if err != nil {
			galaxy.LogError(err)
			return err //没有
		}
		if buff, _ := resp.Bytes(); buff != nil {
			soldierEventCache := &SoldierEventCache{}
			err = proto.Unmarshal(buff, soldierEventCache)
			soldierEvent := &SoldierEvent{}
			//copy
			soldierEvent.eventId = soldierEventCache.GetEventId()
			soldierEvent.startTime = soldierEventCache.GetStartTime()
			soldierEvent.endTime = soldierEventCache.GetEndTime()
			soldierEvent.duration = soldierEventCache.GetDuration()
			soldierEvent.eventType = soldierEventCache.GetEventType()
			soldierEvent.remote = soldierEventCache.GetRemote()
			soldierEvent.local = soldierEventCache.GetLocal()
			soldierEvent.levelTotal = soldierEventCache.GetLevelTotal()

			if this.Events[soldierEventCache.GetEventType()].eventList == nil {
				this.Events[soldierEventCache.GetEventType()].eventList = make(map[int32]*SoldierEvent)
			}

			this.Events[soldierEventCache.GetEventType()].eventList[soldierEventCache.GetEventId()] = soldierEvent

		} else {
			galaxy.LogError("SREM: ", this.getEventkeys(roleId), key)
			galaxy.GxService().Redis().Cmd("SREM", this.getEventkeys(roleId), key)
		}
	}

	if err != nil {
		galaxy.LogError(err)
	}
	return err
}

func (this *UserSoldiers) SoldierAllEvents() *protocol.AllSoldiersEvents {
	var i int32
	events1 := make([]*protocol.SoldierEvent, len(this.Events[1].eventList))

	i = 0
	for _, v := range this.Events[1].eventList {
		events1[i] = v.FillEvent()
		i++
	}

	events2 := make([]*protocol.SoldierEvent, len(this.Events[2].eventList))

	i = 0
	for _, v := range this.Events[2].eventList {
		events2[i] = v.FillEvent()
		i++
	}

	events3 := make([]*protocol.SoldierEvent, len(this.Events[3].eventList))

	i = 0
	for _, v := range this.Events[2].eventList {
		events3[i] = v.FillEvent()
		i++
	}

	m := new(protocol.AllSoldiersEvents)
	m.SetEvents1(events1)
	m.SetEvents2(events2)
	m.SetEvents3(events3)
	return m
}

func (this *SoldierEvent) FillEvent() *protocol.SoldierEvent {
	m := &protocol.SoldierEvent{}
	m.SetEventId(this.eventId)
	m.SetStartTime(this.startTime)
	m.SetEndTime(this.endTime)
	m.SetDuration(this.duration)
	m.SetEventType(this.eventType)
	m.SetRemote(this.remote)
	m.SetLocal(this.local)
	m.SetLevelTotal(this.levelTotal)
	return m
}

func (this *UserSoldiers) newEventCache(event *SoldierEvent) *SoldierEventCache {
	soldierEventCache := &SoldierEventCache{}
	soldierEventCache.SetEventId(event.eventId)
	soldierEventCache.SetStartTime(event.startTime)
	soldierEventCache.SetEndTime(event.endTime)
	soldierEventCache.SetDuration(event.duration)
	soldierEventCache.SetEventType(event.eventType)
	soldierEventCache.SetRemote(event.remote)
	soldierEventCache.SetLocal(event.local)
	soldierEventCache.SetLevelTotal(event.levelTotal)
	return soldierEventCache
}

func (this *UserSoldiers) saveEventsToDb(eventType, eventId int32) bool {
	if event, ok := (*this).Events[eventType].eventList[eventId]; ok {
		soldierEventCache := this.newEventCache(event)

		buff, err := proto.Marshal(soldierEventCache)
		if err != nil {
			galaxy.LogError(err)
			return false
		}
		if _, err = galaxy.GxService().Redis().Cmd("SET", this.getEventkey(this.user.GetUid(), event.eventId), buff); err != nil {
			galaxy.LogError(err)
			return false
		}

		if _, err := galaxy.GxService().Redis().Cmd("SADD", this.getEventkeys(this.user.GetUid()), this.getEventkey(this.user.GetUid(), event.eventId)); err != nil {
			galaxy.LogError(err)
			return false
		}

		return true
	}
	return false
}

func (this *UserSoldiers) removeEventsFromDb(eventType, eventId int32) bool {
	if event, ok := (*this).Events[eventType].eventList[eventId]; ok {
		if _, err := galaxy.GxService().Redis().Cmd("DEL", this.getEventkey(this.user.GetUid(), event.eventId)); err != nil {
			galaxy.LogError(err)
			return false
		}
		if _, err := galaxy.GxService().Redis().Cmd("SREM", this.getEventkeys(this.user.GetUid()), this.getEventkey(this.user.GetUid(), event.eventId)); err != nil {
			galaxy.LogError(err)
			return false
		}
		delete(this.Events[eventType].eventList, eventId)
		return true
	}
	return false
}

//事件处理
func (this *UserSoldiers) SoldierEventDealWith(eventType, eventId int32) bool {
	events, ok := this.Events[eventType]
	if !ok {
		return false
	}
	event, ok := events.eventList[eventId]
	if !ok {
		return false
	}

	if event.startTime > Time() {
		return false //还没开始
	}

	if eventType == 1 {
		//进行中取消事件，结束后移除尸体
		this.RemoveEvents(eventType, eventId)
		if event.startTime+event.duration <= Time() {
			//已经结束，清理尸体
			e := this.Events[eventType].eventList[eventId]
			this.DeleteSoldierInMapFromDb(e.local)
		}
		return true

	} else if eventType == 2 {
		//直接获取奖励
		Event2Award1 := scheme.Commonmap[d.Event2Award1].Value
		Event2Award2 := scheme.Commonmap[d.Event2Award2].Value
		Event2Award3 := scheme.Commonmap[d.Event2Award3].Value
		Event2Award4 := scheme.Commonmap[d.Event2Award4].Value
		Event2Award5 := scheme.Commonmap[d.Event2Award5].Value

		Event2Award1Param := scheme.Commonmap[d.Event2Award1Param].Value
		Event2Award2Param := scheme.Commonmap[d.Event2Award2Param].Value
		Event2Award3Param := scheme.Commonmap[d.Event2Award3Param].Value
		Event2Award4Param := scheme.Commonmap[d.Event2Award4Param].Value
		Event2Award5Param := scheme.Commonmap[d.Event2Award5Param].Value

		if event.levelTotal <= Event2Award1Param {
			Award(Event2Award1, this.user, true)
		} else if Event2Award1Param < event.levelTotal && event.levelTotal <= Event2Award2Param {
			Award(Event2Award2, this.user, true)
		} else if Event2Award2Param < event.levelTotal && event.levelTotal <= Event2Award3Param {
			Award(Event2Award3, this.user, true)
		} else if Event2Award3Param < event.levelTotal && event.levelTotal <= Event2Award4Param {
			Award(Event2Award4, this.user, true)
		} else if Event2Award4Param < event.levelTotal && event.levelTotal <= Event2Award5Param {
			Award(Event2Award5, this.user, true)
		} else {
			return false
		}

		this.RemoveEvents(eventType, eventId)
	} else {
		//好友系统暂时无法处理
	}
	return false
}

//移除这个事件，以及之后发生的事件
func (this *UserSoldiers) RemoveEvents(eventType, eventId int32) {
	for k, _ := range this.Events[eventType].eventList {
		if k >= eventId {
			this.removeEventsFromDb(eventType, eventId)
		}
	}
}

func (this *UserSoldiers) GenerateEvents() bool {
	Event1TimeDown := int64(scheme.Commonmap[d.Event1TimeDown].Value)
	Event1TimeUp := int64(scheme.Commonmap[d.Event1TimeUp].Value)
	Event1CorpseLimit := scheme.Commonmap[d.Event1CorpseLimit].Value
	Event1Duration := scheme.Commonmap[d.Event1Duration].Value
	Event2TimeDown := int64(scheme.Commonmap[d.Event2TimeDown].Value)
	Event2TimeUp := int64(scheme.Commonmap[d.Event2TimeUp].Value)

	Event2Duration := scheme.Commonmap[d.Event2Duration].Value
	Event3TimeDown := int64(scheme.Commonmap[d.Event3TimeDown].Value)
	Event3TimeUp := int64(scheme.Commonmap[d.Event3TimeUp].Value)

	if this.Events == nil {
		this.Events = make(map[int32]*SoldierEvents)
	}

	//用于标记状态 是否可用,1暂时不可用，2永久不可用
	soldiers := make(map[int32]int32)
	//init soldiers
	for autoId, _ := range this.SoldiersInMap {
		soldiers[autoId] = 0
	}

	//设置时间游标
	thatTime := Time()
	//事件生成循环退出条件: 地图中无可用魔物
	for {
		//取出地图中可以生成事件的魔物列表
		for eventType, events := range this.Events {
			fmt.Println(eventType)
			for _, event := range events.eventList {
				if event.startTime > thatTime {
					continue
				} else if event.startTime <= thatTime {
					soldiers[event.local] = 2
					if event.endTime > thatTime {
						//doing event
						if event.remote != 0 {
							if event.duration+event.endTime > thatTime {
								soldiers[event.remote] = 1
							} else {
								soldiers[event.remote] = 0
							}
						}

					} else if event.endTime <= thatTime {
						//overt
						soldiers[event.remote] = 0
					}
				}
			}
		}

		soldiersValid := make(map[int32]int32)
		for k, v := range soldiers {
			if v == 0 {
				soldiersValid[k] = v
			}
		}

		if len(soldiersValid) == 0 {
			break
		}

		//生成事件
		//获取3个事件中最快要结束的事件
		eventsTime := make(map[int32]int64)
		eventLastKey := [3]int32{0, 0, 0}

		for i := 1; i <= 3; i++ {
			if v, ok := this.Events[int32(i)]; ok {
				var keys []int
				for k := range v.eventList {
					keys = append(keys, int(k))
				}
				sort.Ints(keys)
				eventLastKey[i] = int32(keys[len(keys)-1])

				eventsTime[int32(i)] = v.eventList[eventLastKey[i]].endTime
				if v.eventList[eventLastKey[i]].endTime < thatTime {
					eventsTime[int32(i)] = thatTime
				}
			} else {
				eventLastKey[i] = 0
				eventsTime[int32(i)] = thatTime
			}
		}

		Event1TimeLength := Rand(Event1TimeDown, Event1TimeUp)
		Event2TimeLength := Rand(Event2TimeDown, Event2TimeUp)
		Event3TimeLength := Rand(Event3TimeDown, Event3TimeUp)

		eventsTime[1] += int64(Event1TimeLength)
		eventsTime[2] += int64(Event2TimeLength)
		eventsTime[3] += int64(Event3TimeLength)

		newEvent := new(SoldierEvent)
		var joinedSoldiers []int32
		//随机取出两个 最多两个
		for k, _ := range soldiersValid {
			joinedSoldiers = append(joinedSoldiers, k)
			if len(joinedSoldiers) == 2 {
				break
			}
		}
		var err error
		newEvent.eventId, err = this.getEventAutoId()
		if err != nil {
			galaxy.LogError(err)
			return false
		}
		newEvent.local = joinedSoldiers[0]

		if len(soldiersValid) >= 1 && int(Event1CorpseLimit) > len(this.Events[1].eventList) && eventsTime[1] <= eventsTime[2] && eventsTime[1] <= eventsTime[3] {
			//1 生成事件1 恶斗
			newEvent.eventType = 1
			newEvent.duration = int64(Event1Duration)
			newEvent.startTime = eventsTime[1] - int64(Event1TimeLength)
			newEvent.endTime = eventsTime[1]
			newEvent.remote = joinedSoldiers[1]
			thatTime += newEvent.duration
		} else if len(soldiersValid) >= 1 && eventsTime[2] <= eventsTime[1] && eventsTime[2] <= eventsTime[3] {
			//2 生蛋

			newEvent.eventType = 2
			newEvent.remote = joinedSoldiers[1]
			newEvent.duration = int64(Event2Duration)
			newEvent.startTime = eventsTime[2] - int64(Event2TimeLength)
			newEvent.endTime = eventsTime[2]

			soldierLocal, ok := this.SoldiersInMap[newEvent.local]
			if !ok {
				return false
			}

			soldierRemote, ok := this.SoldiersInMap[newEvent.remote]
			if !ok {
				return false
			}

			newEvent.levelTotal = soldierLocal.GetLevel() + soldierRemote.GetLevel()
			thatTime += newEvent.duration
		} else {
			//3 生病
			//随机取出1个
			newEvent.eventType = 3
			newEvent.duration = 0
			newEvent.startTime = eventsTime[3] - int64(Event3TimeLength)
			newEvent.endTime = eventsTime[3]
			thatTime += int64(Min(int32(Event1TimeLength), int32(Event2TimeLength), int32(Event3TimeLength)))
		}
		if this.Events[newEvent.eventType].eventList == nil {
			this.Events[newEvent.eventType].eventList = make(map[int32]*SoldierEvent)
		}

		this.Events[newEvent.eventType].eventList[eventLastKey[newEvent.eventType]+1] = newEvent
	}
	return true
}
