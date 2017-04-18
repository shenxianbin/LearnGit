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
	fixer_timer       fixedTimer
	fixer_timer_state bool
	fixer_timer_mutex sync.Mutex

	cd_timer       cdTimer
	cd_timer_state bool
	cd_timer_mutex sync.Mutex
}

func (this *timerManager) start(timer_type TimerType, args ...interface{}) {
	switch timer_type {
	case TIMER_TYPE_FIXED:
		this.fixer_timer_mutex.Lock()
		if !this.fixer_timer_state {
			this.fixer_timer.init(args[0].(int64))
			this.fixer_timer.run()
			this.fixer_timer_state = true
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

func (this *timerManager) addFixedTimerEvent(fixed_type FixedTimerType, event func()) {
	this.fixer_timer_mutex.Lock()
	defer this.fixer_timer_mutex.Unlock()
	if this.fixer_timer_state {
		this.fixer_timer.add(fixed_type, event)
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
	if this.fixer_timer_state {
		this.fixer_timer.stop()
		this.fixer_timer_state = false
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
	this.fixer_timer.wait()
	this.cd_timer.wait()
}
