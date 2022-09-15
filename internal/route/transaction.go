package route

import (
	handler "roof-stack/internal/handler/transaction"
	"roof-stack/internal/middleware"

	"github.com/labstack/echo/v4"
)


func TransactionRoutes(prefix string,
				   e *echo.Echo,
				   handler handler.ITransactionHandler){

	transactionRoutes := e.Group(prefix + "transactions")

	transactionRoutes.POST("", handler.Create, middleware.IsLoggedIn)
	transactionRoutes.POST("/wallets/:walletId", handler.Create, middleware.IsLoggedIn)
	transactionRoutes.GET("/:transactionId", handler.Get)
	transactionRoutes.GET("", handler.List)
	transactionRoutes.POST("/report", handler.Report)
	
}