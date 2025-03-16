package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"task-management-be/internal/api"
	"task-management-be/internal/generated/openapi/server"
)

const BasePath = "v1" // TODO: Move to config

func main() {
	ctx := context.Background()

	taskAPI := api.NewAPI(ctx)
	serverEcho := echo.New()

	handler := server.NewStrictHandler(taskAPI, nil)
	server.RegisterHandlersWithBaseURL(serverEcho, handler, BasePath)

	wrapper := server.ServerInterfaceWrapper{
		Handler: handler,
	}

	serverEcho.GET("/", wrapper.Healthcheck)

	serverEcho.Logger.Fatal(serverEcho.Start(fmt.Sprintf(":%d", 3000)))
}
