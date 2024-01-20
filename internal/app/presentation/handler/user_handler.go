package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/handler/response"
)

// GetUser get users.
func GetUsers(c echo.Context) error {
	// NOTE: This is the sample user.
	u := &response.User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}
