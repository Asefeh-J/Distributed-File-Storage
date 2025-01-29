package routes

import (
	"github.com/Asefeh-J/Distributed-File-Storage/metadata-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheckHandler)

}
