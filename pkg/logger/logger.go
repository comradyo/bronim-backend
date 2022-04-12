package logger

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"strings"
)

type Level string

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	ErrorLevel Level = "error"
	FatalLevel Level = "fatal"
)

var singletonLogger log.Logger

func Init(lvl Level) {
	singletonLogger = *log.New()
	level, _ := log.ParseLevel(string(lvl))
	singletonLogger.SetLevel(level)
}

func Debug(args ...interface{}) {
	singletonLogger.Debug(args...)
}

func Info(args ...interface{}) {
	singletonLogger.Info(args...)
}

func Error(args ...interface{}) {
	singletonLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	singletonLogger.Error(args...)
}

func DebugAtFunc(currFunc interface{}, args ...interface{}) {
	funcName := getFunctionName(currFunc)
	singletonLogger.Debug(funcName, ":", args)
}

func InfoAtFunc(currFunc interface{}, args ...interface{}) {
	funcName := getFunctionName(currFunc)
	singletonLogger.Info(funcName, ":", args)
}

func ErrorAtFunc(currFunc interface{}, args ...interface{}) {
	funcName := getFunctionName(currFunc)
	singletonLogger.Error(funcName, ":", args)
}

func FatalAtFunc(currFunc interface{}, args ...interface{}) {
	funcName := getFunctionName(currFunc)
	singletonLogger.Error(funcName, ":", args)
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
