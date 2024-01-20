package app

import (
	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/handler"
)

// Router ...
func Router(e *echo.Echo) {
	e.GET("/checkalive", handler.CheckAlive)
	e.GET("/users", handler.GetUsers)
}
