package http

import (
	"errors"
	"log"
	"soulsaveronline-backend-go/internal/model"
	"soulsaveronline-backend-go/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DeviceController struct {
	deviceUsecase *usecase.DeviceUsecase
	validate      *validator.Validate
}

func (c *DeviceController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	device, err := c.deviceUsecase.FindById(ctx.UserContext(), id)

	if err != nil {
		log.Printf("deviceUsecase.FindByID: %v", err)
		return fiber.ErrInternalServerError
	} else if device == nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(model.WebResponse[*model.Device]{Data: device})
}

func (c *DeviceController) Register(ctx *fiber.Ctx) error {
	request := model.RegisterDeviceRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": NewMapErrorJson(err),
		})
	}

	device, err := c.deviceUsecase.Register(ctx.UserContext(), request.ID)

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return fiber.ErrConflict
	} else if err != nil {
		log.Printf("deviceUsecase.Register: %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.Device]{Data: device})
}

func (c *DeviceController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := model.PatchDeviceRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": NewMapErrorJson(err),
		})
	}

	device, err := c.deviceUsecase.Patch(ctx.UserContext(), id, request)

	if err != nil {
		log.Printf("deviceUsecase.Patch: %v", err)
		return fiber.ErrInternalServerError
	} else if device == nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(model.WebResponse[*model.Device]{Data: device})
}

func NewDeviceController(
	deviceUsecase *usecase.DeviceUsecase,
	validate *validator.Validate,
) *DeviceController {
	return &DeviceController{
		deviceUsecase: deviceUsecase,
		validate:      validate,
	}
}
