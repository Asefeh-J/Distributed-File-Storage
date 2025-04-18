package persistent

import (
	"fmt"
	"os"

	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
	"github.com/Asefeh-J/Distributed-File-Storage/shared/models"
	"gorm.io/gorm"
)

func GetDefaultDatabasePath() string {

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
}

// DBInstance holds the database connection instance.
var DBInstance *gorm.DB

// SetDatabase sets the database instance (call during initialization).
func SetDatabase(db *gorm.DB) {
	DBInstance = db
}

// GetDatabase returns the database instance.
func GetDatabase() *gorm.DB {
	return DBInstance
}

func MigrateDatabase() {
	db := GetDatabase() // Retrieve the database instance
	if db == nil {
		logger.Inst().Error("Database instance is not initialized.")
		return
	}

	// Automatically create or migrate the File table based on the model
	err := db.AutoMigrate(&models.File{})
	if err != nil {
		logger.Inst().Error("Failed to migrate database: " + err.Error())
		return
	}
	logger.Inst().Info("Database migration successful.")
}
