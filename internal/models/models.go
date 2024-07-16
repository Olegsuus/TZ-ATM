package models

import (
	"database/sql"
	"github.com/google/uuid"
	"sync"
)

type Account struct {
	ID      string
	Balance float64
	mu      sync.Mutex
	db      *sql.DB
}

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() (float64, error)
}

func (a *Account) InitAccount(id string, db *sql.DB) {
	a.ID = id
	a.db = db
}

func CreateAccount(db *sql.DB) (*Account, error) {
	id := uuid.New().String()
	_, err := db.Exec("INSERT INTO accounts (id, balance) VALUES (?, ?)", id, 0)
	if err != nil {
		return nil, err
	}
	return &Account{ID: id, db: db}, nil
}
