package routes

import (
	"github.com/Asefeh-J/Distributed-File-Storage/file-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheckHandler)
	router.POST("/upload", handlers.UploadFileHandler)
	router.GET("/check-metadata", handlers.CheckMetadataHandler)

}
