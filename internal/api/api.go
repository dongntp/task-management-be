package api

import (
	"context"
	"os"

	"task-management-be/internal/pkg/env"
	"task-management-be/internal/pkg/logger"

	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"
)

type API struct{}

func NewAPI(_ context.Context) *API {

	return &API{}
}

func loadPathsFromDocs(cfg env.Config) *openapi3.Paths {
	data, err := os.ReadFile(cfg.OpenAPIFilePath) // Change to "openapi.json" if needed
	if err != nil {
		logger.Logger.Fatal("Failed to read OpenAPI spec", zap.Error(err))
	}

	// Parse the OpenAPI spec
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		logger.Logger.Fatal("Failed to parse OpenAPI spec", zap.Error(err))
	}
	return doc.Paths
}
