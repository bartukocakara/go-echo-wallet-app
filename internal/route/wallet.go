package route

import (
	handler "roof-stack/internal/handler/wallet"
	"roof-stack/internal/middleware"

	"github.com/labstack/echo/v4"
)


func WalletRoutes(prefix string,
				   e *echo.Echo,
				   handler handler.IWalletHandler){

	walletRoutes := e.Group(prefix + "wallets")

	walletRoutes.POST("", handler.Create, middleware.IsLoggedIn)
	walletRoutes.GET("/:wallerID", handler.Get)
	walletRoutes.GET("", handler.List)
	
}