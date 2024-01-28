package logger

import (
	"log/slog"
	"os"
)

func SetDefaultLogger() {
	var logLevel slog.Level
	env := os.Getenv("LOG_LEVEL")
	switch env {
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "INFO":
		logLevel = slog.LevelInfo
	case "WARN":
		logLevel = slog.LevelWarn
	case "ERROR":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	slog.SetDefault(NewCloudLoggingLogger(os.Stdout, logLevel))
}
