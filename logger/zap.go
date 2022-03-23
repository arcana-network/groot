package logger

import (
	"errors"
	"fmt"
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

//nolint // Global to avoid multiple err allocations, later move to errors.go
var errEmptyService = errors.New("service cannot be empty")

// zapLogger wraps the zap logging library to satisfy Logger interface.
type zapLogger struct {
	*zap.Logger
}

type lumberJackSink struct {
	*lumberjack.Logger
}

func (lumberJackSink) Sync() error { return nil }

// newZap creates a new zap logger instance for specified service, eg. gateway, uploader.
//nolint:ireturn, exhaustivestruct // Return interface to protect zap methods.
func newZap(service string) (Logger, error) {
	if service == "" {
		return nil, errEmptyService
	}

	var baseLocation string
	if baseLocation = os.Getenv("GROOT_LOG_LOCATION"); baseLocation == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return &zapLogger{}, fmt.Errorf("get home dir: %w", err)
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
		return nil, fmt.Errorf("register lumberjack sync: %w", err)
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
		return nil, fmt.Errorf("onfig build: %w", err)
	}

	return &zapLogger{
		logger.Named(service),
	}, nil
}

// Info publishes info level logs.
func (z *zapLogger) Info(msg string, fields Field) {
	z.Logger.Sugar().Infow(msg, unwrapFields(fields)...)
}

// Debug publishes debug logs. Noisy in prod, usually used in dev environments.
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
