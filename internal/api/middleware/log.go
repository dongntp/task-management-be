package middleware

import (
	"net/http"
	"task-management-be/internal/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logConfig = middleware.RequestLoggerConfig{
		LogRequestID:  true,
		LogMethod:     true,
		LogURI:        true,
		LogHost:       true,
		LogStatus:     true,
		LogLatency:    true,
		LogUserAgent:  true,
		LogRemoteIP:   true,
		LogError:      true,
		LogValuesFunc: logFunc,
	}
)

func logFunc(_ echo.Context, v middleware.RequestLoggerValues) error {
	var level zapcore.Level

	fields := []zap.Field{
		zap.String("method", v.Method),
		zap.String("host", v.Host),
		zap.String("uri", v.URI),
		zap.Int("status", v.Status),
		zap.Time("startTime", v.StartTime),
		zap.Duration("latency", v.Latency),
		zap.String("remoteIP", v.RemoteIP),
		zap.String("agent", v.UserAgent),
	}

	if v.Status == http.StatusOK {
		level = zap.InfoLevel
	} else {
		level = zap.ErrorLevel
		fields = append(fields, zap.Error(v.Error))
	}

	logger.Logger.Named(v.RequestID).Log(level, "request", fields...)
	return nil
}

func LogWithRequestID(parentLogger *zap.Logger) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.RequestID(),
		injectLogger(parentLogger),
		middleware.RequestLoggerWithConfig(logConfig),
	}
}

func injectLogger(parentLogger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
			childLogger := parentLogger.Named(requestID)
			newGoCtx := logger.Inject(ctx.Request().Context(), childLogger)
			ctx.SetRequest(ctx.Request().WithContext(newGoCtx))
			return next(ctx)
		}
	}
}
