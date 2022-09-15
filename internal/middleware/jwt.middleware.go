package middleware

import (
	"roof-stack/internal/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
const SignInKey = "greatest-secret-ever"
var next echo.HandlerFunc
var SigningKey = []byte("greatest-secret-ever")

//AuthorizeJWT validates the token user given, return 401 if not valid
var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
    Claims:     &service.JwtCustomClaim{},
    SigningKey: SigningKey,
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        role := claims["role"].(string)
        if role != "admin" {
            return echo.ErrUnauthorized
        }
        return next(c)
    }
}