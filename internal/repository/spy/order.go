package repository

import (
	"context"
	"go-kafka-order-producer/internal/repository"
	"go-kafka-order-producer/internal/repository/order"
)

type OrderRepositorySpy struct {
	repository.OrderRepositoryInterface
	SaveOrderResponse string
	SaveOrderError    error
}

func (s OrderRepositorySpy) SaveOrder(ctx context.Context, order order.Order) (string, error) {
	return s.SaveOrderResponse, s.SaveOrderError
}
