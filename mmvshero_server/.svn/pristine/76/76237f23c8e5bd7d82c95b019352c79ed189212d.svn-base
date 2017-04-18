package logs

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

type colorBrush func(string) string

func newBrush(color string) colorBrush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colorsBrushes = []colorBrush{
	newBrush("1;37"), // Fatal	white
	newBrush("1;31"), // Error      red
	newBrush("1;33"), // Warning    yellow
	newBrush("1;32"), // Info	green
	newBrush("1;34"), // Debug      blue
}

type ConsoleWrite struct {
	*log.Logger
	Level int `json:"level"`
}

func NewConsole() LoggerInterface {
	cw := &ConsoleWrite{Level: LevelDebug}
	cw.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return cw
}

func (this *ConsoleWrite) Init(config string) error {
	if len(config) == 0 {
		return nil
	}
	return json.Unmarshal([]byte(config), this)
}

func (this *ConsoleWrite) WriteMsg(msg string, level int) error {
	if level > this.Level {
		return nil
	}

	if goos := runtime.GOOS; goos == "windows" {
		this.Println(msg)
		return nil
	}

	this.Println(colorsBrushes[level](msg))
	return nil
}

func (this *ConsoleWrite) Flush() {

}

func (this *ConsoleWrite) Destroy() {

}

func init() {
	register("console", NewConsole)
}
