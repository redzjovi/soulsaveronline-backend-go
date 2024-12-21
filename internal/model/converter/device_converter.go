package converter

import (
	"soulsaveronline-backend-go/internal/entity"
	"soulsaveronline-backend-go/internal/model"
)

func DeviceToModel(device *entity.Device) *model.Device {
	return &model.Device{
		ID:        device.ID,
		Name:      device.Name,
		ExpiredAt: device.ExpiredAt,
	}
}
