package route

import (
	handler "roof-stack/internal/handler/auth"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(prefix string, e *echo.Echo, handler handler.IAuthHandler){
	authRoutes := e.Group(prefix + "auth")
	authRoutes.POST("/register", handler.Register)
	authRoutes.POST("/login", handler.Login)
	
}