package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckInterface interface {
	HealthCheck(c *gin.Context)
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Health!"})
}
