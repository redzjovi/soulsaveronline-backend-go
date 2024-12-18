package entity

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	ID          string         `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	ExpiredAt   time.Time      `gorm:"column:expired_at;index"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
