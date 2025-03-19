package middleware

import (
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func AddBuildVersion() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			commitHash := getBuildInfo()
			if commitHash == "" {
				return next(c)
			}

			resp := c.Response()
			resp.Header().Add("Commit-Hash", commitHash)

			return next(c)
		}
	}
}

func getBuildInfo() (commitHash string) {
	var info, ok = debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			commitHash = setting.Value
		}
	}

	return commitHash
}
