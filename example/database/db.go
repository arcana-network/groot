package database

import (
	"errors"

	"github.com/arcana-network/groot/logger"
)

type DB struct {
	connURL string
	log     logger.Logger
}

var errTableNotFound = errors.New("table not found")

func NewDB(connURL string, log logger.Logger) DB {
	return DB{
		connURL: connURL,
		log:     log,
	}
}

func (db *DB) Connect() {
	// Connection logic
	db.log.Debug("Connecting to database", logger.Field{
		"DBuri": db.connURL,
	})

	return
}

func (db *DB) Query(q string) (string, error) {
	// Query execution logic
	db.log.Debug("Executing SQL query", logger.Field{
		"query": q,
	})
	// lets' always return error for demo purpose
	return "", errTableNotFound
}

func (db *DB) Close() {
	// Closure logic
	db.log.Debug("DB connection closed", logger.Field{
		"DBuri": db.connURL,
	})

	return
}
