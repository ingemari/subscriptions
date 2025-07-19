package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(logger *slog.Logger, router *gin.Engine, port string, shutdownTimeout time.Duration) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		logger.Info("Starting server", "port", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	logger.Info("Received shutdown signal", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	logger.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", "error", err)
		os.Exit(1)
	}

	select {
	case <-ctx.Done():
		logger.Warn("Shutdown timeout exceeded")
	default:
		logger.Info("Server exited gracefully")
	}
}
