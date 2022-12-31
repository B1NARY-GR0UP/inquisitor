package core

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const Inquisitor = "---INQUISITOR---"

var logger ILogger = &inquisitor{
	stdLog:    log.New(os.Stderr, Inquisitor, log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
	level:     LevelInfo,
	calldepth: 4,
}

type inquisitor struct {
	stdLog    *log.Logger
	level     Level
	calldepth int
}

func (i *inquisitor) log(lv Level, format *string, a ...any) {
	if i.level > lv {
		return
	}
	sb := &strings.Builder{}
	switch lv {
	case LevelTrace:
		sb.WriteString("[TRACE] ")
	case LevelDebug:
		sb.WriteString("[DEBUG] ")
	case LevelInfo:
		sb.WriteString("[INFO] ")
	case LevelWarn:
		sb.WriteString("[WARN] ")
	case LevelError:
		sb.WriteString("[ERROR] ")
	case LevelFatal:
		sb.WriteString("[FATAL] ")
	default:
		sb.WriteString(fmt.Sprintf("[?%d?] ", lv))
	}
	if format != nil {
		sb.WriteString(fmt.Sprintf(*format, a...))
	} else {
		sb.WriteString(fmt.Sprint(a...))
	}
	_ = i.stdLog.Output(i.calldepth, sb.String())
	if lv == LevelFatal {
		os.Exit(1)
	}
}

func (i *inquisitor) SetLevel(lv Level) {
	i.level = lv
}

func (i *inquisitor) SetOutput(w io.Writer) {
	i.stdLog.SetOutput(w)
}

func (i *inquisitor) Trace(a ...any) {
	i.log(LevelTrace, nil, a...)
}

func (i *inquisitor) Debug(a ...any) {
	i.log(LevelDebug, nil, a...)
}

func (i *inquisitor) Info(a ...any) {
	i.log(LevelInfo, nil, a...)
}

func (i *inquisitor) Warn(a ...any) {
	i.log(LevelWarn, nil, a...)
}

func (i *inquisitor) Error(a ...any) {
	i.log(LevelError, nil, a...)
}

func (i *inquisitor) Fatal(a ...any) {
	i.log(LevelFatal, nil, a...)
}

func (i *inquisitor) Tracef(format string, a ...any) {
	i.log(LevelTrace, &format, a...)
}

func (i *inquisitor) Debugf(format string, a ...any) {
	i.log(LevelDebug, &format, a...)
}

func (i *inquisitor) Infof(format string, a ...any) {
	i.log(LevelInfo, &format, a...)
}

func (i *inquisitor) Warnf(format string, a ...any) {
	i.log(LevelWarn, &format, a...)
}

func (i *inquisitor) Errorf(format string, a ...any) {
	i.log(LevelError, &format, a...)
}

func (i *inquisitor) Fatalf(format string, a ...any) {
	i.log(LevelFatal, &format, a...)
}
