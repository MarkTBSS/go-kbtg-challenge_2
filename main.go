package main

import (
	"github.com/MarkTBSS/go-kbtg-challenge_2/postgres"
	"github.com/MarkTBSS/go-kbtg-challenge_2/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/MarkTBSS/go-kbtg-challenge_2/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	databaseInstance, err := postgres.New()
	if err != nil {
		panic(err)
	}
	echoInstance := echo.New()
	echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(databaseInstance)
	echoInstance.GET("/api/v1/wallets", handler.WalletHandler)
	echoInstance.Logger.Fatal(echoInstance.Start(":1323"))
}
