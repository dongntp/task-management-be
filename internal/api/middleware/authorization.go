package middleware

import (
	"context"
	"slices"

	"task-management-be/internal/generated/sql"
	"task-management-be/internal/pkg/db"
	"task-management-be/internal/pkg/env"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthContextKeyType string

const usernameContextKey AuthContextKeyType = "username"

func adminAuthConfig(config env.Config) echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(_, password string, c echo.Context) (bool, error) {
			if password == string(config.AdminAccessToken) {
				setAuth(c, "admin")
				return true, nil
			}
			return false, nil
		},
	})
}

func defaultBasicAuthConfig(config env.Config, dbClient *db.Client, paths *openapi3.Paths) echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			reqPath := c.Request().URL.Path
			reqMethod := c.Request().Method

			path := paths.Find(reqPath)
			tags := path.Operations()[reqMethod].Tags

			params := sql.GetRoleByUserParams{
				Username: username,
				Password: password,
			}
			role, err := dbClient.GetRoleByUser(c.Request().Context(), params)
			if err != nil {
				return false, nil
			}

			if slices.Contains(tags, string(role)) {
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
