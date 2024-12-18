package repository

import (
	"soulsaveronline-backend-go/internal/entity"
	"time"

	"gorm.io/gorm"
)

type DeviceRepository struct {
	Repository[entity.Device]
}

func (r *DeviceRepository) Now(db *gorm.DB) (*time.Time, error) {
	var now time.Time

	if err := db.Raw("SELECT NOW()").Scan(&now).Error; err != nil {
		return nil, err
	}

	return &now, nil
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{}
}
