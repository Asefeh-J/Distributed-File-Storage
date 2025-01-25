package api

import (
	"github.com/Asefeh-J/Distributed-File-Storage/file-service/api/routes"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8082")
}
