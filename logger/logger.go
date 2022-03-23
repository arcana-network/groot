package logger

// Logger exposes all log functions.
type Logger interface {
	Info(msg string, fields Field)
	Debug(msg string, fields Field)
	Error(msg string, fields Field)
	Warn(msg string, fields Field)
	Fatal(msg string, fields Field)
}

// The semi structured log fields.
type Field map[string]interface{}

//nolint:ireturn // Return interface to protect zap methods.
func NewZapLogger(service string) Logger {
	z, err := newZap(service)
	if err != nil {
		panic(err)
	}

	return z
}
