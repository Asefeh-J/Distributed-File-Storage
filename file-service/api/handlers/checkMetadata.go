package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
	"github.com/gin-gonic/gin"
)

func CheckMetadataHandler(c *gin.Context) {
	metadataURL := os.Getenv("METADATA_SERVICE_URL")
	resp, err := http.Get(metadataURL + "/health")
	if err != nil {
		logger.Inst().Error("Failed to contact metadata service")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "metadata service not reachable"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Inst().Error("Error reading response body")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response from metadata service"})
		return
	}

	var metadataResp map[string]interface{}
	if err := json.Unmarshal(body, &metadataResp); err != nil {
		logger.Inst().Error("Error parsing metadata service response")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid response from metadata service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Metadata service responded",
		"body":    metadataResp,
	})
}
