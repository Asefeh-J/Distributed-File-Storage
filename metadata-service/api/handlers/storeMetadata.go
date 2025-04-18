package handlers

import (
	"net/http"
	"strconv"

	"github.com/Asefeh-J/Distributed-File-Storage/metadata-service/api/persistent"
	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
	"github.com/Asefeh-J/Distributed-File-Storage/shared/models"
	"github.com/gin-gonic/gin"
)

func StoreMetadataHandler(c *gin.Context) {
	name := c.PostForm("name")
	sizeStr := c.PostForm("size")
	metadata := c.PostForm("metadata") // optional

	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		logger.Inst().Error("Invalid file size")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file size"})
		return
	}

	file := models.File{
		Name:     name,
		Size:     size,
		Metadata: metadata,
	}

	db := persistent.GetDatabase()
	if db == nil {
		logger.Inst().Error("Database not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not available"})
		return
	}

	result := db.Create(&file)
	if result.Error != nil {
		logger.Inst().Error("Failed to store metadata")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not store metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metadata stored", "fileID": file.ID})
}
