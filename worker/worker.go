package worker

import (
	"context"
	"database/sql"
	"go-kafka-poc-consumer/internal/events"
	"go-kafka-poc-consumer/internal/handlers/order"
	"go-kafka-poc-consumer/internal/infra/kafka"
	orderRepository "go-kafka-poc-consumer/internal/repository/order"

	geutils "github.com/GenaroDaniel/geutils/pkg/events"
)

var (
	OrderHandler *order.OrderHandler
)

type WorkerInterface interface {
	Start()
}

type Worker struct {
	Db         *sql.DB
	Kafka      kafka.KafkaInterface
	Dispatcher *geutils.EventDispatcher
}

func NewWorker(db *sql.DB, kafka kafka.KafkaInterface, dispatcher *geutils.EventDispatcher) WorkerInterface {
	return &Worker{
		Db:         db,
		Kafka:      kafka,
		Dispatcher: dispatcher,
	}
}

func initializateHandlers(
	db *sql.DB,
) {
	repository := orderRepository.NewOrderRepository(db)
	service := order.NewOrderService(repository)
	OrderHandler = order.NewOrderHandler(service)

}

func (w *Worker) Start() {
	ctx := context.Background()
	initializateHandlers(w.Db)

	w.Dispatcher.Register(ctx, string(events.CreateOrder), OrderHandler)
	w.Kafka.Consume(ctx)
}
