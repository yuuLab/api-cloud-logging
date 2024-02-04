package logger

import (
	"context"
	"fmt"
	"io"
	"os"

	"log/slog"
)

type contextKey string

const (
	userIDCtxKey  contextKey = "userIDCtxKey"
	traceIDCtxKey contextKey = "traceIDCtxKey"
	spanIDCtxKey  contextKey = "spanIDCtxKey"
)

// cloudLoggingHandler is the logging handler for the cloud logging.
type cloudLoggingHandler struct {
	handler slog.Handler
}

var cloudLoggingKey = struct {
	severity       string
	message        string
	trace          string
	spanID         string
	labels         string
	sourceLocation string
}{
	severity:       "severity",
	message:        "message",
	trace:          "logging.googleapis.com/trace",
	spanID:         "logging.googleapis.com/spanId",
	labels:         "logging.googleapis.com/labels",
	sourceLocation: "logging.googleapis.com/sourceLocation",
}

func NewCloudLoggingLogger(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(NewCloudLoggingHandler(w, level))
}

func NewCloudLoggingHandler(w io.Writer, level slog.Level) slog.Handler {
	return &cloudLoggingHandler{
		handler: slog.NewJSONHandler(
			w,
			&slog.HandlerOptions{
				AddSource:   true,
				Level:       level,
				ReplaceAttr: replaceAttrFunc(),
			},
		),
	}
}

func replaceAttrFunc() func(groups []string, a slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		switch {
		// msg -> message
		case a.Key == slog.MessageKey:
			return slog.String(cloudLoggingKey.message, a.Value.String())
		// level -> severity, WARN -> WARNING
		case a.Key == slog.LevelKey:
			levle := a.Value.Any().(slog.Level)
			switch levle {
			case slog.LevelWarn:
				return slog.String(cloudLoggingKey.severity, "WARNING")
			case LevelNotice:
				return slog.String(cloudLoggingKey.severity, "NOTICE")
			case LevelCritical:
				return slog.String(cloudLoggingKey.severity, "CRITICAL")
			case LevelAlert:
				return slog.String(cloudLoggingKey.severity, "ALERT")
			case LevelEmergency:
				return slog.String(cloudLoggingKey.severity, "EMERGENCY")
			}
			return slog.String(cloudLoggingKey.severity, a.Value.String())
		// level -> severity
		case a.Key == slog.LevelKey:
			return slog.Attr{
				Key:   cloudLoggingKey.severity,
				Value: a.Value,
			}
		// source -> logging.googleapis.com/sourceLocation
		case a.Key == slog.SourceKey:
			return slog.Attr{
				Key:   cloudLoggingKey.sourceLocation,
				Value: a.Value,
			}
		}
		return a
	}
}

func (h *cloudLoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *cloudLoggingHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(traceIDCtxKey).(string); ok {
		r.AddAttrs(
			slog.String(cloudLoggingKey.trace, fmt.Sprintf("projects/%s/traces/%s", os.Getenv("PROJECT_ID"), traceID)),
		)
	}

	if spanID, ok := ctx.Value(spanIDCtxKey).(string); ok {
		r.AddAttrs(slog.String(cloudLoggingKey.spanID, spanID))
	}

	if uid, ok := ctx.Value(userIDCtxKey).(string); ok {
		r.AddAttrs(slog.Group(
			cloudLoggingKey.labels, slog.String("uid", uid),
		))
	}

	return h.handler.Handle(ctx, r)
}

func (h *cloudLoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.handler.WithAttrs(attrs)
}

func (h *cloudLoggingHandler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}

var _ slog.Handler = (*cloudLoggingHandler)(nil)

func SetUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, userIDCtxKey, uid)
}

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtxKey, traceID)
}

func SetSpanID(ctx context.Context, spanID string) context.Context {
	return context.WithValue(ctx, spanIDCtxKey, spanID)
}
