package lib

import (
	"errors"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleQuery(ctx echo.Context) (map[string]interface{}, error){
	from, err := TimeParam(ctx, "from")
	to, err := TimeParam(ctx, "to")
	limit, err := LimitParam(ctx)
	status := parseStatus(ctx)
	m := make(map[string]interface{} )
	m["from"] = from
	m["to"] = to
	m["limit"] = limit
	m["status"] = status
	if err != nil {
		return nil, err
	}
	return m, nil
}

func TimeParam(ctx echo.Context, name string) (time.Time, error) {
	// Now is default time
	t := time.Now()
	value := ctx.QueryParam(name)
	if value == "" {
		return t, nil
	}

	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return t, err
	}

	return parsed, nil
}

const defaultLimit = 10
func LimitParam(ctx echo.Context) (interface{}, error) {
	param := ctx.QueryParam("limit")
	typeCheck := IsNumericOnly(param)
	if typeCheck == false {
		return nil, errors.New("Limit doesnt have integer value")
	}
	limit, err := strconv.Atoi(param)
	if err != nil {
		return nil, err
	}
	return limit, nil
}

func parseStatus(ctx echo.Context) bool {
	param := ctx.QueryParam("status")

	status := true
	switch param {
	  case "false": 
	  status = false;
		break;
	  case "true": 
	  status = true;
		break;
	  default:
		status = true;
	}
	return status
}

func IsNumericOnly(str string) bool {

    if len(str) == 0 {
        return false
    }

    for _, s := range str {
        if s < '0' || s > '9' {
            return false
        }
    }
    return true
}