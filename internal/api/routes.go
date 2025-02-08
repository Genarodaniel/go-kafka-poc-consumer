package api

import (
	"database/sql"
	"go-kafka-order-producer/internal/api/healthcheck"
	"go-kafka-order-producer/internal/api/order"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine, db *sql.DB) {
	v1 := e.Group("/v1")

	orderGroup := v1.Group("/order")
	healthCheckGroup := v1.Group("/healthcheck")

	healthcheck.Router(healthCheckGroup)
	order.Router(orderGroup, db)
}
