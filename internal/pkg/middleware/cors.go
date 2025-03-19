package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var (
	CorsConfig = middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}
)
