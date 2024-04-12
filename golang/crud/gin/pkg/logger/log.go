package logger

import (
	"fmt"
	"jixiangup.me/examples/gin/internal/constants"
	"jixiangup.me/examples/gin/pkg/util/str"
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var logLevelNames = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

// Log 导出日志对象 默认日志级别为INFO
var Log *Logger

func SetupLogger() {
	loggerLevel := os.Getenv(constants.LogLevel)
	if loggerLevel == "" {
		loggerLevel = "INFO"
	}
	loggerLevel = strings.ToUpper(loggerLevel)
	contains, index := str.ContainsString(logLevelNames, loggerLevel)
	if contains {
		Log = NewLogger(LogLevel(index))
	} else {
		Log = NewLogger(INFO)
	}
}

type Logger struct {
	Level LogLevel
}

func NewLogger(level LogLevel) *Logger {
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[Examples - Gin] - ")
	return &Logger{Level: level}
}

// 输出日志信息，根据日志级别添加前缀
func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level < l.Level {
		return
	}
	prefix := fmt.Sprintf("[Examples - Gin] - %s ", logLevelNames[level])
	log.SetPrefix(prefix)
	log.Printf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
	os.Exit(1)
}
