package logger

import (
	"bytes"
	"fmt"
)

type testLogger struct {
	b *bytes.Buffer
}

func NewTestLogger() *testLogger {
	return &testLogger{
		b: &bytes.Buffer{},
	}
}

func (t *testLogger) Info(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}
func (t *testLogger) Debug(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}
func (t *testLogger) Error(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}
func (t *testLogger) Warn(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}
func (t *testLogger) Fatal(msg string, fields Field) {
	fmt.Fprintln(t.b, msg, fields)
}

func (t *testLogger) GetLogs() string {
	return t.b.String()
}
