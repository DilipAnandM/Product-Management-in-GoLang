package routes

import (
	"prodapi/app/services/ItemActiveInWarehouse"
	"prodapi/app/services/Validation"

	"github.com/gofiber/fiber/v2"
)

func WarehouseRoutes(app fiber.Router) {
	app.Post("/create", Validation.ValidateWarehouseItem, ItemActiveInWarehouse.CreateWarehouseStatus)
	app.Put("/update_status", ItemActiveInWarehouse.UpdateStatusWarehouse)
	app.Get("/list", ItemActiveInWarehouse.GetWarehouseStatusList)
	app.Get("/get", ItemActiveInWarehouse.GetWarehouseStatus)
}
