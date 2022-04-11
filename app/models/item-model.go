package models

import (
	"github.com/kamva/mgm/v3"
)

type MainProductData struct {
	Data Product `json:"data" validate:"required"`
}

type Product struct {
	mgm.DefaultModel    `bson:",inline"`
	Account_id          string             `json:"account_id" bson:"account_id" validate:"required,omitempty"`
	Client_code         string             `json:"client_code" bson:"client_code,omitempty"`
	Product_identifiers ProductIdentifiers `json:"product_identifiers" bson:"product_identifiers,omitempty"`
	Product_details     ProductDetails     `json:"product_details" bson:"product_details,omitempty"`
	Category_details    CategoryDetails    `json:"category_details" bson:"category_details,omitempty"`
	Package_details     PackageDetails     `json:"package_details" bson:"package_details,omitempty"`
	Images              Image              `json:"images" bson:"images"`
	Fulfillment_details FulfillmentDetails `json:"fulfillment_details" bson:"fulfillment_details,omitempty"`
	Price_details       PriceDetails       `json:"price_details" bson:"price_details,omitempty"`
	Item_dimension      ItemDimension      `json:"item_dimension" bson:"item_dimension,omitempty"`
	Is_item_valid       string             `json:"is_item_valid" bson:"is_item_valid,omitempty"`
	Status              string             `json:"status" bson:"status,omitempty"`
	Is_archive          bool               `json:"is_archive" bson:"is_archive,omitempty"`
	New_item            bool               `json:"new_item" bson:"new_item,omitempty"`
	Entity_errors       map[string]string  `json:"entity_errors" bson:"entity_errors,omitempty"`
}

// "required,min=3,max=12"
// validate:"required,numeric"

type ProductIdentifiers struct {
	Sku_id           string `json:"sku_id" bson:"sku_id,omitempty" validate:"required"`
	Product_id_type  string `json:"product_id_type" bson:"product_id_type,omitempty"`
	Product_id_value string `json:"product_id_value" bson:"product_id_value,omitempty"`
}

type ProductDetails struct {
	Title          string `json:"title" bson:"title,omitempty"`
	Description    string `json:"description" bson:"description,omitempty"`
	Size           int    `json:"size" bson:"size,omitempty"`
	Color          string `json:"color" bson:"color,omitempty"`
	Condition_type string `json:"condition_type" bson:"condition_type,omitempty"`
}

type CategoryDetails struct {
	Category_id      int    `json:"category_id" bson:"category_id,omitempty"`
	Category_name    string `json:"category_name" bson:"category_name,omitempty"`
	SubCategory_id   int    `json:"sub_category_id" bson:"sub_category_id,omitempty"`
	SUbCategory_name string `josn:"sub_category_name" bson:"sub_category_name,omitempty"`
}

type PackageDetails struct {
	Pakage_length      int  `json:"package_length" bson:"package_length,omitempty"`
	Package_width      int  `json:"package_width" bosn:"package_width,omitempty"`
	Package_volume     int  `josn:"package_volume" bson:"package_volume,omitempty"`
	Dimension_verified bool `jons:"dimension_verified" bson:"dimension_verified"`
}

type FulfillmentDetails struct {
	Storage_type       string `json:"storage_type" bson:"storage_type,omitempty"`
	Fulfillment_type   string `json:"fulfillment_type" bson:"fulfillment_type,omitempty"`
	Lot_required       int    `json:"lot_required" bson:"lot_required,omitempty"`
	Serial_no_required int    `json:"serial_no_required" bson:"serial_no_required,omitempty"`
	Instruction        string `json:"instruction" bson:"instruction,omitempty"`
}

type PriceDetails struct {
	Retail_price   int `json:"max_retail_price" bson:"retail_price,omitempty"`
	Selling_price  int `json:"selling_price" bson:"selling_price,omitempty"`
	Sourcing_price int `json:"sourcing_price" bson:"sourcing_price,omitempty"`
}

type Image struct {
	Front_view string `json:"front_view" bson:"front_view"`
	Left_view  string `json:"left_view" bson:"left_view"`
	Right_view string `json:"right_view" bson:"right_view"`
	Top_view   string `json:"top_view" bson:"top_view"`
	Back_view  string `json:"back_view" bson:"back_view"`
}

type ItemDimension struct {
	Item_length         int    `json:"item_length" bson:"item_length,omitempty"`
	Item_width          int    `json:"item_width" bson:"item_width,omitempty"`
	Item_height         int    `json:"item_height" bson:"item_height,omitempty"`
	Item_weight         int    `json:"item_weight" bson:"item_weight,omitempty"`
	Item_volume         int    `json:"item_volume" bson:"item_volume,omitempty"`
	Item_instructions   string `json:"item_instructions" bson:"item_instructions,omitempty"`
	Dimensions_verified bool   `json:"dimensions_verified" bson:"dimensions_verified,omitempty"`
}

// common
type ProductIdentifiersUS struct {
	Sku_id string `json:"sku_id" bson:"sku_id"`
}

// update status
type MainUpdateStatus struct {
	Data UpdateStatus `json:"data" bson:"data,omitempty"`
}

type UpdateStatus struct {
	Product_identifiers ProductIdentifiersUS `json:"product_identifiers" bson:"product_identifiers"`
	Status              string               `json:"status" bson:"status"`
	Account_id          string               `json:"account_id" bson:"account_id"`
}

//update price
type MainUpdatePrice struct {
	Data UpdatePrice `json:"data" bson:"data"`
}

type UpdatePrice struct {
	Account_id          string               `json:"account_id" bson:"account_id"`
	Product_identifiers ProductIdentifiersUS `json:"product_identifiers" bson:"product_identifiers"`
	Price_details       PriceDetails         `json:"price_details" bson:"price_details"`
}

// delete
type MainDeleteProduct struct {
	Data DeleteProduct `json:"data" validate:"required"`
}
type DeleteProduct struct {
	Product_identidifiers ProductIdentifiersUS `json:"product_identifiers" validate:"required"`
	Account_id            string               `json:"account_id" validate:"required"`
}

//  update fulfillment details
type MainFulfillmentUpdate struct {
	Data UpdateFulfillment `json:"data" bson:"data"`
}

type UpdateFulfillment struct {
	Account_id            string               `json:"account_id" bson:"account_id"`
	Product_identidifiers ProductIdentifiersUS `json:"product_identifiers" bson:"product_identifiers"`
	Fulfillment_details   FulfillmentDetails   `json:"fulfillment_details" bson:"fulfillment_details"`
}

type Res struct {
	Status bool `json:"status,omitempty" bson:"status"`
}
