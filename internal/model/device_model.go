package model

import "time"

type Device struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ExpiredAt time.Time `json:"expired_at"`
}

type RegisterDeviceRequest struct {
	ID string `json:"id" validate:"required"`
}

type PatchDeviceRequest struct {
	Name      *string    `json:"name"`
	ExpiredAt *time.Time `json:"expired_at"`
}
