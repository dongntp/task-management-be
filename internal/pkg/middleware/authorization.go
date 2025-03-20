package middleware

import (
	"context"
	"strings"

	"task-management-be/internal/generated/sql"
	"task-management-be/internal/pkg/db"
	"task-management-be/internal/pkg/env"
	"task-management-be/internal/pkg/hash"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthContextKeyType string

const usernameContextKey AuthContextKeyType = "username"

func defaultBasicAuthConfig(config env.Config, dbClient *db.Client) echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/" || strings.Contains(c.Request().URL.Path, "healthcheck")
		},
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if password == string(config.AdminAccessToken) {
				setAuth(c, "admin")
				return true, nil
			}

			user, err := dbClient.GetUserByUserName(c.Request().Context(), username)
			if err != nil {
				return false, nil
			}

			role := sql.RoleEnumEmployee
			if strings.Contains(c.Request().URL.Path, "employer") {
				role = sql.RoleEnumEmployer
			}

			if user.Active && user.Role == role && hash.CheckPasswordHash(password, user.Password) {
				setAuth(c, username)
				return true, nil
			}
			return false, nil
		},
	})
}

func setAuth(c echo.Context, username string) {
	req := c.Request()
	reqCtx := req.Context()

	reqCtxWithUsername := context.WithValue(reqCtx, usernameContextKey, username)

	c.SetRequest(req.WithContext(reqCtxWithUsername))
}

func GetUserName(ctx context.Context) *string {
	username := ctx.Value(usernameContextKey)
	if username == nil {
		return nil
	}
	rs := username.(string)
	return &rs
}
