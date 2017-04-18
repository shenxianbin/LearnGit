package event

import (
	. "galaxy/logs"
	"galaxy/utils"
	"math"
)

type event struct {
	function func(args ...interface{})
	args     []interface{}
}

type EventFlow struct {
	name      string
	eventChan chan *event
}

func NewEventFlow(name string) *EventFlow {
	ef := new(EventFlow)
	ef.name = name
	ef.eventChan = make(chan *event, math.MaxUint16)
	go ef.startEventFlow()
	return ef
}

func (this *EventFlow) startEventFlow() {
	GxLogDebug("Event:", this.name, "goroutine start")
	defer utils.Stack()
Exit:
	for {
		select {
		case e := <-this.eventChan:
			// GxLogDebug("EventFlow chan size : ", len(this.eventChan))
			if e == nil {
				break Exit
			}
			if e.function != nil {
				e.function(e.args...)
			}
		}
	}

	//	for e := range this.eventChan {
	//		GxLogDebug("EventFlow chan size : ", len(this.eventChan))
	//		if e.function != nil {
	//			e.function(e.args...)
	//		}
	//	}

	GxLogDebug("Event:", this.name, "goroutine end")
}

func (this *EventFlow) Close() {
	close(this.eventChan)
}

func (this *EventFlow) Execute(eventFunc func(args ...interface{}), eventArgs ...interface{}) {
	event := &event{function: eventFunc, args: eventArgs}
	this.eventChan <- event
}
