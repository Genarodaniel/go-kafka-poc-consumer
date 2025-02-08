package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	g.GET("/", HealthCheck)
}
