package models

import (
	"fmt"

	"github.com/arcana-network/groot/example/database"
	"github.com/arcana-network/groot/logger"
)

type DBUser struct {
	log logger.Logger
	db  database.DB
}

func NewDBUser(db database.DB, log logger.Logger) UserModel {
	return DBUser{
		log: log,
		db:  db,
	}
}

func (u DBUser) getUserID(name string) (string, error) {
	// Example purpose only, use prepare statements to avoid sql injection.
	result, err := u.db.Query(fmt.Sprintf("SELECT * FROM users where name=%s", name))
	if err != nil {
		u.log.Error("Unable to get user id", logger.Field{
			"username": name,
		})

		return "", err
	}

	return result, nil
}
