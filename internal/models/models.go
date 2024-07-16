package models

import (
	"database/sql"
	"sync"
)

type Account struct {
	ID      string
	Balance float64
	MU      sync.Mutex `json:"MU,omitempty"`
	DB      *sql.DB    `json:"DB,omitempty"`
}

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() (float64, error)
}
