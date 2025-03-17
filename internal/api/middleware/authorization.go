package middleware

import (
	"task-management-be/internal/pkg/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func defaultBasicAuthConfig(config env.Config) echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(_, _ string, _ echo.Context) (bool, error) {
			return true, nil
		},
	})
}
