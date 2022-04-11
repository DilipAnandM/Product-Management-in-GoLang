package routes

import (
	// "prodapi/app/services/Validation""
	"prodapi/app/services/Association"

	"github.com/gofiber/fiber/v2"
)

func AssociationRoutes(app fiber.Router) {
	app.Post("/create", Association.CreateAssociation)
	// app.Put("/update", ItemActiveInWarehouse.UpdateStatusWarehouse)
	// app.Post("/archive")
	app.Put("/update_status", Association.UpdateAssociationStatus)
	app.Get("/list", Association.GetAssociationList)
	app.Get("/get", Association.GetAssociation)
}
