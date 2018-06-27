package context

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

const (
	CodeSuccess   Code = iota
	CodeUndefined      = 999
)

type Code int

type SError struct {
	Code Code
	Msg  string
}

func (err *SError) Error() string {
	return err.Msg
}

func ErrBadRequest(detail string) *SError {
	return &SError{
		Code: CodeUndefined,
		Msg:  fmt.Sprintf("请求失败:%s", detail),
	}
}

func Say(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   CodeSuccess,
		"result": result,
	})
}

var HTTPErrorHandler = func(err error, c echo.Context) {
	if de, ok := err.(*SError); ok && de != nil {
		if !c.Response().Committed() {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":  de.Code,
				"msg": de.Msg,
			})
		}
		return
	}
	if he, ok := err.(*echo.HTTPError); ok && he != nil {
		if !c.Response().Committed() {
			c.String(he.Code, he.Message)
		}
		return
	}
	if !c.Response().Committed() {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
