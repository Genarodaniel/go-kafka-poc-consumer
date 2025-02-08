package repository

import (
	"context"
	"go-kafka-order-producer/internal/repository/order"
)

type OrderRepositoryInterface interface {
	SaveOrder(ctx context.Context, order order.Order) (string, error)
}
