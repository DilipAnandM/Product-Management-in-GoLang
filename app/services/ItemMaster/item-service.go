package ItemMaster

import (
	"context"
	"encoding/json"

	// "encoding/json"

	// "encoding/json"
	"fmt"
	"net/http"
	"prodapi/app/models"

	"strconv"

	// "prodapi/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"

	"github.com/qri-io/jsonschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"prodapi/app/services"
	"prodapi/app/services/AppError"
	"prodapi/app/services/EntitySchema"
	// "prodapi/app/services/EntitySchema"
)

func CreateItem(c *fiber.Ctx) error {
	var product *models.MainProductData
	response := &models.Response{}
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}
	sku_id := product.Data.Product_identifiers.Sku_id
	account_id := product.Data.Account_id

	//``````````````````````````````````````````````````

	var by []byte
	by, er := json.Marshal(product)
	_ = er
	_ = by

	rs := &jsonschema.Schema{}
	e := json.Unmarshal(EntitySchema.CreateSchema, rs)
	if e != nil {
		fmt.Println(e)
	}
	errs, errr := rs.ValidateBytes(c.Context(), by)
	if errr != nil {
		fmt.Println(errr.Error())
	}
	if len(errs) > 0 {
		product.Data.Entity_errors[errs[0].Message] = errs[0].Error()
	}
	fmt.Println(errs)
	// fmt.Println(rs)
	// if unerr := json.Unmarshal(EntitySchema.CreateSchema, rs); unerr != nil {
	// 	fmt.Println("unmarshal schema: " + unerr.Error())
	// }

	//000000000000000000000000000000000000000000000000000000000000
	get := getAItem(sku_id, account_id, c)
	if get == true {
		response.Status = false
		response.Message = "A product already exist with this SKU ID."
		return c.Status(http.StatusCreated).JSON(response)
	} else {
		check := createAItem(&product.Data, c)
		// fmt.Println(check)
		if check["status"] {
			response.Status = true
			response.Message = "Product Created Successfully."
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

func InternalCreateItem(c *fiber.Ctx) error {
	response := &models.Response{}
	var internal_product models.MainProductData
	c.BodyParser(&internal_product)
	sku_id := internal_product.Data.Product_identifiers.Sku_id
	account_id := internal_product.Data.Account_id

	fetched_data := getAItem(sku_id, account_id, c)
	if fetched_data == true {
		response.Status = false
		response.Message = "A product already exist with this SKU ID."
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	createAItem := createAItem(&internal_product.Data, c)
	if createAItem["status"] {
		response.Status = true
		response.Message = "Product Created Successfully."
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

func GetItemList(c *fiber.Ctx) error {
	var product []models.Product
	err := mgm.CollectionByName("products").SimpleFind(&product, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}
	fmt.Println(len(product))
	if len(product) == 0 {
		return c.Status(http.StatusCreated).JSON(&fiber.Map{"status": http.StatusCreated, "message": "No products in the database", "data": product})
	}
	return c.Status(http.StatusCreated).JSON(&fiber.Map{"status": http.StatusCreated, "message": "product fetched successfully", "data": product})
}

func GetItem(c *fiber.Ctx) error {
	var product *models.Product
	query := c.Query("sku_id")
	filter := bson.M{"product_identifiers.sku_id": query}
	err := mgm.CollectionByName("products").FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(&fiber.Map{"status": http.StatusCreated, "message": "product fetched successfully", "data": product})
}

func UpdateItem(c *fiber.Ctx) error {
	var product *models.MainProductData
	response := &models.Response{}
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}
	product.Data.Is_archive = false

	sku_id := product.Data.Product_identifiers.Sku_id
	account_id := product.Data.Account_id

	get := getAItem(sku_id, account_id, c)
	if get != true {
		response.Status = false
		response.Message = "No Products with this SKU_ID found to Update."
		return c.Status(http.StatusCreated).JSON(response)
	} else {
		query := product.Data.Product_identifiers.Sku_id
		filter := bson.M{"product_identifiers.sku_id": query}
		check := updateAItem(&product.Data, filter, c)
		if check["status"] {
			response.Status = true
			response.Message = "Product Updated Successfully."
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

func DeleteItem(c *fiber.Ctx) error {
	response := &models.Response{}
	var deleteproduct *models.MainDeleteProduct
	c.BodyParser(&deleteproduct)
	sku_id := deleteproduct.Data.Product_identidifiers.Sku_id
	account_id := deleteproduct.Data.Account_id
	filter := bson.M{"product_identifiers.sku_id": sku_id, "account_id": account_id}

	fetched_item := getAItem(sku_id, account_id, c)
	if fetched_item != true {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}

	_, err := mgm.CollectionByName("products").DeleteOne(context.Background(), filter)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_UPDATE",
		}
		return AppError.AppError(errobj, c)
	}

	response.Status = true
	response.Message = "Item archived"
	return c.Status(http.StatusBadRequest).JSON(response)
}

func UpdateItemStatus(c *fiber.Ctx) error {
	response := &models.Response{}
	var updatestatusproduct models.MainUpdateStatus
	c.BodyParser(&updatestatusproduct)

	sku_id := updatestatusproduct.Data.Product_identifiers.Sku_id
	account_id := updatestatusproduct.Data.Account_id

	fetched_item := getAItem(sku_id, account_id, c)
	if fetched_item != true {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}

	status := updatestatusproduct.Data.Status

	filter := bson.M{"product_identifiers.sku_id": sku_id, "account_id": account_id}
	update := bson.M{"$set": bson.M{"status": status}}

	var result models.Product
	err := mgm.CollectionByName("products").FindOneAndUpdate(context.Background(), filter, update).Decode(&result)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_UPDATE",
		}
		return AppError.AppError(errobj, c)
	}
	response.Status = true
	response.Message = "Product Status Updated Successfully."
	return c.Status(http.StatusAccepted).JSON(response)
}

func UpdateItemFulfillmentDetails(c *fiber.Ctx) error {
	response := &models.Response{}
	var updateitem models.MainFulfillmentUpdate
	c.BodyParser(&updateitem)

	sku_id := updateitem.Data.Product_identidifiers.Sku_id
	account_id := updateitem.Data.Account_id

	filter := bson.M{"product_identifiers.sku_id": sku_id, "account_id": account_id}

	var old_product models.Product
	var empty models.Product
	res := mgm.CollectionByName("products").FindOne(context.Background(), filter).Decode(&old_product)
	if res != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}
	// if (models.Product{} == old_product)
	// check struct is empty or not
	if old_product.Account_id == empty.Account_id {
		response.Status = false
		response.Message = "No products with this sku_id found"
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	fulfillment_type_new := updateitem.Data.Fulfillment_details.Fulfillment_type
	fulfillment_type_old := old_product.Fulfillment_details.Fulfillment_type

	// not allowing changes of fulfillment type WFN/MFN to VSKU or vice-versa
	if fulfillment_type_old == "VSKU" {
		updateitem.Data.Fulfillment_details.Fulfillment_type = "VSKU"
	} else if fulfillment_type_new == "VSKU" {
		updateitem.Data.Fulfillment_details.Fulfillment_type = old_product.Fulfillment_details.Fulfillment_type
	}

	update_fulfillment := updateitem.Data.Fulfillment_details
	update := bson.M{"$set": bson.M{"fulfillment_details": update_fulfillment}}

	fmt.Println(update)
	var result models.Product
	update_err := mgm.CollectionByName("products").FindOneAndUpdate(context.Background(), filter, update).Decode(&result)

	if update_err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_UPDATE",
		}
		return AppError.AppError(errobj, c)
	}
	response.Status = true
	response.Message = "Product Fulfillment Details Updated Successfully."
	return c.Status(http.StatusAccepted).JSON(response)
}

func UpdateItemPrice(c *fiber.Ctx) error {
	response := &models.Response{}
	var updatepriceproduct models.MainUpdatePrice
	c.BodyParser(&updatepriceproduct)

	fmt.Println(updatepriceproduct)
	sku_id := updatepriceproduct.Data.Product_identifiers.Sku_id
	account_id := updatepriceproduct.Data.Account_id
	filter := bson.M{"product_identifiers.sku_id": sku_id, "account_id": account_id}

	var fetched_product models.Product
	err := mgm.CollectionByName("products").FindOne(context.Background(), filter).Decode(&fetched_product)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}
	price_det := updatepriceproduct.Data.Price_details
	update := bson.M{"$set": bson.M{"price_details": price_det}}

	fmt.Println(update)
	var updated_one models.Product
	update_err := mgm.CollectionByName("products").FindOneAndUpdate(context.Background(), filter, update).Decode(&updated_one)
	if update_err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_UPDATE",
		}
		return AppError.AppError(errobj, c)
	}

	response.Status = true
	response.Message = "Product Price Details Updated Successfully."
	return c.Status(http.StatusAccepted).JSON(response)
}

