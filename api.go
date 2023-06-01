package inquisitor

import "io"

func SetLevel(lv Level) {
	logger.SetLevel(lv)
}

func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

//////////////////
// CommonLogger //
//////////////////

func Trace(a ...any) {
	logger.Trace(a...)
}

func Debug(a ...any) {
	logger.Debug(a...)
}

func Info(a ...any) {
	logger.Info(a...)
}

func Warn(a ...any) {
	logger.Warn(a...)
}

func Error(a ...any) {
	logger.Error(a...)
}

func Fatal(a ...any) {
	logger.Fatal(a...)
}

//////////////////
// FormatLogger //
//////////////////

func Tracef(format string, a ...any) {
	logger.Tracef(format, a...)
}

func Debugf(format string, a ...any) {
	logger.Debugf(format, a...)
}

func Infof(format string, a ...any) {
	logger.Infof(format, a...)
}

func Warnf(format string, a ...any) {
	logger.Warnf(format, a...)
}

func Errorf(format string, a ...any) {
	logger.Errorf(format, a...)
}

func Fatalf(format string, a ...any) {
	logger.Fatalf(format, a...)
}
