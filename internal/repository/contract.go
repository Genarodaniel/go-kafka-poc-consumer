package repository

import (
	"context"
	"go-kafka-poc-consumer/internal/repository/order"
)

type OrderRepositoryInterface interface {
	SaveOrder(ctx context.Context, order order.Order) (string, error)
}
