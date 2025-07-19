package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

func SetupRouter(db *pgxpool.Pool, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	// repo

	// service

	// repo

	// open routes
	//router.POST("/register", authHandler.HandleRegister)

	logger.Info("Endpoints registered")
	return router
}
