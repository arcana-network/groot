package models

import (
	"github.com/arcana-network/groot/example/database"
	"github.com/arcana-network/groot/logger"
)

type DBTransaction struct {
	log logger.Logger
	db  database.DB
}

func NewDBTransaction(db database.DB, log logger.Logger) TransactionModel {
	return DBTransaction{
		log: log,
		db:  db,
	}
}

func (t *DBTransaction) getTransactionAmount(userName string) int {
	// Logic to retrieve txns from DB and assemble.
	t.log.Info("Getting transaction details", logger.Field{
		"username": userName,
	})

	return 20
}
