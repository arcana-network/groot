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
	log := logger.NewZap("example")

	log.Info("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Info("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Warn("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Warn("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Debug("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Debug("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Error("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Error("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
}

func logTest() {
	log := logger.NewTestLogger()

	log.Info("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Info("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Warn("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Warn("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Debug("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Debug("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Error("Message1", logger.Field{
		"k1.1": "v1.1",
		"k1.2": "v1.2",
	})
	log.Error("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
	log.Fatal("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})

	log.Info(log.GetLogs(), map[string]interface{}{})
}

func logFatal() {
	// Will panic here
	log := logger.NewZap("exampleFatal")
	log.Fatal("Message2", logger.Field{
		"k2.1": "v2.1",
		"k2.2": "v2.2",
	})
}
