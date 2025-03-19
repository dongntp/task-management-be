package middleware

import (
	"task-management-be/internal/pkg/db"
	"task-management-be/internal/pkg/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const bodyLimit string = "5M"

// SetUp sets up middleware for API.
func SetUp(server *echo.Echo, parentLogger *zap.Logger, config env.Config, dbClient *db.Client) {
	var (
		timeoutConfig = middleware.TimeoutConfig{Timeout: config.Timeout}
	)
	// Before
	server.Pre(middleware.TimeoutWithConfig(timeoutConfig))
	server.Pre(middleware.Secure())
	server.Pre(middleware.BodyLimit(bodyLimit))
	server.Pre(middleware.CORSWithConfig(CorsConfig))
	//nolint
	// server.Mux.Pre(sentryecho.New(sentryecho.Options{
	// 	Repanic: true,
	// }))

	// After
	server.Use(AddBuildVersion())
	server.Use(middleware.CORSWithConfig(CorsConfig))
	server.Use(middleware.Gzip())
	server.Use(middleware.Recover())
	server.Use(LogWithRequestID(parentLogger)...)
	server.Use(defaultBasicAuthConfig(config, dbClient))
}
