package logs

import (
	"strings"
)

var debuglog *GxLogger

func SetGxLogLevel(level int) {
	debuglog.SetLevel(level)
}

func SetGxLogLogger(adapterName string, config string) {
	debuglog.SetLogger(adapterName, config)
}

func GxLogFatal(v ...interface{}) {
	debuglog.Fatal(generateFmtStr(len(v)), v...)
}

func GxLogError(v ...interface{}) {
	debuglog.Error(generateFmtStr(len(v)), v...)
}

func GxLogWarning(v ...interface{}) {
	debuglog.Warning(generateFmtStr(len(v)), v...)
}

func GxLogInfo(v ...interface{}) {
	debuglog.Info(generateFmtStr(len(v)), v...)
}

func GxLogDebug(v ...interface{}) {
	debuglog.Debug(generateFmtStr(len(v)), v...)
}

func init() {
	debuglog = NewLogger(1000)
	debuglog.SetFuncCallDepth(3)
	debuglog.SetLogger("console", "")
	debuglog.SetLogger("file", `{"filename":"./gxlog/gxlog"}`)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
