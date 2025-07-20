package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"subscriptions/internal/handler/dto"
	"subscriptions/internal/handler/mapper"
	"subscriptions/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type SubService interface {
	CreateSub(ctx context.Context, sub model.Subscription) (model.Subscription, error)
	GetByID(ctx context.Context, id int) (model.Subscription, error)
	UpdateSubPrice(ctx context.Context, sub model.Subscription) (model.Subscription, error)
	DeleteSub(ctx context.Context, id int) error
	ListSubs(ctx context.Context, sub model.Subscription) ([]model.Subscription, error)
	SumSubs(ctx context.Context, from, to time.Time) (int, error)
}

type SubHandler struct {
	subService SubService
	logger     *slog.Logger
}

func NewSubHandler(as SubService, logger *slog.Logger) *SubHandler {
	return &SubHandler{subService: as, logger: logger}
}

// ErrorResponse describes an error message
type ErrorResponse struct {
	Error string `json:"error"`
}

// HandlerCreateSub godoc
// @Summary Создать подписку
// @Description Создаёт новую подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body dto.SubReq true "Подписка"
// @Success 201 {object} model.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
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

// HandlerGetSubs godoc
// @Summary Получить подписку по ID
// @Description Возвращает подписку по идентификатору
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} model.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [get]
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

// HandlerUpdateSubPrice godoc
// @Summary Обновить цену подписки
// @Description Обновляет цену подписки по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Param price body dto.UpdatePriceRequest true "Новая цена"
// @Success 200 {object} model.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [put]
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

// HandlerDeleteSub godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *SubHandler) HandlerDeleteSub(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Warn("Failed to map request to model", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.subService.DeleteSub(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("Failed to delete subscription", slog.Int("id", id), slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted"})
}

// HandlerListSubs godoc
// @Summary Список подписок
// @Description Возвращает список подписок по user_id
// @Tags subscriptions
// @Produce json
// @Param user_id query string true "UUID пользователя"
// @Success 200 {array} model.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [get]
func (h *SubHandler) HandlerListSubs(c *gin.Context) {
	var req dto.ListReq
	req.UserID = c.Query("user_id")

	sub, err := mapper.ListReqToModel(req)
	if err != nil {
		h.logger.Warn("Failed to map request to model", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subs, err := h.subService.ListSubs(c.Request.Context(), sub)
	if err != nil {
		h.logger.Error("Failed to list subscriptions", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch subscriptions"})
		return
	}

	c.JSON(http.StatusOK, subs)
}

// HandlerSumSubs godoc
// @Summary Сумма подписок
// @Description Возвращает сумму подписок за период
// @Tags subscriptions
// @Produce json
// @Param start_date_from query string true "Дата начала (MM-YYYY)"
// @Param start_date_to query string true "Дата окончания (MM-YYYY)"
// @Success 200 {object} map[string]int
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/sum [get]
func (h *SubHandler) HandlerSumSubs(c *gin.Context) {
	var req dto.SumReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Warn("Failed to bind query params", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date_from and start_date_to are required"})
		return
	}

	from, err := mapper.ParseMonthYear(req.StartDateFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date_from"})
		return
	}

	to, err := mapper.ParseMonthYear(req.StartDateTo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date_to"})
		return
	}

	sum, err := h.subService.SumSubs(c.Request.Context(), from, to)
	if err != nil {
		h.logger.Error("Failed to calculate subscriptions sum", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate sum"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
