package logs

import (
	"fmt"
	"galaxy/utils"
	"path"
	"runtime"
	"sync"
)

const (
	LevelFatal = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

type LoggerInterface interface {
	Init(config string) error
	WriteMsg(msg string, level int) error
	Destroy()
	Flush()
}

type loggerType func() LoggerInterface

var adapter = make(map[string]loggerType)

func register(name string, log loggerType) {
	if log == nil {
		panic("logs : Register log provider is nil")
	}

	if _, has := adapter[name]; has {
		panic("logs: Register called twice for provider " + name)
	}

	adapter[name] = log
}

type logMsg struct {
	level int
	msg   string
}

type GxLogger struct {
	mutex               sync.Mutex
	level               int
	loggerFuncCallDepth int
	msgChan             chan *logMsg
	switchChan          chan bool
	outputs             map[string]LoggerInterface
}

func NewLogger(channelLen int) *GxLogger {
	gl := new(GxLogger)
	gl.level = LevelDebug
	gl.loggerFuncCallDepth = 2
	gl.msgChan = make(chan *logMsg, channelLen)
	gl.switchChan = make(chan bool)
	gl.outputs = make(map[string]LoggerInterface)
	go gl.startLogger()
	return gl
}

func (this *GxLogger) SetLogger(adapterName string, config string) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if logger, ok := adapter[adapterName]; ok {
		lg := logger()
		err := lg.Init(config)
		if err != nil {
			fmt.Println("logs.GxLogger.SetLogger: " + err.Error())
			return err
		}
		this.outputs[adapterName] = lg
	} else {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}

	return nil
}

func (this *GxLogger) DelLogger(adapterName string) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if lg, ok := this.outputs[adapterName]; ok {
		lg.Destroy()
		delete(this.outputs, adapterName)
		return nil
	} else {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}
}

func (this *GxLogger) writeMsg(level int, msg string) error {
	if level > this.level {
		return nil
	}

	lm := new(logMsg)
	lm.level = level
	_, file, line, ok := runtime.Caller(this.loggerFuncCallDepth)
	if ok {
		_, filename := path.Split(file)
		lm.msg = fmt.Sprintf("[%s:%d] %s", filename, line, msg)
	} else {
		lm.msg = msg
	}

	this.msgChan <- lm
	return nil
}

func (this *GxLogger) SetLevel(level int) {
	this.level = level
}

func (this *GxLogger) SetFuncCallDepth(depth int) {
	this.loggerFuncCallDepth = depth
}

func (this *GxLogger) startLogger() {
	defer utils.Stack()
Exit:
	for {
		select {
		case lm := <-this.msgChan:
			for _, lg := range this.outputs {
				err := lg.WriteMsg(lm.msg, lm.level)
				if err != nil {
					fmt.Println("ERROR, unable to WriteMsg:", err)
				}
			}
		case <-this.switchChan:
			this.close()
			break Exit
		}
	}
}

func (this *GxLogger) Flush() {
	for _, l := range this.outputs {
		l.Flush()
	}
}

func (this *GxLogger) Close() {
	close(this.switchChan)
}

func (this *GxLogger) close() {
	for {
		if len(this.msgChan) > 0 {
			lm := <-this.msgChan
			for _, l := range this.outputs {
				err := l.WriteMsg(lm.msg, lm.level)
				if err != nil {
					fmt.Println("ERROR, unable to WriteMsg (while closing logger):", err)
				}
			}
			continue
		}
		break
	}

	for _, l := range this.outputs {
		l.Flush()
		l.Destroy()
	}
}

func (this *GxLogger) Fatal(format string, args ...interface{}) {
	msg := fmt.Sprintf("[F]"+format, args...)
	this.writeMsg(LevelFatal, msg)
}

func (this *GxLogger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf("[E]"+format, args...)
	this.writeMsg(LevelError, msg)
}

func (this *GxLogger) Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf("[W]"+format, args...)
	this.writeMsg(LevelWarning, msg)
}

func (this *GxLogger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf("[I]"+format, args...)
	this.writeMsg(LevelInfo, msg)
}

func (this *GxLogger) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf("[D]"+format, args...)
	this.writeMsg(LevelDebug, msg)
}
