package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"task-management-be/internal/api"
	"task-management-be/internal/generated/openapi/server"
	"task-management-be/internal/pkg/env"
	"task-management-be/internal/pkg/logger"
	"task-management-be/internal/pkg/middleware"
)

const BasePath = "v1" // TODO: Move to config

func main() {
	ctx := context.Background()

	loadedConfig := env.GetConfig()

	taskAPI := api.NewAPI(ctx, loadedConfig)
	serverEcho := echo.New()

	middleware.SetUp(serverEcho, logger.Logger, loadedConfig, taskAPI.DBClient)

	handler := server.NewStrictHandler(taskAPI, nil)
	server.RegisterHandlersWithBaseURL(serverEcho, handler, BasePath)

	wrapper := server.ServerInterfaceWrapper{
		Handler: handler,
	}

	serverEcho.GET("/", wrapper.Healthcheck)

	serverEcho.Logger.Fatal(serverEcho.Start(fmt.Sprintf(":%d", 3000)))
}
