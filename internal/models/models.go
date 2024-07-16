package models

import (
	"database/sql"
	"sync"
)

type Account struct {
	ID      string
	Balance float64
	MU      sync.Mutex
	DB      *sql.DB
}
