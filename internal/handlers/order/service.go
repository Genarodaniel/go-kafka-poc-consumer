package order

import (
	"context"
	"fmt"
	"go-kafka-poc-consumer/internal/repository"

	orderRepository "go-kafka-poc-consumer/internal/repository/order"
)

type OrderServiceInterface interface {
	PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error)
}

type OrderService struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewOrderService(orderRepository repository.OrderRepositoryInterface) *OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
	}
}

func (a *OrderService) PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error) {
	orderEntity := orderRepository.Order{
		ID:                order.OrderID,
		StoreID:           order.StoreID,
		ClientID:          order.ClientID,
		NotificationEmail: order.NotificationEmail,
		Status:            order.Status,
	}

	result, err := a.OrderRepository.SaveOrder(ctx, orderEntity)
	if err != nil {
		return nil, fmt.Errorf("error to save order %s", err.Error())
	}

	return &PostOrderResponse{
		OrderID: result,
	}, nil
}
