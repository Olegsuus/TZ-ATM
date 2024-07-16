package main

import (
	"TZ-ATM/internal/db"
	"TZ-ATM/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() (float64, error)
}

func main() {
	storage := db.InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/accounts", handlers.CreateAccountHandler(storage))

	e.POST("/accounts/:id/deposit", handlers.DepositHandler(storage))
	e.POST("/accounts/:id/withdraw", handlers.WithdrawHandler(storage))
	e.GET("/accounts/:id/balance", handlers.BalanceHandler(storage))

	e.Logger.Fatal(e.Start(":8080"))
}
