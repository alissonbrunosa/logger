package logger

import (
	"fmt"
	"io"
	"runtime"
)

const (
	Cyan   = 36
	Yellow = 33
	Red    = 31
	Reset  = 0
)

type Appender struct {
	Output io.Writer
}

func (a *Appender) log(logType string, color int, mensage interface{}) {
	fmt.Fprintf(a.Output, "%s[%s]\t- [%s] %v %s \n", format(color), logType, file(), mensage, reset())
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
