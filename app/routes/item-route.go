package routes

import (
	"prodapi/app/services/ItemMaster"
	"prodapi/app/services/Validation"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(app fiber.Router) {

	// main_route := app.Group("api/v2/product_management")
	// item_route := main_route.Group("/item")
	// item_route.Post("/create", utils.ValidateItem ,ItemMaster.CreateItem)
	// item_route.Get("/getall", ItemMaster.GetAllItems)
	// item_route.Get("/get", ItemMaster.GetSingleItem)
	// item_route.Put("/update", ItemMaster.UpdateItem)
	// item_route.Delete("/delete", ItemMaster.DeleteItem)

	app.Post("/create", Validation.ValidateCreateItem, ItemMaster.CreateItem)
	app.Post("/internal_create", ItemMaster.InternalCreateItem)
	app.Get("/list", ItemMaster.GetItemList)
	app.Get("/get", ItemMaster.GetItem)
	app.Put("/update", ItemMaster.UpdateItem)
	app.Delete("/delete", Validation.ValidateDeleteItem, ItemMaster.DeleteItem)

	app.Put("/update_status", ItemMaster.UpdateItemStatus)
	app.Put("/update_fulfillment_details", ItemMaster.UpdateItemFulfillmentDetails)
	app.Put("/update_price", ItemMaster.UpdateItemPrice)
}