// ***************************************************************************************************************

func getAItem(sku_id string, account_id string, c *fiber.Ctx) interface{} {
	var productt *models.Product
	filter := bson.M{"product_identifiers.sku_id": sku_id}
	err := mgm.CollectionByName("products").FindOne(context.Background(), filter).Decode(&productt)
	if err != nil {
		errobj := AppError.Error{
			Status:     false,
			StatusCode: 406,
			ErrorCode:  "FAILED_TO_FETCH_DATA",
		}
		return AppError.AppError(errobj, c)
	}
	return productt.Product_identifiers.Sku_id == sku_id
}

func createAItem(product *models.Product, c *fiber.Ctx) map[string]bool {
	var catdet models.CategoryDetails
	var category int
	var sub_category int

	errors := make(map[string]string)

	// new
	// rs := &jsonschema.Schema{}
	// fmt.Println(rs)
	// if unerr := json.Unmarshal(EntitySchema.CreateSchema, rs); unerr != nil {
	// 	fmt.Println("unmarshal schema: " + unerr.Error())
	// }

	// var valid = []byte(product)
	// errs, valerr := rs.ValidateBytes(context.Background(), valid)
	// if valerr != nil {
	// 	fmt.Println(valerr)
	// }

	// if len(errs) > 0 {
	// 	// fmt.Println(errs[0].Error())
	// }

	// --------------
	if product.Category_details != catdet {
		if product.Category_details.Category_id != 0 {
			category = product.Category_details.Category_id
		}
		if product.Category_details.SubCategory_id != 0 {
			sub_category = product.Category_details.SubCategory_id
		}

		isCategoryDetailsValid := CheckCategoryDetails(category, sub_category)

		if isCategoryDetailsValid["status"] == "false" {
			_, isPresent := isCategoryDetailsValid["category"]
			_, isPresentsub := isCategoryDetailsValid["sub_category"]

			if isPresent && isPresentsub {
				errors["category_id"] = "Invalid category"
				errors["sub_category_id"] = "Invalid sub_category"
			} else {
				if isPresent {
					errors["category_id"] = "Invalid category"
				} else {
					errors["sub_category_id"] = "Invalid sub_category"
				}
			}
		}
	}

	var operation string = "itemDrafted"
	_ = operation // to avoid unused error

	// chech map is empty or not len(map_name) == 0
	if len(errors) != 0 {
		product.Entity_errors = errors
		// need to implement this
		// product.Is_item_valid = getItemValidReason(errors)
		product.New_item = true
	} else {
		product.Is_item_valid = "valid"
		product.New_item = false
		operation = "itemCreated"
	}
	if len(errors) == 0 && product.Product_identifiers.Product_id_value == "VSKU" {
		operation = "virtualSkuCreated"
	}
	res := make(map[string]bool)
	err := mgm.CollectionByName("products").Create(product)
	if err != nil {
		res["status"] = false
	} else {
		res["status"] = true
	}
	return res
}

