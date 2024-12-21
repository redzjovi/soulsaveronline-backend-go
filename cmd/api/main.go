package main

import (
	"soulsaveronline-backend-go/internal/config"
	"soulsaveronline-backend-go/internal/delivery/http"
	"soulsaveronline-backend-go/internal/delivery/http/route"
	"soulsaveronline-backend-go/internal/repository"
	"soulsaveronline-backend-go/internal/usecase"
)

func main() {
	config.NewGodotenv()

	viper := config.NewViper()

	db := config.NewDB(viper)

	validator := config.NewValidator()

	deviceRepository := repository.NewDeviceRepository()

	deviceUsecase := usecase.NewDeviceUsecase(db, deviceRepository)

	deviceController := http.NewDeviceController(deviceUsecase, validator)

	app := config.NewFiber(viper)

	app = route.NewRoute(app, deviceController)

	app.Listen(viper.GetString("APP_ADDR"))
}
