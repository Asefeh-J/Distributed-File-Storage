package models

import "gorm.io/gorm"

type File struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Metadata   string `json:"metadata"` // JSON string to store metadata
}
