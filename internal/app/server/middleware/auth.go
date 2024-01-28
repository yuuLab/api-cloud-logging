package middleware

import (
	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/logger"
)

func Authenticate() echo.MiddlewareFunc {
	return authenticate
}

func authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//NOTE: ユーザー認証を行い、ユーザーIDもしくはユーザーオブジェクトを本来は取得する
		request := c.Request().WithContext(logger.SetUserID(c.Request().Context(), "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"))
		c.SetRequest(request)
		return next(c)
	}
}
