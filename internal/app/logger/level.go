package logger

import "log/slog"

const (
	// NOTE: slog default log level
	// LevelDebug Level = -4
	// LevelInfo  Level = 0
	// LevelWarn  Level = 4
	// LevelError Level = 8
	LevelNotice    = slog.Level(2)
	LevelCritical  = slog.Level(12)
	LevelAlert     = slog.Level(16)
	LevelEmergency = slog.Level(20)
)
