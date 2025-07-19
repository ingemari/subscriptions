package db

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	slog.Info("Successfully connected to the database!")

	return pool
}
