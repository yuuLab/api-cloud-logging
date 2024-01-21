package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/handler/response"
)

// GetUser get users.
func GetUsers(c echo.Context) error {
	// NOTE: This is the sample user.
	u := &response.User{
		UserID: "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
		Name:   "Jon",
		Email:  "jon@labstack.com",
	}
	slog.Info("Successfully get user info.", "uid", u.UserID, "Name", u.Name)
	return c.JSON(http.StatusOK, u)
}
