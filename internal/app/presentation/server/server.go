package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	slog.Info("server shutdown completed.")
}
