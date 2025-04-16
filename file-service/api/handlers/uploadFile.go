package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		logger.Inst().Error("File upload error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Save file locally
	if err := saveUploadedFile(header.Filename, file); err != nil {
		logger.Inst().Error("File save error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Send metadata
	if err := sendFileMetadata(header.Filename, header.Size); err != nil {
		logger.Inst().Error("Failed to send metadata: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func saveUploadedFile(filename string, file multipart.File) error {
	uploadDir := "uploads"
	os.MkdirAll(uploadDir, 0755)

	filePath := filepath.Join(uploadDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}

func sendFileMetadata(name string, size int64) error {
	metadataURL := os.Getenv("METADATA_SERVICE_URL") + "/store"

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	writer.WriteField("name", name)
	writer.WriteField("size", strconv.FormatInt(size, 10))
	writer.Close()

	req, err := http.NewRequest("POST", metadataURL, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("metadata service returned status: %s", resp.Status)
	}

	return nil
}

// =========================== Developer's note ==========================
// 1) Gin handles the incoming file upload request from the client.
// 2) Once the file is uploaded and saved, Gin is not needed for the next step.
// 3) The next step involves contacting the metadata-service to store the metadata.
//    For this, we need to send an HTTP request from our server to the metadata service.
//    Here, the net/http package is used because:

//       It provides a client to send HTTP requests (http.Client).

//       It allows us to construct the HTTP request (http.NewRequest).

//       It enables us to set headers, send the request, and receive the response.

// Using Gin directly to make an outgoing request would not be appropriate
// because it is focused on managing incoming requests.
// For this particular task (communicating with another service), we need net/http.
