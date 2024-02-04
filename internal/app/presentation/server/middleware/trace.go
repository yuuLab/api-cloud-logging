package middleware

import (
	"regexp"

	"github.com/labstack/echo"
	"github.com/yuuLab/api-cloud-logging/internal/app/logger"
)

const traceHeader = "X-Cloud-Trace-Context"

var gcpTraceRegexp = regexp.MustCompile(`^([a-f0-9]{32})\/(\d+)(?:;o=(\d+))?$`)

func Trace() echo.MiddlewareFunc {
	return trace
}

func trace(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		trace := c.Request().Header.Get(traceHeader)
		if trace == "" {
			return next(c)
		}

		traceID, spanID := extract(trace)
		if traceID != "" {
			request := c.Request().WithContext(logger.SetTraceID(c.Request().Context(), traceID))
			c.SetRequest(request)
		}
		if spanID != "" {
			request := c.Request().WithContext(logger.SetSpanID(c.Request().Context(), spanID))
			c.SetRequest(request)
		}

		return next(c)
	}
}

func extract(trace string) (traceID, spanID string) {
	matches := gcpTraceRegexp.FindStringSubmatch(trace)
	if len(matches) < 4 {
		return
	}

	return matches[1], matches[2]
}
