package route

import (
	"soulsaveronline-backend-go/internal/delivery/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewRoute(
	app *fiber.App,
	deviceController *http.DeviceController,
) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"now": time.Now(),
		})
	})
	app.Get("/device/:id", deviceController.FindByID)
	app.Patch("/device/:id", deviceController.Patch)
	app.Post("/device/register", deviceController.Register)

	return app
}
