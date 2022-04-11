package Association

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

func CreateAssociation(c *fiber.Ctx) error {
	response := &models.Response{}
	var product models.MainAssociationData
	c.BodyParser(&product)
	account_id := product.Data.Account_id
	parent_sku_id := product.Data.Parent_sku_id

	get := getAssociation(account_id, parent_sku_id)
	if get["status"] {
		response.Status = false
		response.Message = "Association with this parent product already exists."
		return c.Status(http.StatusBadRequest).JSON(response)
	} else {
		check := createAssociation(product.Data)
		if check["status"] {
			response.Status = true
			response.Message = "Association Created"
			return c.Status(http.StatusAccepted).JSON(response)
		} else {
			errobj := AppError.Error{
				Status:     false,
				StatusCode: 406,
				ErrorCode:  "FAILED_TO_CREATE_ASSOCIATION",
			}
			return AppError.AppError(errobj, c)
		}
	}
}

func UpdateAssociationStatus(c *fiber.Ctx) error {
	response := &models.Response{}
	var product models.MainAssociationData
	c.BodyParser(&product)
	account_id := product.Data.Account_id
	parent_sku_id := product.Data.Parent_sku_id
	get := getAssociation(account_id, parent_sku_id)
	if get["status"] {
		check := updateAssociationStatus(product.Data)
		if check["status"] {
			response.Status = true
			response.Message = "Updated association status"
			return c.Status(http.StatusAccepted).JSON(response)
		} else {
			errobj := AppError.Error{
				Status:     false,
				StatusCode: 406,
				ErrorCode:  "FAILED_TO_UPDATE",
			}
			return AppError.AppError(errobj, c)
		}
	} else {
		response.Status = false
		response.Message = "No association with this parent_sku_id to update"
		return c.Status(http.StatusBadRequest).JSON(response)
	}
}

func GetAssociationList(c *fiber.Ctx) error {
	var products []models.Association
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

	err := mgm.CollectionByName("association").SimpleFind(&products, filter, pagination)
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

func GetAssociation(c *fiber.Ctx) error {
	var queryitem models.Association
	c.QueryParser(&queryitem)
	var item models.Association
	account_id := queryitem.Account_id
	parent_sku_id := queryitem.Parent_sku_id
	filter := bson.M{"account_id": account_id, "parent_sku_id": parent_sku_id}
	err := mgm.CollectionByName("association").FindOne(context.Background(), filter).Decode(&item)
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

//--------------------------

func getAssociation(account_id string, parent_sku_id string) map[string]bool {
	res := make(map[string]bool)
	var old_one models.Association
	filter := bson.M{"account_id": account_id, "parent_sku_id": parent_sku_id}
	mgm.CollectionByName("association").FindOne(context.Background(), filter).Decode(&old_one)
	if old_one.Parent_sku_id == parent_sku_id {
		res["status"] = true
	} else {
		res["status"] = false
	}
	return res
}

func createAssociation(product models.Association) map[string]bool {
	res := make(map[string]bool)
	err := mgm.CollectionByName("association").Create(&product)
	if err != nil {
		res["status"] = false
	} else {
		res["status"] = true
	}
	return res
}

func updateAssociationStatus(product models.Association) map[string]bool {
	res := make(map[string]bool)
	parent_sku_id := product.Parent_sku_id
	filter := bson.M{"parent_sku_id": parent_sku_id}
	status := product.Status
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := mgm.CollectionByName("association").UpdateOne(context.Background(), filter, update)
	if err != nil {
		res["status"] = false
	} else {
		res["status"] = true
	}
	return res
}
