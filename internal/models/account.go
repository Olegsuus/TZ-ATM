package models

import (
	"database/sql"
	"fmt"
)

func (a *Account) Deposit(amount float64) error {
	a.MU.Lock()
	defer a.MU.Unlock()
	_, err := a.DB.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, a.ID)
	return err
}

func (a *Account) Withdraw(amount float64) error {
	a.MU.Lock()
	defer a.MU.Unlock()

	var currentBalance float64
	err := a.DB.QueryRow("SELECT balance FROM accounts WHERE id = ?", a.ID).Scan(&currentBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("account with ID %s does not exist", a.ID)
		}
		return err
	}

	if currentBalance < amount {
		return fmt.Errorf("insufficient funds: current balance is %.2f, requested withdrawal amount is %.2f", currentBalance, amount)
	}

	_, err = a.DB.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, a.ID)
	if err != nil {
		return fmt.Errorf("failed to withdraw from account: %v", err)
	}

	return nil
}

func (a *Account) GetBalance() (float64, error) {
	var currentBalance float64
	err := a.DB.QueryRow("SELECT balance FROM accounts WHERE id = ?", a.ID).Scan(&currentBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("account with ID %s does not exist", a.ID)
		}
		return 0, err
	}
	return currentBalance, nil
}
