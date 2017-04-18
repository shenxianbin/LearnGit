package utils

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func Stack() {
	if r := recover(); r != nil {
		now := time.Now().Unix()
		file, _ := os.Create(fmt.Sprintf("panic_%v", now))
		buf := make([]byte, 102400)
		l := runtime.Stack(buf, true)
		output := fmt.Sprintf("%s\n%s", r, string(buf[:l]))
		file.Write([]byte(output))
		file.Close()
		os.Exit(1)
	}
}
