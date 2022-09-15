package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"roof-stack/internal/database"
	"roof-stack/internal/helper"
	lib "roof-stack/internal/lib/error"
	"roof-stack/internal/repository"
	"roof-stack/internal/service"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = database.SetupDatabaseConnection()
	IuserRepository repository.IUserRepository = repository.NewUserRepository(DB)
	IjwtService     service.IJWTService        = service.NewJWTService()
	IauthService    service.IAuthService       = service.NewAuthService(IuserRepository)
)

var mockDB = `{
    "email": "bartukocakaraa@gmail.com",
    "password": "123456"
}`

var testRegisterResponseJSON = `{"status":true,"message":"Created","errors":null,"data":{"id":3,"first_name": "","last_name": "","email": "bartukocakaraa@gmail.com","token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlIjoidXNlciIsImZpcnN0X25hbWUiOiIiLCJleHAiOjE2NjM1MTIwNTR9.bKHGR1CMWZndmkIluy8GHhGrmrT_PrPCE5Z3RuKHBxo","created_at": "2022-09-15T17:40:54.755696+03:00","updated_at": "2022-09-15T17:40:54.755696+03:00"}}`

func TestRegister(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(mockDB))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	IauthHandler := NewAuthHandler(IauthService, IjwtService)
	fmt.Println(rec.Body.String())
	// Assertions
	response := helper.BuildResponse(true, lib.StatusCreated, testRegisterResponseJSON)
	response.Data = nil
	fmt.Println(response)
	if assert.NoError(t, IauthHandler.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, response, rec.Body.String())
	}
}