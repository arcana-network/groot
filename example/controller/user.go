package controller

import (
	"net/http"

	"github.com/arcana-network/groot/example/models"
	"github.com/arcana-network/groot/logger"
)

type UserController struct {
	user models.UserModel
	log  logger.Logger
}

func NewUserController(user models.UserModel, log logger.Logger) UserController {
	return UserController{
		user: user,
		log:  log,
	}
}

// GetUserID handles /users/id API.
func (uc *UserController) GetUserID(w http.ResponseWriter, r *http.Request) {
	// logic to get user details from models
}
