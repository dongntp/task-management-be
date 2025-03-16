package api

import (
	"context"

	"task-management-be/internal/generated/openapi/server"
)

func (a *API) Healthcheck(_ context.Context, _ server.HealthcheckRequestObject) (server.HealthcheckResponseObject, error) {
	return server.Healthcheck200JSONResponse("OK"), nil
}
