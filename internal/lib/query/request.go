package lib

import (
	"errors"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleQuery(ctx echo.Context) (map[string]interface{}, error){
	from, err := timeParam(ctx, "from")
	to, err := timeParam(ctx, "to")
	report := isReportable(ctx)
	limit, err := LimitParam(ctx)
	status := parseStatus(ctx)
	m := make(map[string]interface{} )
	m["from"] = from
	m["to"] = to
	m["limit"] = limit
	m["status"] = status
	m["reportable"] = report
	if err != nil {
		return nil, err
	}
	return m, nil
}

func timeParam(ctx echo.Context, name string) (time.Time, error) {
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
	typeCheck := isNumericOnly(param)
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

func isNumericOnly(str string) bool {

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

func isReportable(ctx echo.Context) bool {
	reportable := ctx.QueryParam("reportable")
	if reportable == "1" { return true }
	return false
}