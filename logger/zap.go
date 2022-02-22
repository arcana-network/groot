package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	*zap.Logger
}

// NewZap creates a new zap logger instance.
// Service could be gateway, uploader etc.
func NewZap(service string) Logger {

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:      "service",
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
		},
	}

	logger, err := cfg.Build(zap.AddCallerSkip(1)) // If we don't skip caller, it will always show the caller as this file
	if err != nil {
		log.Panic("Unable to build zap logger")
	}

	return &zapLogger{
		logger.Named(service),
	}
}

func (z *zapLogger) Info(msg string, fields Field) {
	z.Logger.Sugar().Infow(msg, unwrapFields(fields)...)
}

func (z *zapLogger) Debug(msg string, fields Field) {
	z.Logger.Sugar().Debugw(msg, unwrapFields(fields)...)
}

func (z *zapLogger) Error(msg string, fields Field) {
	z.Logger.Sugar().Errorw(msg, unwrapFields(fields)...)
}

func (z *zapLogger) Warn(msg string, fields Field) {
	z.Logger.Sugar().Warnw(msg, unwrapFields(fields)...)
}

func (z *zapLogger) Fatal(msg string, fields Field) {
	z.Logger.Sugar().Fatalw(msg, unwrapFields(fields)...)
}

func unwrapFields(fields Field) (unwrapped []interface{}) {
	for k, v := range fields {
		unwrapped = append(unwrapped, k, v)
	}
	return
}
