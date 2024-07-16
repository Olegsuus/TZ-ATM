package main

import (
	"TZ-ATM/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() (float64, error)
}

//func main() {
//	db := initDB() // Предполагаем, что эта функция настроена соответствующим образом
//
//	e := echo.New()
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	// Setup routes
//	e.POST("/accounts", func(c echo.Context) error {
//		return createAccountHandler(c, db)
//	})
//
//	e.POST("/accounts/:id/deposit", depositHandler(db))
//	e.POST("/accounts/:id/withdraw", withdrawHandler(db))
//	e.GET("/accounts/:id/balance", balanceHandler(db))
//
//	// Start server
//	e.Logger.Fatal(e.Start(":8080"))
//}

func main() {
	db := db.InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup routes

	e.POST("/accounts/:id/deposit", func(c echo.Context) error {
		accountID := c.Param("id")
		amount, err := strconv.ParseFloat(c.QueryParam("amount"), 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
		}

		resultChan := make(chan error)
		go func() {
			defer close(resultChan)
			account := &Account{ID: accountID, db: db}
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
	})

	e.POST("/accounts/:id/withdraw", func(c echo.Context) error {
		accountID := c.Param("id")
		amount, err := strconv.ParseFloat(c.QueryParam("amount"), 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
		}

		resultChan := make(chan error)
		go func() {
			defer close(resultChan)
			account := &Account{ID: accountID, db: db}
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
	})

	e.GET("/accounts/:id/balance", func(c echo.Context) error {
		accountID := c.Param("id")

		resultChan := make(chan *BalanceResult)
		go func() {
			defer close(resultChan)
			account := &Account{ID: accountID, db: db}
			balance, err := account.GetBalance()
			resultChan <- &BalanceResult{Balance: balance, Err: err}
		}()

		result := <-resultChan
		if result.Err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Err.Error())
		}

		return c.JSON(http.StatusOK, map[string]float64{"balance": result.Balance})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
