package router

import (
	"subscriptions/internal/handler"
	"subscriptions/internal/repository"
	"subscriptions/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

func SetupRouter(db *pgxpool.Pool, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	// repo
	subRepo := repository.NewSubRepository(db, logger)
	// service
	subService := service.NewSubService(subRepo, logger)
	// handler
	subHandler := handler.NewSubHandler(subService, logger)

	router.POST("/subscriptions", subHandler.HandlerCreateSub)
	router.GET("/subscriptions/:id", subHandler.HandlerGetSubs)
	router.PUT("/subscriptions/:id", subHandler.HandlerUpdateSubPrice)

	logger.Info("Endpoints registered")
	return router
}
