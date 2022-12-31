package core

import (
	"io"
)

type ILogger interface {
	CommonLogger
	FormatLogger
	Config
}

type Config interface {
	SetLevel(lv Level)
	SetOutput(w io.Writer)
}

type CommonLogger interface {
	Trace(a ...any)
	Debug(a ...any)
	Info(a ...any)
	Warn(a ...any)
	Error(a ...any)
	Fatal(a ...any)
}

type FormatLogger interface {
	Tracef(format string, a ...any)
	Debugf(format string, a ...any)
	Infof(format string, a ...any)
	Warnf(format string, a ...any)
	Errorf(format string, a ...any)
	Fatalf(format string, a ...any)
}
