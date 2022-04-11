package ItemActiveInWarehouse

import (
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	"prodapi/app/models"
	"prodapi/app/services/AppError"
	"prodapi/app/services/EntitySchema"

	// "prodapi/app/services/EntitySchema"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/qri-io/jsonschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"strconv"

	// "github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateWarehouseStatus(c *fiber.Ctx) error {
	var productinwarehouse models.WarehouseMain
	response := &models.Response{}

	// var productcheck EntitySchema.WarehouseValidation
	// c.BodyParser(&productcheck)

	err := c.BodyParser(&productinwarehouse)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}

	sku_id := productinwarehouse.Data.Sku_id
	account_id := productinwarehouse.Data.Account_id
	warehouse_id := productinwarehouse.Data.Warehouse_id
	//``````````````````````````````````````
	// in := make(map[string]string)
	// var validate = validator.New()
	// // productinwarehouse.Data.Entity_errors["any"] = "any"
	// verr := validate.Struct(productcheck)
	// if verr != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		// required_errors = append(required_errors, err.Field())
	// 		fmt.Println("get")
	// 		// key := err.StructField()
	// 		in["key"] = err.Field()
	// 	}
	// }
	// fmt.Println(in)
	// productinwarehouse.Data.Entity_errors = in

	var by []byte
	by, er := json.Marshal(productinwarehouse)
	_ = er
	_ = by

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(EntitySchema.SchemaData, rs); err != nil {
		fmt.Println("unmarshal schema: " + err.Error())
	}
	errs, err := rs.ValidateBytes(context.Background(), by)
	if err != nil {
		fmt.Println(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs)
	}

	//````````````````````````````````````````````````
	get := getAItem(sku_id, account_id, warehouse_id)
	if get {
		response.Status = true
		response.Message = "Product already exists in this Warehouse, please try to update quantity."
		return c.Status(http.StatusCreated).JSON(response)
	} else {
		productinwarehouse.Data.Last_sync_timestamp = time.Now()
		check := createActiveInWarehouse(&productinwarehouse.Data)
		if check {
			response.Status = true
			response.Message = "Warehouse status item Created"
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

func UpdateStatusWarehouse(c *fiber.Ctx) error {
	var updatestatusproduct *models.WarehouseMain
	response := &models.Response{}

	err := c.BodyParser(&updatestatusproduct)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}
	updatestatusproduct.Data.Last_sync_timestamp = time.Now()
	sku_id := updatestatusproduct.Data.Sku_id
	account_id := updatestatusproduct.Data.Account_id
	warehouse_id := updatestatusproduct.Data.Warehouse_id

	get := getAItem(sku_id, account_id, warehouse_id)
	if !get {
		response.Status = false
		response.Message = "No product found to warehouse with this sku_id"
		return c.Status(http.StatusCreated).JSON(response)
	} else {
		filter := bson.M{"sku_id": sku_id, "account_id": account_id, "warehouse_id": warehouse_id}
		updatestatus := updatestatusproduct.Data
		update := bson.M{"$set": updatestatus}
		check := updateAItem(&updatestatusproduct.Data, filter, update)
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

func GetWarehouseStatusList(c *fiber.Ctx) error {
	var products []models.Warehouse
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

	err := mgm.CollectionByName("warehouse").SimpleFind(&products, filter, pagination)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}

	if len(products) == 0 {
		response.Status = true
		response.Message = "No products in the database"
		return c.Status(http.StatusAccepted).JSON(response)
	}

	return c.Status(http.StatusAccepted).JSON(&fiber.Map{"status": http.StatusAccepted, "message": "Items fetched successfully", "data": products})
}

func GetWarehouseStatus(c *fiber.Ctx) error {
	var queryitem models.Warehouse
	c.QueryParser(&queryitem)

	var item *models.Warehouse

	sku_id := queryitem.Sku_id
	account_id := queryitem.Account_id
	warehouse_id := queryitem.Warehouse_id

	err := mgm.CollectionByName("warehouse").FindOne(context.Background(), bson.M{"sku_id": sku_id, "account_id": account_id, "warehouse_id": warehouse_id}).Decode(&item)
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

//**************************************************************************************************************

func getAItem(sku_id string, account_id string, warehouse_id string) bool {
	var product models.Warehouse
	mgm.CollectionByName("warehouse").FindOne(context.Background(), bson.M{"sku_id": sku_id, "account_id": account_id, "warehouse_id": warehouse_id}).Decode(&product)
	return product.Sku_id == sku_id
}

func createActiveInWarehouse(product *models.Warehouse) bool {
	err := mgm.CollectionByName("warehouse").Create(product)
	return err == nil
}

func updateAItem(updatestatusproduct *models.Warehouse, filter primitive.M, update primitive.M) bool {
	var updatedOne models.Warehouse
	err := mgm.CollectionByName("warehouse").FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedOne)
	return err == nil
}
