package logger

import "sync"

// Logger exposes all log functions.
type Logger interface {
	Info(msg string, fields Field)
	Debug(msg string, fields Field)
	Error(msg string, fields Field)
	Warn(msg string, fields Field)
	Fatal(msg string, fields Field)
}

//nolint
var (
	zlog    Logger
	zapOnce sync.Once
)

// The semi structured log fields.
type Field map[string]interface{}

// NewZapLogger returns a zap logger instance.
//nolint:ireturn // Return interface to protect zap methods.
func NewZapLogger(service string) Logger {
	z, err := newZap(service)
	if err != nil {
		panic(err)
	}

	return z
}

// NewZapGlobal returns a global zap logger instance.
// When called multiple times it will return same object with the
// service name of first caller.
// **This primarily to use the logger in code where we can't pass around
// logger from main function. Avoid globals:)**.
func NewZapGlobal(service string) Logger {
	if service == "" {
		panic(errEmptyService)
	}

	getZap := func() {
		var err error
		zlog, err = newZap(service)

		if err != nil {
			panic(err)
		}
	}
	zapOnce.Do(getZap)

	return zlog
}
