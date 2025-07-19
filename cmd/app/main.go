package main

import (
	"log/slog"
	"subscriptions/internal/config"
	"subscriptions/internal/db"
	"subscriptions/internal/handler/router"
	"subscriptions/internal/middleware/logs"
	"subscriptions/internal/server"
	"time"
)

func main() {
	logger := logs.SetupLogger()
	slog.SetDefault(logger)

	cfg := config.LoadConfig()
	logger.Info("Configuration loaded", "config", cfg)

	database := db.InitDB(config.MakeDSN(*cfg))
	defer database.Close()
	logger.Info("Database connection established")

	r := router.SetupRouter(database, logger)

	server.Run(
		logger,
		r,
		cfg.ServerPort,
		30*time.Second,
	)
}
