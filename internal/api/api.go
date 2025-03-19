package api

import (
	"context"

	"task-management-be/internal/pkg/db"
	"task-management-be/internal/pkg/env"
	"task-management-be/internal/pkg/logger"

	"go.uber.org/zap"
)

type API struct {
	DBClient *db.Client
}

func NewAPI(ctx context.Context, loadedConfig env.Config) *API {
	dbClient, err := db.NewDBClient(ctx, string(loadedConfig.DBConnectionString))
	if err != nil {
		logger.Logger.Fatal("cannot get db connection pool:", zap.Error(err))
	}

	return &API{
		DBClient: dbClient,
	}
}
