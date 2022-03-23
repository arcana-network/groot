package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arcana-network/groot/example/controller"
	"github.com/arcana-network/groot/example/database"
	"github.com/arcana-network/groot/example/models"
	"github.com/arcana-network/groot/logger"
)

func main() {
	log := logger.NewZapLogger("example")
	db := database.NewDB("postgres://user?pass:db", log)
	userModel := models.NewDBUser(db, log)
	userController := controller.NewUserController(userModel, log)
	http.HandleFunc("/user/id", userController.GetUserID)

	url := os.Getenv("SERVER_URL")
	port := os.Getenv("PORT")

	if url == "" && port == "" {
		url = "0.0.0.0"
		port = "9000"

		log.Warn("No url port provided, using default port", logger.Field{
			"url":  url,
			"port": port,
		})
	}

	log.Info("starting http server", logger.Field{
		"url":  url,
		"port": port,
	})

	addr := fmt.Sprintf("%s:%s", url, port)
	err := http.ListenAndServe(addr, nil).Error()

	log.Fatal("Unable to start server: ", logger.Field{
		"error": err,
	})
}
