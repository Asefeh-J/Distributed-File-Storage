package main

import (
	"fmt"
	"os"

	"github.com/Asefeh-J/Distributed-File-Storage/metadata-service/api"
	"github.com/Asefeh-J/Distributed-File-Storage/metadata-service/api/persistent"
	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitLogger() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("can't get working directory: %v", err)
		os.Exit(-1)
	}
	logger.InitLog(path, "metadata-service.log")
	logger.Inst().Info("metadata-service logger initialized")

}

func InitDatabase() {
	connStr := persistent.GetDefaultDatabasePath()

	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Inst().Error(fmt.Sprintf("Failed to connect to database: %v", err))
		os.Exit(-1)
	}
	logger.Inst().Info("Database initialized successfully.")

	// Make the database accessible in the persistent package
	persistent.SetDatabase(db)
	// Call the MigrateDatabase function to handle migrations
	persistent.MigrateDatabase()
}
func Cleanup() {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Inst().Error(fmt.Sprintf("Failed to get underlying DB: %v", err))
		return
	}
	if err := sqlDB.Close(); err != nil {
		logger.Inst().Error(fmt.Sprintf("Error closing database connection: %v", err))
	} else {
		logger.Inst().Info("Database connection closed.")
	}
}

func Init() {
	InitLogger()
	InitDatabase()
}

func main() {
	fmt.Println("Starting application...")
	Init()
	defer Cleanup()
	fmt.Println("Starting API server...")
	api.StartServer()
}
