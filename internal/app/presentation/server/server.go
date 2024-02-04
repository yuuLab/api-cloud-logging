package server

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"
	"github.com/yuuLab/api-cloud-logging/internal/app/logger"
	"github.com/yuuLab/api-cloud-logging/internal/app/presentation/server/middleware"
)

// Run runs the server.
func Run() {
	logger.SetDefaultLogger()

	e := echo.New()

	e.Debug = true
	e.Use(echomiddleware.Recover(), middleware.Authenticate(), middleware.Trace())
	Router(e)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		slog.Warn(fmt.Sprintf("Defaulting to port %s", port))
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
