package order

import (
	"database/sql"
	orderRepository "go-kafka-order-producer/internal/repository/order"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, db *sql.DB) {
	repository := orderRepository.NewOrderRepository(db)
	service := NewOrderService(repository)
	handler := NewOrderHandler(service)

	g.POST("/", handler.HandlePostOrder)
}
