package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"task-management-be/internal/api"
)

const BasePath = "v1" // TODO: Move to config

func main() {
	ctx := context.Background()

	_ = api.NewAPI(ctx)

	server := echo.New()
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", 3000)))
}
