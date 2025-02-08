package server

import (
	"database/sql"
	"go-kafka-order-producer/internal/api"

	"github.com/gin-gonic/gin"
)

func Init(db *sql.DB) *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	api.Router(router, db)

	return router
}
