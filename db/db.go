package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"  // Or your preferred database driver
	"fengge/models"
	"log"
	"sync"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// GetDB returns the database instance
func GetDB() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("fengge.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Auto migrate the schemas
		err = db.AutoMigrate(&models.File{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}

		instance = db
	})
	return instance
} 