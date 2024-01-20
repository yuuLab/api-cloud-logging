package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/handler/response"
)

func CheckAlive(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.CheckAlive{Status: "OK"})
}
