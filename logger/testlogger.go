package logger

import (
	"bytes"
	"fmt"
)

// TestLogger holds the logs in-memory for testing purposes.
type TestLogger struct {
	b *bytes.Buffer
}

// NewTestLogger returns a logger that can be used for unit tests.
func NewTestLogger() *TestLogger {
	return &TestLogger{
		b: &bytes.Buffer{},
	}
}

// Info publishes info level logs.
func (t *TestLogger) Info(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

// Info publishes debug logs. Usually used in dev environments.
func (t *TestLogger) Debug(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

// Error publishes errors to stderr.
func (t *TestLogger) Error(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

// Warn publishes less severe errors that usually doesn't need alerts.
func (t *TestLogger) Warn(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

// Fatal publishes the log and stops the execution of program.
func (t *TestLogger) Fatal(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

// Panic publishes the log and stops the execution of goroutine.
func (t *TestLogger) Panic(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

// GetLogs is a helper function to get in-memory logs.
func (t *TestLogger) GetLogs() string {
	return t.b.String()
}
