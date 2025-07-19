package handler

import (
	"context"
	"log/slog"
	"net/http"
	"subscriptions/internal/handler/dto"
	"subscriptions/internal/handler/mapper"
	"subscriptions/internal/model"

	"github.com/gin-gonic/gin"
)

type SubService interface {
	CreateSub(ctx context.Context, c *gin.Context)
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
		h.logger.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	pvz := mapper.CreateProductReqToPvz(req)
	product := mapper.CreateProductReqToProduct(req)

	result, err := h.productService.CreateProduct(c.Request.Context(), product, pvz)
	if err != nil {
		h.logger.Error("Failed to create product", "err", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Failed to create product"})
		return
	}

	//resp := mapper.ProductToCreateProductResp(result)

	//c.JSON(http.StatusCreated, resp)
}
