package timer

import (
	. "galaxy/logs"
	"sync"
)

type TimerType int

const (
	TIMER_TYPE_FIXED TimerType = 1 + iota
	TIMER_TYPE_CD
)

type timerManager struct {
	fixer_timer       map[int64]*fixedTimer
	fixer_timer_mutex sync.Mutex

	cd_timer       cdTimer
	cd_timer_state bool
	cd_timer_mutex sync.Mutex
}

func (this *timerManager) start(timer_type TimerType, args ...interface{}) {
	switch timer_type {
	case TIMER_TYPE_FIXED:
		this.fixer_timer_mutex.Lock()
		if this.fixer_timer == nil {
			this.fixer_timer = make(map[int64]*fixedTimer)
		}

		if _, has := this.fixer_timer[args[0].(int64)]; !has {
			timer := new(fixedTimer)
			timer.init(args[0].(int64))
			timer.run()
			this.fixer_timer[args[0].(int64)] = timer
			GxLogInfo("FixedTimer Start Success ! Ex: ", args[0].(int64))
		}
		this.fixer_timer_mutex.Unlock()
	case TIMER_TYPE_CD:
		this.cd_timer_mutex.Lock()
		if !this.cd_timer_state {
			this.cd_timer.init()
			this.cd_timer.run()
			this.cd_timer_state = true
			GxLogInfo("CdTimer Start Success !")
		}
		this.cd_timer_mutex.Unlock()
	}
}

func (this *timerManager) addFixedTimerEvent(fixed_type FixedTimerType, offset int64, event func()) {
	this.fixer_timer_mutex.Lock()
	defer this.fixer_timer_mutex.Unlock()
	if timer, has := this.fixer_timer[offset]; has {
		timer.add(fixed_type, event)
	}
}

func (this *timerManager) addCDTimerEvent(interval int32, count int32, event func()) {
	this.cd_timer_mutex.Lock()
	defer this.cd_timer_mutex.Unlock()
	if this.cd_timer_state {
		this.cd_timer.add(interval, count, event)
	}
}

func (this *timerManager) stop() {
	this.fixer_timer_mutex.Lock()
	for _, v := range this.fixer_timer {
		v.stop()
	}

	this.fixer_timer_mutex.Unlock()

	this.cd_timer_mutex.Lock()
	if this.cd_timer_state {
		this.cd_timer.stop()
		this.cd_timer_state = false
	}
	this.cd_timer_mutex.Unlock()
}

func (this *timerManager) wait() {
	for _, v := range this.fixer_timer {
		v.wait()
	}
	this.fixer_timer = make(map[int64]*fixedTimer)
	this.cd_timer.wait()
}
