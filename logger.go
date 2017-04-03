package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	Appenders []*Appender
}

func New() *Logger {
	file, err := os.Create("developement.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Logger: %v", err)
	}
	fileAppender := &Appender{Output: file}
	console := &Appender{Output: os.Stdout}
	appenders := []*Appender{fileAppender, console}
	return &Logger{Appenders: appenders}
}

func (l *Logger) Add(appender *Appender) {
	l.Appenders = append(l.Appenders, appender)
}

func (l *Logger) Debug(mensage interface{}) {
	for _, appender := range l.Appenders {
		appender.log(DEBUG, Reset, mensage)
	}
}

func (l *Logger) Info(mensage interface{}) {
	for _, appender := range l.Appenders {
		appender.log(INFO, Cyan, mensage)
	}
}

func (l *Logger) Error(mensage interface{}) {
	for _, appender := range l.Appenders {
		appender.log(ERROR, Red, mensage)
	}
}

func (l *Logger) Warn(mensage interface{}) {
	for _, appender := range l.Appenders {
		appender.log(ERROR, Red, mensage)
	}
}
