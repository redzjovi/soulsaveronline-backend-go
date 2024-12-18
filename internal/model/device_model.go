package model

import "time"

type Device struct {
	ID        string    `json:"id"`
	ExpiredAt time.Time `json:"expired_at"`
}

type RegisterDeviceRequest struct {
	ID string `json:"id" validate:"required"`
}
