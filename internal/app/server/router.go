package server

import (
	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/handler"
)

// Router ...
func Router(e *echo.Echo) {
	e.GET("/api/v1/users", handler.GetUsers)
}
