package logger

import (
	"fmt"
	"runtime"
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

func Debug(mensage interface{}) {
	log(DEBUG, 0, mensage)
}

func Info(mensage interface{}) {
	log(INFO, Cyan, mensage)
}

func Error(mensage interface{}) {
	log(ERROR, Red, mensage)
}

func Warn(mensage interface{}) {
	log(WARN, Yellow, mensage)
}

func log(logType string, color int, mensage interface{}) {
	fmt.Printf("%s[%s]\t- [%s] %v %s \n", format(color), logType, file(), mensage, reset())
}

func file() string {
	_, file, _, _ := runtime.Caller(3)
	return file
}

func format(sequence int) string {
	return fmt.Sprintf("\x1b[%d;%dm", 0, sequence)
}

func reset() string {
	return format(Reset)
}
