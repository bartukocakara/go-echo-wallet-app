package handler

import (
	"net/http"
	"roof-stack/internal/dto"
	"roof-stack/internal/helper"
	lib "roof-stack/internal/lib/error"
	"roof-stack/internal/service"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type IWalletHandler interface {
	Create(ctx echo.Context) error
	List(ctx echo.Context) error	
}

type walletHandler struct {
	walletService service.IWalletService 
}

func NewWalletHandler(walletService service.IWalletService)*walletHandler{
	return &walletHandler{
		walletService: walletService,
	}
}

func(h *walletHandler) Create(ctx echo.Context) error {
	var createDTO dto.CreateWalletDTO
	errDTO := ctx.Bind(&createDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(lib.ErrBadRequest, errDTO.Error(), helper.EmptyObj{})
		return echo.NewHTTPError(http.StatusBadRequest, response)
	} else {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(*service.JwtCustomClaim)
		userID := claims.UserID
		createDTO.UserID = userID

		result, err := h.walletService.Create(createDTO)
		if err != nil {
			res := helper.BuildErrorResponse(lib.ErrBadRequest, "No data", result)
			return ctx.JSON(http.StatusOK, res)
		}
		response := helper.BuildResponse(true, lib.StatusOK, result)
		return ctx.JSON(http.StatusCreated, response)
	}
}

func (h *walletHandler) List(ctx echo.Context) error {
	wallets, err := h.walletService.List()
	if err != nil {
		res := helper.BuildErrorResponse(lib.ErrBadRequest, "No data", wallets)
		return ctx.JSON(http.StatusOK, res)
	}
	res := helper.BuildResponse(true, "OK", wallets)
	return ctx.JSON(http.StatusOK, res)
}

func (h *walletHandler) Get(ctx echo.Context) error {
	return nil
}
