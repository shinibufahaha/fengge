package models

import (
    "gorm.io/gorm"
	"time"
)

type File struct {
    ID         uint64   `gorm:"primaryKey;autoIncrement"` // Explicit ID field
    FilePath   string   `gorm:"type:varchar(255);not null;uniqueIndex"` // File path, unique index for faster lookups
    HashValue  string   `gorm:"type:varchar(64);not null"`
    CpgPath    string   `gorm:"type:varchar(255);not null;uniqueIndex"`           // Hash value of the file
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"` // Soft delete support
}