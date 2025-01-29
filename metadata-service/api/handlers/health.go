package handlers

import (
	"net/http"

	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	logger.Inst().Info("Health check endpoint accessed")
	c.JSON(http.StatusOK, gin.H{
		"status": "metadata-service is running",
	})
}
