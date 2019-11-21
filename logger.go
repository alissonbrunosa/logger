package logger

import (
	"io"
	"os"
)

const (
	Cyan   = 36
	Yellow = 33
	Red    = 31
	Reset  = 0
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

var Default = &Logger{
	appenders: []*appender{
		{output: os.Stdout},
	},
}

type Logger struct {
	appenders []*appender
}

func New(outputs ...io.Writer) *Logger {
	var appenders []*appender

	for _, output := range outputs {
		appenders = append(appenders, &appender{output: output})
	}

	return &Logger{appenders: appenders}
}

func (l *Logger) Add(appender *appender) {
	l.appenders = append(l.appenders, appender)
}

func (l *Logger) Debug(mensage interface{}) {
	for _, appender := range l.appenders {
		appender.log(DEBUG, Reset, mensage)
	}
}

func (l *Logger) Info(mensage interface{}) {
	for _, appender := range l.appenders {
		appender.log(INFO, Cyan, mensage)
	}
}

func (l *Logger) Error(mensage interface{}) {
	for _, appender := range l.appenders {
		appender.log(ERROR, Red, mensage)
	}
}

func (l *Logger) Warn(mensage interface{}) {
	for _, appender := range l.appenders {
		appender.log(WARN, Yellow, mensage)
	}
}
