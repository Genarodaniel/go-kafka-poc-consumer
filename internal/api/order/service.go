package order

import (
	"context"
	"fmt"
	"go-kafka-order-producer/internal/repository"

	orderRepository "go-kafka-order-producer/internal/repository/order"
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
		StoreID:  order.StoreID,
		ClientID: order.ClientID,
	}

	result, err := a.OrderRepository.SaveOrder(ctx, orderEntity)
	if err != nil {
		return nil, fmt.Errorf("error to save order %s", err.Error())
	}

	return &PostOrderResponse{
		OrderID: result,
	}, nil
}
