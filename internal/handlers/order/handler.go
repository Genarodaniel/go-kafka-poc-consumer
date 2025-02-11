package order

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	geutils "github.com/GenaroDaniel/geutils/pkg/events"
)

type OrderHandlerInterface interface {
	Handle(event geutils.EventInterface, wg *sync.WaitGroup)
}

type OrderHandler struct {
	OrderService OrderServiceInterface
	log.Logger
}

func NewOrderHandler(orderService OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (h *OrderHandler) Handle(event geutils.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()
	request := &PostOrderRequest{}
	if err := json.Unmarshal(event.GetPayload(), request); err != nil {
		fmt.Printf("error to unserializee order payload: %s", err.Error())
		return
	}

	result, err := h.OrderService.PostOrder(ctx, request)
	if err != nil {
		fmt.Printf("error to create order: %s", err.Error())
		return
	}

	fmt.Printf("order created: %s", result.OrderID)

}
