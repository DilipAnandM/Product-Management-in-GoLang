package routes

import (
	// "prodapi/app/services/Validation""
	"prodapi/app/services/ItemAssociation"

	"github.com/gofiber/fiber/v2"
)

func ItemAssociationRoutes(app fiber.Router) {
	app.Post("/create", ItemAssociation.CreateItemAssociation)
	// app.Put("/update_status", ItemAssociation.UpdateAssociationStatus)
	// app.Put("/update", ItemAssociation.UpdateItemAssociation)
	// app.Delete("/remove", ItemAssociation.RemoveItemAssociation)
	app.Get("/list", ItemAssociation.GetAssociationItemList)
	app.Get("/get", ItemAssociation.GetAssociationItem)
}
