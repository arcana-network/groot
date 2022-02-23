package main

import (
	"github.com/arcana-network/gologger/logger"
)

func main() {
	logAllLevels()
	logTest()
	logFatal()
}

func logAllLevels() {
	l := logger.NewZap("log all levels")

	l.Info("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Info("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Warn("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Warn("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Debug("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Debug("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Error("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Error("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
}

func logTest() {
	l := logger.NewTestLogger()

	l.Info("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Info("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Warn("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Warn("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Debug("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Debug("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Error("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	l.Error("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	l.Fatal("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})

	l.Info(l.GetLogs(), map[string]interface{}{})
}

func logFatal() {
	// Will panic here
	l := logger.NewZap("log fatal levels")
	l.Fatal("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
}
