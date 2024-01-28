package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/logger"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/handler/response"
)

// GetUser get users.
func GetUsers(c echo.Context) error {
	// NOTE: This is the sample user.
	u := &response.User{
		UserID: "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
		Name:   "Jon",
		Age:    20,
	}

	ctx := c.Request().Context()

	slog.DebugContext(ctx, "DEBUG Message", "Name", u.Name, "Age", u.Age)
	slog.InfoContext(ctx, "INFO Message", "Name", u.Name, "Age", u.Age)
	slog.WarnContext(ctx, "WARN Message", "Name", u.Name, "Age", u.Age)
	slog.ErrorContext(ctx, "ERROR Message", "Name", u.Name, "Age", u.Age)

	slog.Default().Log(ctx, logger.LevelNotice, "NOTICE Message", "Name", u.Name, "Age", u.Age)
	slog.Default().Log(ctx, logger.LevelAlert, "ALERT Message", "Name", u.Name, "Age", u.Age)
	slog.Default().Log(ctx, logger.LevelCritical, "CRITICAL Message", "Name", u.Name, "Age", u.Age)
	slog.Default().Log(ctx, logger.LevelEmergency, "EMERGENCY Message", "Name", u.Name, "Age", u.Age)

	return c.JSON(http.StatusOK, u)
}
