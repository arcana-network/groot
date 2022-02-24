package logger

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	maxSize    = 512
	maxAge     = 90
	maxBackups = 30
)

// zapLogger wraps the zap logging library to satisfy Logger interface.
type zapLogger struct {
	*zap.Logger
}

type lumberJackSink struct {
	*lumberjack.Logger
}

func (lumberJackSink) Sync() error { return nil }

// NewZap creates a new zap logger instance for specified service, eg. gateway, uploader.
//nolint: ireturn, exhaustivestruct // We need to return interface to not to expose zap methods.
// Zap takes care of uninitialized struct fields.
func NewZap(service string) Logger {
	var baseLocation string
	if baseLocation = os.Getenv("GROOT_LOG_LOCATION"); baseLocation == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		baseLocation = path.Join(home, "arcana/logs", service)
	}

	// XXX: Sanitize service string. If the service contains
	// restricted characters the filesystem won't allow to throw error
	logFile := path.Join(baseLocation, service+".log")

	ljLogger := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   true,
	}

	err := zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return lumberJackSink{
			Logger: &ljLogger,
		}, nil
	})
	if err != nil {
		log.Panicf("Unable to register lumberjack sync: %s", err)
	}

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout", fmt.Sprintf("lumberjack:%s", logFile)},
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

// Info publishes info level logs.
func (z *zapLogger) Info(msg string, fields Field) {
	z.Logger.Sugar().Infow(msg, unwrapFields(fields)...)
}

// Info publishes debug logs. Usually used in dev environments.
func (z *zapLogger) Debug(msg string, fields Field) {
	z.Logger.Sugar().Debugw(msg, unwrapFields(fields)...)
}

// Error publishes errors to stderr.
func (z *zapLogger) Error(msg string, fields Field) {
	z.Logger.Sugar().Errorw(msg, unwrapFields(fields)...)
}

// Warn publishes less severe errors that usually doesn't need alerts.
func (z *zapLogger) Warn(msg string, fields Field) {
	z.Logger.Sugar().Warnw(msg, unwrapFields(fields)...)
}

// Fatal publishes the log and stops the execution of program.
func (z *zapLogger) Fatal(msg string, fields Field) {
	z.Logger.Sugar().Fatalw(msg, unwrapFields(fields)...)
}

// unwrapFields is a helper function.
func unwrapFields(fields Field) (unwrapped []interface{}) {
	for k, v := range fields {
		unwrapped = append(unwrapped, k, v)
	}

	return unwrapped
}
