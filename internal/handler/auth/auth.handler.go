package handler

import (
	"fmt"
	"net/http"
	"roof-stack/internal/dto"
	"roof-stack/internal/entity"
	"roof-stack/internal/helper"
	lib "roof-stack/internal/lib/error"
	"roof-stack/internal/service"

	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type authHandler struct {
	authService service.IAuthService
	jwtService service.IJWTService
}

func NewAuthHandler(authService service.IAuthService, jwtService service.IJWTService) *authHandler {
	return &authHandler{
		authService: authService,
		jwtService: jwtService,
	}
}

func(h *authHandler) Register(ctx echo.Context) error {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.Bind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(lib.ErrBadRequest, errDTO.Error(), helper.EmptyObj{})
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	if !h.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse(lib.ErrConflict, lib.EmailDuplicate, helper.EmptyObj{})
		return ctx.JSON(http.StatusConflict, response)
	} 
	createdUser := h.authService.CreateUser(registerDTO)
	token := h.jwtService.GenerateToken(createdUser)
	createdUser.Token = token
	response := helper.BuildResponse(true, lib.StatusCreated, createdUser)
	return ctx.JSON(http.StatusCreated, response)
}

func (h *authHandler) Login(ctx echo.Context) error  {
	var loginDTO dto.LoginDTO
	errDTO := ctx.Bind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(lib.ErrBadRequest, errDTO.Error(), helper.EmptyObj{})
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	authResult := h.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	// return ctx.JSON(http.StatusUnauthorized, authResult)
	if verifiedUser, ok := authResult.(entity.User); ok {
		generatedToken := h.jwtService.GenerateToken(verifiedUser )
		fmt.Println(verifiedUser)
		fmt.Println(generatedToken)

		verifiedUser.Token = generatedToken
		response := helper.BuildResponse(true, lib.StatusOK, verifiedUser)
		return ctx.JSON(http.StatusOK, response)
	}
	response := helper.BuildErrorResponse(lib.ErrUnauthorized, "Invalid Credential", helper.EmptyObj{})
	return ctx.JSON(http.StatusUnauthorized, response)
}