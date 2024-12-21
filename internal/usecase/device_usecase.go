package usecase

import (
	"context"
	"errors"
	"fmt"
	"soulsaveronline-backend-go/internal/entity"
	"soulsaveronline-backend-go/internal/model"
	"soulsaveronline-backend-go/internal/model/converter"
	"soulsaveronline-backend-go/internal/repository"
	"time"

	"gorm.io/gorm"
)

type DeviceUsecase struct {
	db               *gorm.DB
	deviceRepository *repository.DeviceRepository
}

func (u *DeviceUsecase) FindById(ctx context.Context, id string) (*model.Device, error) {
	e := entity.Device{}

	err := u.deviceRepository.FindById(u.db, &e, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("deviceRepository.FindById: %v", err)
	}

	return converter.DeviceToModel(&e), nil
}

func (u *DeviceUsecase) Register(ctx context.Context, id string) (*model.Device, error) {
	m, err := u.FindById(ctx, id)

	if m != nil {
		return nil, gorm.ErrDuplicatedKey
	} else if err != nil {
		return nil, err
	}

	e := entity.Device{
		ID:        id,
		ExpiredAt: time.Now().AddDate(0, 0, 1),
	}

	if err = u.deviceRepository.Create(u.db, &e); err != nil {
		return nil, fmt.Errorf("deviceRepository.Create: %v", err)
	}

	return converter.DeviceToModel(&e), nil
}

func (u *DeviceUsecase) Patch(ctx context.Context, id string, req model.PatchDeviceRequest) (*model.Device, error) {
	e := entity.Device{}

	err := u.deviceRepository.FindById(u.db, &e, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("deviceRepository.FindById: %v", err)
	}

	if req.Name != nil {
		e.Name = *req.Name
	}

	if req.ExpiredAt != nil {
		e.ExpiredAt = *req.ExpiredAt
	}

	if err = u.deviceRepository.Update(u.db, &e); err != nil {
		return nil, fmt.Errorf("deviceRepository.Update: %v", err)
	}

	return converter.DeviceToModel(&e), nil
}

func NewDeviceUsecase(
	db *gorm.DB,
	deviceRepository *repository.DeviceRepository,
) *DeviceUsecase {
	return &DeviceUsecase{
		db:               db,
		deviceRepository: deviceRepository,
	}
}
