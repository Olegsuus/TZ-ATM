package handlers

import (
	"TZ-ATM/internal/models"
	"database/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func createAccount(db *sql.DB) (*models.Account, error) {
	id := uuid.New().String()
	_, err := db.Exec("INSERT INTO accounts (id, balance) VALUES (?, ?)", id, 0)
	if err != nil {
		return nil, err
	}
	return &models.Account{ID: id, DB: db}, nil
}

func CreateAccountHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		account, err := createAccount(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, account)
	}
}

func DepositHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID := c.Param("id")
		amount, err := strconv.ParseFloat(c.QueryParam("amount"), 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
		}
		account := &models.Account{ID: accountID, DB: db}

		resultChan := make(chan error)
		go func(bankAccount models.BankAccount) {
			defer close(resultChan)
			err := bankAccount.Deposit(amount)
			if err != nil {
				resultChan <- err
			} else {
				resultChan <- nil
			}
		}(account)

		if err := <-resultChan; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{"status": "deposit completed"})
	}
}

func WithdrawHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID := c.Param("id")
		amountStr := c.QueryParam("amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
		}

		account := &models.Account{ID: accountID, DB: db}

		resultChan := make(chan error)
		go func(bankAccount models.BankAccount) {
			defer close(resultChan)
			err := bankAccount.Withdraw(amount)
			if err != nil {
				resultChan <- err
			} else {
				resultChan <- nil
			}
		}(account)

		if err := <-resultChan; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{"status": "withdrawal completed"})
	}
}

func BalanceHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID := c.Param("id")
		account := &models.Account{ID: accountID, DB: db}

		resultChan := make(chan *models.BalanceResult)
		go func(bankAccount models.BankAccount) {
			defer close(resultChan)

			balance, err := bankAccount.GetBalance()
			resultChan <- &models.BalanceResult{Balance: balance, Err: err}
		}(account)

		result := <-resultChan
		if result.Err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Err.Error())
		}

		return c.JSON(http.StatusOK, map[string]float64{"balance": result.Balance})
	}
}
