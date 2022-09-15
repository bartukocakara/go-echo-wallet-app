package lib

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"`
}

func Error(err error, ctx echo.Context){
	errObj := HTTPError{
		Code: http.StatusInternalServerError,
		Message: err.Error(),
	}
	switch err {
	case ErrBadRequest:
		errObj.Code = http.StatusBadRequest
	case ErrNotFound:
		errObj.Code = http.StatusNotFound
	case ErrDuplicateEntry, ErrConflict:
		errObj.Code = http.StatusConflict
	case ErrForbidden:
		errObj.Code = http.StatusForbidden
	case ErrUnprocessableEntity:
		errObj.Code = http.StatusUnprocessableEntity
	case ErrPartialOk:
		errObj.Code = http.StatusPartialContent
	case ErrGone:
		errObj.Code = http.StatusGone
	case ErrUnauthorized:
		errObj.Code = http.StatusUnauthorized
	}
	he, ok := err.(*echo.HTTPError)
	if ok {
		errObj.Code = he.Code
		errObj.Message = fmt.Sprintf("%v", he.Message)
	}
	errObj.Name = http.StatusText(errObj.Code)
	if !ctx.Response().Committed {
		if ctx.Request().Method == echo.HEAD {
			ctx.NoContent(errObj.Code)
		} else {
			ctx.JSON(errObj.Code, errObj)
		}
	}
}