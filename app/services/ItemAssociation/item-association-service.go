package ItemAssociation

import (
	"context"
	"net/http"
	"prodapi/app/models"
	"prodapi/app/services/AppError"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateItemAssociation(c *fiber.Ctx) error {
	response := &models.Response{}
	var product models.MainItemAssociationData
	c.BodyParser(&product)
	account_id := product.Data.Account_id
	sku_id := product.Data.Sku_id
	parent_sku_id := product.Data.Parent_sku_id

	get := getAssociationItem(account_id, sku_id, parent_sku_id)
	if get["status"] {
		response.Status = false
		response.Message = "Association has already been created."
		return c.Status(http.StatusBadRequest).JSON(response)
	} else {
		check := createIterAssociation(product.Data)
		if check["status"] {
			response.Status = true
			response.Message = "Associated item details record created"
			return c.Status(http.StatusAccepted).JSON(response)
		} else {
			errobj := AppError.Error{
				Status:     false,
				StatusCode: 406,
				ErrorCode:  "FAILED_TO_CREATE",
			}
			return AppError.AppError(errobj, c)
		}
	}
}

// func UpdateAssociationStatus(c *fiber.Ctx) error {

// }

// func UpdateItemAssociation(c *fiber.Ctx) error {

// }

// func RemoveItemAssociation(c *fiber.Ctx) error {

// }

func GetAssociationItemList(c *fiber.Ctx) error {
	var products []models.ItemAssociation
	response := &models.Response{}

	limit, _ := strconv.Atoi(c.Query("limit"))
	pageno, _ := strconv.Atoi(c.Query("pageno"))

	var skipPages int
	if pageno > 0 {
		skipPages = (pageno - 1) * limit
	} else {
		skipPages = 0
	}

	var que models.Query
	c.QueryParser(&que)
	var filter primitive.M
	if que.Status == "" {
		filter = bson.M{}
	} else {
		filter = bson.M{"status": que.Status}
	}
	pagination := options.Find().SetSkip(int64(skipPages)).SetLimit(int64(limit))

	err := mgm.CollectionByName("item_association").SimpleFind(&products, filter, pagination)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}

	if len(products) == 0 {
		response.Status = true
		response.Message = "No products in the database"
		return c.Status(http.StatusAccepted).JSON(response)
	}

	return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": http.StatusAccepted, "message": "Associations fetched successfully", "data": products})
}

func GetAssociationItem(c *fiber.Ctx) error {
	var queryitem models.ItemAssociation
	c.QueryParser(&queryitem)
	var item models.ItemAssociation
	account_id := queryitem.Account_id
	sku_id := queryitem.Sku_id
	parent_sku_id := queryitem.Parent_sku_id
	filter := bson.M{"account_id": account_id, "sku_id": sku_id, "parent_sku_id": parent_sku_id}
	err := mgm.CollectionByName("item_association").FindOne(context.Background(), filter).Decode(&item)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}
	return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": http.StatusAccepted, "message": "Association fetched successfully", "data": item})
}

//-----------------------------------------------------------

func getAssociationItem(account_id string, sku_id string, parent_sku_id string) map[string]bool {
	res := make(map[string]bool)
	var old_product models.ItemAssociation
	filter := bson.M{"account_id": account_id, "sku_id": sku_id, "parent_sku_id": parent_sku_id}
	mgm.CollectionByName("item_association").FindOne(context.Background(), filter).Decode(&old_product)
	if old_product.Parent_sku_id == parent_sku_id {
		res["status"] = true
	} else {
		res["status"] = false
	}
	return res
}

func createIterAssociation(product models.ItemAssociation) map[string]bool {
	res := make(map[string]bool)
	err := mgm.CollectionByName("item_association").Create(&product)
	if err != nil {
		res["status"] = false
	} else {
		res["status"] = true
	}
	return res
}
