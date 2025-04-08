
package models

import (
    "gorm.io/gorm"
	"time"
)

type Project struct {
    ID          uint64    `gorm:"primaryKey;autoIncrement"` // Explicit ID field
    Name        string    `gorm:"type:varchar(255);not null;uniqueIndex"` // Project name, unique index for faster lookups
    Status		string    `gorm:"type:text"`
    FileID      uint64    `gorm:"not null"` // Foreign key reference to File
    File        File      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Establishing the relationship with File struct
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"` // Soft delete support
}