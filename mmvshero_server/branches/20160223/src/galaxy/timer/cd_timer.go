package timer

import (
	. "galaxy/event"
	"galaxy/utils"
	"sync"
	"time"
)

const CDTimeNoCount = -1

type cdTime struct {
	timestamp     int64
	interval      int32
	count         int32
	cd_time_event func()
}

type cdTimer struct {
	timer_list  map[int64]*cdTime
	mutex       sync.Mutex
	exit_chan   chan bool
	wait_group  sync.WaitGroup
	timer_index int64
}

func (this *cdTimer) init() {
	this.timer_list = make(map[int64]*cdTime)
	this.exit_chan = make(chan bool)
}

func (this *cdTimer) run() {
	this.wait_group.Add(1)
	go func() {
		defer utils.Stack()
		timer := time.NewTicker(time.Second)
	Exit:
		for {
			select {
			case <-timer.C:
				this.mutex.Lock()
				this.execute()
				this.mutex.Unlock()
			case <-this.exit_chan:
				break Exit
			}
		}
		this.wait_group.Done()
	}()
}

func (this *cdTimer) execute() {
	now := time.Now().Unix()
	for k, v := range this.timer_list {
		if now >= v.timestamp {
			if v.count != CDTimeNoCount {
				v.count--
				GxEvent().Execute(func(args ...interface{}) {
					v.cd_time_event()
				})
				if v.count <= 0 {
					delete(this.timer_list, k)
				} else {
					v.timestamp = now + int64(v.interval)
				}
			} else {
				GxEvent().Execute(func(args ...interface{}) {
					v.cd_time_event()
				})
				v.timestamp = now + int64(v.interval)
			}
		}
	}
}

func (this *cdTimer) add(interval int32, count int32, event func()) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.timer_index++
	this.timer_list[this.timer_index] = &cdTime{
		timestamp:     time.Now().Unix(),
		interval:      interval,
		count:         count,
		cd_time_event: event,
	}
}

func (this *cdTimer) stop() {
	close(this.exit_chan)
}

func (this *cdTimer) wait() {
	this.wait_group.Wait()
}
