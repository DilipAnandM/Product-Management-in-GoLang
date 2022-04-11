package routes

import (
	"prodapi/app/services/ItemActiveInMarketplace"

	"github.com/gofiber/fiber/v2"
)

func MarketplaceRoutes(app fiber.Router) {
	app.Post("/create", ItemActiveInMarketplace.CreateMarketplaceStatus)
	app.Put("/update", ItemActiveInMarketplace.UpdateMarketplaceStatus)
	app.Put("/update_status", ItemActiveInMarketplace.UpdateStatusMarketplace)
	app.Get("/list", ItemActiveInMarketplace.GetMarketplaceStatusList)
	app.Get("/get", ItemActiveInMarketplace.GetMarketplaceStatus)
}
