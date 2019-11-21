package logger

import (
	"fmt"
	"io"
	"runtime"
)

const LINE_TEMPLATE = "%s[%s]\t- [%s] %v %s \n"

type appender struct {
	output io.Writer
}

func (a *appender) log(logType string, color int, mensage interface{}) {
	fmt.Fprintf(
		a.output,
		LINE_TEMPLATE,
		format(color),
		logType,
		file(),
		mensage,
		reset(),
	)
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
