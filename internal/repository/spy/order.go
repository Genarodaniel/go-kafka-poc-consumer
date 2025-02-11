package repository

import (
	"context"
	"go-kafka-poc-consumer/internal/repository"
	"go-kafka-poc-consumer/internal/repository/order"
)

type OrderRepositorySpy struct {
	repository.OrderRepositoryInterface
	SaveOrderResponse string
	SaveOrderError    error
}

func (s OrderRepositorySpy) SaveOrder(ctx context.Context, order order.Order) (string, error) {
	return s.SaveOrderResponse, s.SaveOrderError
}
