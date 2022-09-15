package main

import (
	"net/http"
	"roof-stack/internal/config"
	"roof-stack/internal/container"
	"roof-stack/internal/database"
	"roof-stack/internal/route"

	"github.com/labstack/echo/v4"
)

const (
	prefix     = "api"
	version    = "v1"
)

func main() {
	defer database.CloseDatabaseConnection(container.DB)

	e := echo.New()
	config.CORSConfig(e)
	e.GET("/ping",func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	prefix := joinStrings(prefix, "/", version, "/")

	route.AuthRoutes(prefix, e, container.IauthHandler)
	route.WalletRoutes(prefix, e, container.IwalletHandler)
	route.TransactionRoutes(prefix, e, container.ItransactionHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func joinStrings(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}