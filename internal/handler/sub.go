package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"subscriptions/internal/handler/dto"
	"subscriptions/internal/handler/mapper"
	"subscriptions/internal/model"

	"github.com/gin-gonic/gin"
)

type SubService interface {
	CreateSub(ctx context.Context, sub model.Subscription) (model.Subscription, error)
	GetByID(ctx context.Context, id int) (model.Subscription, error)
	UpdateSubPrice(ctx context.Context, sub model.Subscription) (model.Subscription, error)
}

type SubHandler struct {
	subService SubService
	logger     *slog.Logger
}

func NewSubHandler(as SubService, logger *slog.Logger) *SubHandler {
	return &SubHandler{subService: as, logger: logger}
}

func (h *SubHandler) HandlerCreateSub(c *gin.Context) {
	var req dto.SubReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Failed to parse request", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := mapper.CreateReqToModel(req)
	if err != nil {
		h.logger.Warn("Failed to map request to model", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err = h.subService.CreateSub(c.Request.Context(), sub)
	if err != nil {
		h.logger.Error("Failed to create subscription", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.ModelToResp(sub)

	c.JSON(http.StatusCreated, resp)
}

func (h *SubHandler) HandlerGetSubs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Warn("Failed to map request to model", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.subService.GetByID(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("Failed to find subscription",
			slog.Int("id", id),
			slog.Any("err", err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.ModelToResp(sub)

	c.JSON(http.StatusOK, resp)
}

func (h *SubHandler) HandlerUpdateSubPrice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Warn("Failed to map request to model", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req dto.UpdatePriceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Failed to parse request", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub := mapper.UpdatePriceReqToModel(req, id)

	sub, err = h.subService.UpdateSubPrice(c.Request.Context(), sub)
	if err != nil {
		h.logger.Error("Failed to change subscription",
			slog.Int("id", id),
			slog.Any("err", err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.ModelToUpdatePriceResp(sub)
	c.JSON(http.StatusOK, resp)
}
