package order

import (
	"go-kafka-order-producer/internal/server/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandlerInterface interface {
	HandleGetAddressByZipCode(c *gin.Context)
}

type OrderHandler struct {
	OrderService OrderServiceInterface
}

func NewOrderHandler(orderService OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (h *OrderHandler) HandlePostOrder(ctx *gin.Context) {
	request := &PostOrderRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	if err := request.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	response, err := h.OrderService.PostOrder(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.DefaultErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
