package db

import (
	"context"
	"errors"
	"fmt"
	"task-management-be/internal/generated/sql"
	"task-management-be/internal/pkg/logger"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// Database Client
type Client struct {
	*sql.Queries
	Pool *pgxpool.Pool
}

const (
	// 30s for connection timeout
	connectionContextTimeout = 30
	// 2 connection ticker
	connectionTickerInterval = 2
)

// Connect acquires a connection to the database
func New(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, connectionContextTimeout*time.Second)
	defer cancelCtx()

	ticker := time.NewTicker(connectionTickerInterval * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("connection to DB failed: %w", ctx.Err())
		default:
			db, err := pgxpool.New(ctx, connString)

			if db == nil || err != nil {
				fmt.Printf("Failed set DB: %v. Retrying...\n", err)
				continue
			}

			if db.Ping(ctx) != nil {
				fmt.Printf("Failed to ping DB: %v. Retrying...\n", err)
				continue
			}
			return db, nil
		}
	}

	return nil, fmt.Errorf("connection to DB failed: something unexpected happened")
}

func NewDBClient(ctx context.Context, connString string) (*Client, error) {
	db, err := New(ctx, connString)
	if err != nil {
		return nil, err
	}
	client := Client{
		Queries: sql.New(db),
		Pool:    db,
	}
	return &client, nil
}

// Transact SQL
func Transact[T any](ctx context.Context, c *Client, f func(*sql.Queries) (T, error)) (T, error) {
	var res T
	tx, err := c.Pool.Begin(ctx)
	if err != nil {
		return res, fmt.Errorf("unable to begin a new tx with error: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			logger.Logger.Error("unable to rollback tx", zap.Error(err))
		}
	}()

	queries := c.Queries.WithTx(tx)

	if res, err = f(queries); err != nil {
		return res, fmt.Errorf("unable to execute tx with error: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return res, fmt.Errorf("unable to commit tx with error: %w", err)
	}

	return res, nil
}
