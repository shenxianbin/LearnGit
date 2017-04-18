package timer

/*FixTimer 固定时间定时器
用于每年，每月，每周，每天, 每小时触发
允许一个偏移值，支持类似每天5点重置的机制
*/

import (
	. "galaxy/event"
	"galaxy/utils"
	"sync"
	"time"
)

type FixedTimerType int

const (
	FIXED_TIMER_TYPE_YEAR FixedTimerType = iota + 1
	FIXED_TIMER_TYPE_MONTH
	FIXED_TIMER_TYPE_WEEK
	FIXED_TIMER_TYPE_DAY
)

type fixedTime struct {
	fixed_time_type  FixedTimerType
	fixed_time_event func()
}

type fixedTimer struct {
	timer_list   map[FixedTimerType][]*fixedTime
	timer_offset int64
	mutex        sync.Mutex
	exit_chan    chan bool
	wait_group   sync.WaitGroup

	year_flag       int
	yeat_timestamp  int64
	month_flag      int
	month_timestamp int64
	week_flag       int
	week_timestamp  int64
	day_flag        int
	day_timestamp   int64
}

func (this *fixedTimer) init(offset int64) {
	this.timer_list = make(map[FixedTimerType][]*fixedTime)
	this.timer_list[FIXED_TIMER_TYPE_YEAR] = make([]*fixedTime, 0)
	this.timer_list[FIXED_TIMER_TYPE_MONTH] = make([]*fixedTime, 0)
	this.timer_list[FIXED_TIMER_TYPE_WEEK] = make([]*fixedTime, 0)
	this.timer_list[FIXED_TIMER_TYPE_DAY] = make([]*fixedTime, 0)

	this.timer_offset = offset
	this.exit_chan = make(chan bool)

	fix_time := time.Now().Unix() - this.timer_offset
	t := time.Unix(fix_time, 0)
	this.year_flag = t.Year()
	this.month_flag = int(t.Month())
	this.week_flag = int(t.Weekday())
	this.day_flag = t.Day()
}

func (this *fixedTimer) run() {
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

func (this *fixedTimer) execute() {
	fix_time := time.Now().Unix() - this.timer_offset
	t := time.Unix(fix_time, 0)
	now_year := t.Year()
	now_month := int(t.Month())
	now_week := int(t.Weekday())
	now_day := t.Day()

	switch {
	case this.year_flag != now_year:
		for _, v := range this.timer_list[FIXED_TIMER_TYPE_YEAR] {
			GxEvent().Execute(func(args ...interface{}) {
				v.fixed_time_event()
			})
		}
		this.year_flag = now_year
	case this.month_flag != now_month:
		for _, v := range this.timer_list[FIXED_TIMER_TYPE_MONTH] {
			GxEvent().Execute(func(args ...interface{}) {
				v.fixed_time_event()
			})
		}
		this.month_flag = now_month
	case this.week_flag == 1 && now_week == 0:
		for _, v := range this.timer_list[FIXED_TIMER_TYPE_WEEK] {
			GxEvent().Execute(func(args ...interface{}) {
				v.fixed_time_event()
			})
		}
		this.week_flag = now_week
	case this.day_flag != now_day:
		for _, v := range this.timer_list[FIXED_TIMER_TYPE_DAY] {
			GxEvent().Execute(func(args ...interface{}) {
				v.fixed_time_event()
			})
		}
		this.day_flag = now_day
	}
}

func (this *fixedTimer) add(fixed_type FixedTimerType, event func()) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	list, has := this.timer_list[fixed_type]
	if has {
		this.timer_list[fixed_type] = append(list, &fixedTime{
			fixed_time_type:  fixed_type,
			fixed_time_event: event,
		})
	}
}

func (this *fixedTimer) stop() {
	close(this.exit_chan)
}

func (this *fixedTimer) wait() {
	this.wait_group.Wait()
}
