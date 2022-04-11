package routes

import "github.com/gofiber/fiber/v2"

func RoutesRegistry(app *fiber.App) {
	basepath := "/api/v2/product_management"
	startroute := app.Group(basepath)

	item_master := startroute.Group("/item")
	ItemRoutes(item_master)

	warehouse_status := startroute.Group("/warehouse_status")
	WarehouseRoutes(warehouse_status)

	marketplace_status := startroute.Group("/marketplace_status")
	MarketplaceRoutes(marketplace_status)

	association := startroute.Group("/association")
	AssociationRoutes(association)
}
