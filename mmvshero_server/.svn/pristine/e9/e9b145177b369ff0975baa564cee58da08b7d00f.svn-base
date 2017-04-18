package timer

var timer_manager *timerManager

func init() {
	timer_manager = new(timerManager)
}

func Start(timer_type TimerType, args ...interface{}) {
	timer_manager.start(timer_type, args...)
}

func AddFixedTimerEvent(fixed_type FixedTimerType, event func()) {
	timer_manager.addFixedTimerEvent(fixed_type, event)
}

func AddCdTimerEvent(interval int32, count int32, event func()) {
	timer_manager.addCDTimerEvent(interval, count, event)
}

func Stop() {
	timer_manager.stop()
}

func Wait() {
	timer_manager.wait()
}
