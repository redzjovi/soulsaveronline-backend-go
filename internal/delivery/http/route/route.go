package route

import (
	"soulsaveronline-backend-go/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func NewRoute(
	app *fiber.App,
	deviceController *http.DeviceController,
) *fiber.App {
	app.Use(healthcheck.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/device/:id", deviceController.FindByID)
	app.Post("/device/register", deviceController.Register)

	return app
}
