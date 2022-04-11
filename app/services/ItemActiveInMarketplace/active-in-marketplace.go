package ItemActiveInMarketplace

import (
	"context"
	"fmt"
	"net/http"
	"prodapi/app/models"
	"prodapi/app/services/AppError"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMarketplaceStatus(c *fiber.Ctx) error {
	var productinmarket models.MainMarketplaceData
	c.BodyParser(&productinmarket)

	response := &models.Response{}
	sku_id := productinmarket.Data.Sku_id
	account_id := productinmarket.Data.Account_id
	channel_id := productinmarket.Data.Channel_id
	store_id := productinmarket.Data.Store_id

	get := getAItem(sku_id, account_id, channel_id, store_id)
	if get {
		response.Status = true
		response.Message = "Product already exists in this marketplace."
		return c.Status(http.StatusCreated).JSON(response)
	} else {
		productinmarket.Data.Last_sync_timestamp = time.Now()
		check := createActiveInMarketplace(&productinmarket.Data)
		if check {
			response.Status = true
			response.Message = "Marketplace status item Created"
			return c.Status(http.StatusCreated).JSON(response)
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

func UpdateMarketplaceStatus(c *fiber.Ctx) error {
	var product models.MainMarketplaceData
	c.BodyParser(&product)
	response := &models.Response{}
	sku_id := product.Data.Sku_id
	account_id := product.Data.Account_id
	channel_id := product.Data.Channel_id
	store_id := product.Data.Store_id
	product.Data.Last_sync_timestamp = time.Now()
	fmt.Println(product)
	filter := bson.M{"sku_id": sku_id, "account_id": account_id, "channel_id": channel_id, "store_id": store_id}
	update := bson.M{"$set": product.Data}
	var updatedOne models.Marketplace
	err := mgm.CollectionByName("marketplace").FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedOne)
	fmt.Println(updatedOne)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_UPDATE",
		}
		return AppError.AppError(errobj, c)
	}
	response.Status = true
	response.Message = "Item Updated"
	return c.Status(http.StatusAccepted).JSON(response)
}

func UpdateStatusMarketplace(c *fiber.Ctx) error {
	var product models.MainMarketplaceData
	c.BodyParser(&product)
	response := &models.Response{}

	sku_id := product.Data.Sku_id
	account_id := product.Data.Account_id
	channel_id := product.Data.Channel_id
	store_id := product.Data.Store_id

	product.Data.Last_sync_timestamp = time.Now()
	status := product.Data.Status
	get := getAItem(sku_id, account_id, channel_id, store_id)
	if !get {
		response.Status = false
		response.Message = "No documents with this SKU_ID in the database"
		return c.Status(http.StatusBadRequest).JSON(response)
	} else {
		filter := bson.M{"sku_id": sku_id, "account_id": account_id, "channel_id": channel_id, "store_id": store_id}
		update := bson.M{"$set": bson.M{"status": status, "last_sync_timestamp": product.Data.Last_sync_timestamp}}
		check := updateAItem(filter, update)
		if check {
			response.Status = true
			response.Message = "Item Status Updated"
			return c.Status(http.StatusAccepted).JSON(response)
		} else {
			errobj := AppError.Error{
				Status:     false,
				StatusCode: 406,
				ErrorCode:  "FAILED_TO_UPDATE",
			}
			return AppError.AppError(errobj, c)
		}
	}
}

func GetMarketplaceStatusList(c *fiber.Ctx) error {
	var products []models.Marketplace
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

	err := mgm.CollectionByName("marketplace").SimpleFind(&products, filter, pagination)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}
	if len(products) == 0 {
		response.Status = true
		response.Message = "No products in the database"
		return c.Status(http.StatusAccepted).JSON(response)
	}

	return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": http.StatusAccepted, "message": "Items fetched successfully", "data": products})
}

func GetMarketplaceStatus(c *fiber.Ctx) error {
	var queryitem models.Marketplace
	c.QueryParser(&queryitem)

	var item models.Marketplace

	sku_id := queryitem.Sku_id
	account_id := queryitem.Account_id
	channel_id := queryitem.Channel_id
	store_id := queryitem.Store_id

	filter := bson.M{"sku_id": sku_id, "account_id": account_id, "channel_id": channel_id, "store_id": store_id}
	err := mgm.CollectionByName("marketplace").FindOne(context.Background(), filter).Decode(&item)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}

	return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": http.StatusAccepted, "message": "Item fetched successfully", "data": item})
}

//--------------------------------------------------//

func getAItem(sku_id string, account_id string, channel_id string, store_id string) bool {
	var product models.Marketplace
	filter := bson.M{"sku_id": sku_id, "account_id": account_id, "channel_id": channel_id, "store_id": store_id}
	mgm.CollectionByName("marketplace").FindOne(context.Background(), filter).Decode(&product)
	return product.Sku_id == sku_id
}

func createActiveInMarketplace(product *models.Marketplace) bool {
	err := mgm.CollectionByName("marketplace").Create(product)
	return err == nil
}

func updateAItem(filter primitive.M, update primitive.M) bool {
	var updatedOne models.Marketplace
	err := mgm.CollectionByName("marketplace").FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedOne)
	return err == nil
}
