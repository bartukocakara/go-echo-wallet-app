package handler

import (
	"fmt"
	"net/http"
	"roof-stack/internal/dto"
	"roof-stack/internal/helper"
	lib "roof-stack/internal/lib/error"
	query "roof-stack/internal/lib/query"
	"roof-stack/internal/service"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ITransactionHandler interface {
	Create(ctx echo.Context) error
	List(ctx echo.Context) error	
	Get(ctx echo.Context) error
	Report(ctx echo.Context) error
}

type transactionHandler struct {
	transactionService service.ITransactionService 
}

func NewTransactionHandler(transactionService service.ITransactionService)*transactionHandler{
	return &transactionHandler{
		transactionService: transactionService,
	}
}

func(h *transactionHandler) Create(ctx echo.Context) error {
	var createDTO dto.CreateTransactionDTO
	errDTO := ctx.Bind(&createDTO)
	id, err := strconv.ParseUint(ctx.Param("walletId"), 0, 0)

	if err != nil {
		response := helper.BuildErrorResponse(lib.ErrBadRequest, errDTO.Error(), helper.EmptyObj{})
		return echo.NewHTTPError(http.StatusBadRequest, response)
	} else {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(*service.JwtCustomClaim)
		userID := claims.UserID
		createDTO.UserID = userID
		createDTO.WalletID = id

		result, err := h.transactionService.Create(createDTO)
		if err != nil {
			fmt.Println(err)
			response := helper.BuildErrorResponse(true, err.Error(), err)
			return ctx.JSON(http.StatusBadRequest, response)
		}

		response := helper.BuildResponse(true, lib.StatusOK, result)
		return ctx.JSON(http.StatusCreated, response)
	}
}

func (h *transactionHandler) List(ctx echo.Context) error {
	requestParams, err := query.HandleQuery(ctx)
	if err != nil {
		res := helper.BuildErrorResponse(lib.ErrBadRequest, err.Error(), err.Error())
		return ctx.JSON(http.StatusBadRequest, res)
	}
	transactions, err  := h.transactionService.List(requestParams)
	res := helper.BuildResponse(true, "OK", transactions)
	return ctx.JSON(http.StatusOK, res)
}

func (h *transactionHandler) Get(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("transactionId"), 0, 0)
	transaction, err  := h.transactionService.Get(id)
	if err != nil {
		res := helper.BuildErrorResponse(lib.ErrBadRequest, "No data", transaction)
		return ctx.JSON(http.StatusOK, res)
	}
	res := helper.BuildResponse(true, "OK", transaction)
	return ctx.JSON(http.StatusOK, res)
}

func (h *transactionHandler) Report(ctx echo.Context) error {
	requestParams, err := query.HandleQuery(ctx)
	if err != nil {
		res := helper.BuildErrorResponse(lib.ErrBadRequest, err.Error(), err.Error())
		return ctx.JSON(http.StatusBadRequest, res)
	}
	transactions, err  := h.transactionService.List(requestParams)
	res := helper.BuildResponse(true, "OK", transactions)
	return ctx.JSON(http.StatusOK, res)
}