package models

type UserModel interface {
	getUserID(name string) (string, error)
}

type User struct {
	name string
}

type TransactionModel interface{}

type Transaction struct {
	amount int
	user   User
}
