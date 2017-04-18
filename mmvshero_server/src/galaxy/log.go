package galaxy

import (
	"galaxy/logs"
	"strings"
)

var logger *logs.GxLogger

func LogLevel(level int) {
	logger.SetLevel(level)
}

func LogSetLogger(adaptername string, config string) error {
	err := logger.SetLogger(adaptername, config)
	if err != nil {
		return err
	}
	return nil
}

func LogFatal(v ...interface{}) {
	logger.Fatal(generateFmtStr(len(v)), v...)
}

func LogError(v ...interface{}) {
	logger.Error(generateFmtStr(len(v)), v...)
}

func LogWarning(v ...interface{}) {
	logger.Warning(generateFmtStr(len(v)), v...)
}

func LogInfo(v ...interface{}) {
	logger.Info(generateFmtStr(len(v)), v...)
}

func LogDebug(v ...interface{}) {
	logger.Debug(generateFmtStr(len(v)), v...)
}

func init() {
	logger = logs.NewLogger(10000)
	logger.SetFuncCallDepth(3)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