func updateAItem(product_new *models.Product, filter primitive.M, c *fiber.Ctx) map[string]bool {
	var updatedOne models.Product
	var product_old models.Product

	// fetch the old product
	mgm.CollectionByName("products").FindOne(context.Background(), bson.M{"product_identifiers.sku_id": product_new.Product_identifiers.Sku_id}).Decode(&product_old)

	// s3 image
	var empty models.Product
	if product_new != &empty {
		basepath := product_new.Account_id + "/product_images/" + product_new.Product_identifiers.Sku_id + "/"
		// product_new.Images = utils.UploadImageToS3(product_new.Images, basepath)
		_ = basepath // unused

	}

	// dimension unnecessary
	if product_old.Item_dimension.Dimensions_verified {
		item_instructions := product_new.Item_dimension.Item_instructions
		product_new.Item_dimension.Item_instructions = item_instructions
	}

	isValidUpdate := true
	isVSKU := false

	fulfillment_type_new := product_new.Fulfillment_details.Fulfillment_type
	fulfillment_type_old := product_old.Fulfillment_details.Fulfillment_type

	if fulfillment_type_old == "VSKU" {
		isVSKU = true
	}
	if (fulfillment_type_new == "VSKU" && fulfillment_type_old != "VSKU") || (fulfillment_type_new != "VSKU" && fulfillment_type_old == "VSKU") {
		product_new.Fulfillment_details.Fulfillment_type = ""
		isValidUpdate = false
	}

	if isValidUpdate {
		isVSKU = (fulfillment_type_new == "VSKU")
	}
	_ = isVSKU //unused error

	// check the entity errors

	// final data to update
	data_to_update := bson.M{"$set": product_new}

	res := make(map[string]bool)
	update_err := mgm.CollectionByName("products").FindOneAndUpdate(context.Background(), filter, data_to_update).Decode(&updatedOne)
	if update_err != nil {
		res["status"] = false
	}
	res["status"] = true

	return res
}

func CheckCategoryDetails(category int, sub_category int) map[string]string {
	var categorieslist = services.GetCategories()

	res := make(map[string]string)

	cfou := false
	sfou := false
	for cat, sub_cat := range categorieslist {
		if cat == category {
			cfou = true
			for _, scat := range sub_cat {
				if scat == sub_category {
					sfou = true
				}
			}
		}
	}

	if !cfou {
		res["status"] = "false"
		res["category"] = strconv.Itoa(category)
	}
	if !sfou {
		res["status"] = "false"
		res["sub_category"] = strconv.Itoa(sub_category)
	}
	if cfou && sfou {
		res["status"] = "true"
	}

	return res
}
