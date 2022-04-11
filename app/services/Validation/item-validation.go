package Validation

import (
	"fmt"
	"net/http"
	"prodapi/app/models"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()
var required_errors []string

func ValidateCreateItem(c *fiber.Ctx) error {
	response := &models.Response{}
	var product *models.MainProductData
	c.BodyParser(&product)
	err := validate.Struct(product)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.Requiredfields = append(response.Requiredfields, err.Field())
		}
		response.Message = "Required fields are missing"
		return c.Status(http.StatusAccepted).JSON(response)
	}
	return c.Next()
}

func ValidateWarehouseItem(c *fiber.Ctx) error {
	var productinwarehouse models.WarehouseMain
	c.BodyParser(&productinwarehouse)
	err := validate.Struct(productinwarehouse.Data)
	fmt.Println(err)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			required_errors = append(required_errors, err.Field())
		}
		return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": false, "message": "Required fields are missing", "missing_fields": required_errors})
	}
	return c.Next()
}

func ValidateDeleteItem(c *fiber.Ctx) error {
	var deleteproduct models.MainDeleteProduct
	c.BodyParser(&deleteproduct)
	err := validate.Struct(deleteproduct.Data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			required_errors = append(required_errors, err.Field())
		}
		return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": false, "message": "Required fields are missing", "missing_fields": required_errors})
	}
	return c.Next()
}
