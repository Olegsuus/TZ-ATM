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

func createAccountHandler(c echo.Context, db *sql.DB) error {
	account, err := createAccount(db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, account)
}

func depositHandler(db *sql.DB, c echo.Context) error {
	accountID := c.Param("id")
	amount, err := strconv.ParseFloat(c.QueryParam("amount"), 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
	}

	resultChan := make(chan error)
	go func() {
		defer close(resultChan)
		account := &models.Account{ID: accountID, DB: db}
		err := account.Deposit(amount)
		if err != nil {
			resultChan <- err
		} else {
			resultChan <- nil
		}
	}()

	if err := <-resultChan; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "deposit completed"})
}

func withdrawHandler(db *sql.DB, c echo.Context) error {
	accountID := c.Param("id")
	amount, err := strconv.ParseFloat(c.QueryParam("amount"), 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
	}

	resultChan := make(chan error)
	go func() {
		defer close(resultChan)
		account := &models.Account{ID: accountID, DB: db}
		err := account.Withdraw(amount)
		if err != nil {
			resultChan <- err
		} else {
			resultChan <- nil
		}
	}()

	if err := <-resultChan; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "withdrawal completed"})
}

func balanceHandler(db *sql.DB, c echo.Context) error {

}
