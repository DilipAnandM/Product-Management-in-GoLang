package services

type ItemMasterCreateSchema struct {
	Data Product `json:"data" validate:"required,dive,required"`
}

type Product struct {
	Account_id          string             `json:"account_id"`
	Client_code         string             `json:"client_code"`
	Product_identifiers ProductIdentifiers `json:"product_identifiers"`
	Product_details     ProductDetails     `json:"product_details"`
	Category_details    CategoryDetails    `json:"category_details"`
	Package_details     PackageDetails     `json:"package_details"`
	Images              Image              `json:"images"`
	Fulfillment_details FulfillmentDetails `json:"fulfillment_details"`
	Price_details       PriceDetails       `json:"price_details"`
	Item_dimension      ItemDimension      `json:"item_dimension"`
	Is_item_valid       string             `json:"is_item_valid"`
	Status              string             `json:"status"`
	Is_archive          bool               `json:"is_archive"`
	New_item            bool               `json:"new_item"`
	Entity_errors       map[string]string  `json:"entity_errors"`
}

// "required,min=3,max=12"
// validate:"required,numeric"

type ProductIdentifiers struct {
	Sku_id           string `json:"sku_id"`
	Product_id_type  string `json:"product_id_type"`
	Product_id_value string `json:"product_id_value"`
}

type ProductDetails struct {
	Title string `json:"title"`
	Size  int    `json:"size"`
	Color string `json:"color"`
}

type CategoryDetails struct {
	Category_id      int    `json:"category_id"`
	Category_name    string `json:"category_name"`
	SubCategory_id   int    `json:"sub_category_id"`
	SUbCategory_name string `josn:"sub_category_name"`
}

type PackageDetails struct {
	Pakage_length      int  `json:"package_length"`
	Package_width      int  `json:"package_width"`
	Package_volume     int  `josn:"package_volume"`
	Dimension_verified bool `jons:"dimension_verified"`
}

type FulfillmentDetails struct {
	Storage_type       string `json:"storage_type"`
	Fulfillment_type   string `json:"fulfillment_type"`
	Lot_required       int    `json:"lot_required"`
	Serial_no_required int    `json:"serial_no_required"`
	Instruction        string `json:"instruction"`
}

type PriceDetails struct {
	Retail_price   int `json:"max_retail_price"`
	Selling_price  int `json:"selling_price"`
	Sourcing_price int `json:"sourcing_price"`
}

type Image struct {
	Front_view string `json:"front_view"`
	Left_view  string `json:"left_view"`
	Right_view string `json:"right_view"`
	Top_view   string `json:"top_view"`
	Back_view  string `json:"back_view"`
}

type ItemDimension struct {
	Item_length         int    `json:"item_length"`
	Item_width          int    `json:"item_width"`
	Item_height         int    `json:"item_height"`
	Item_weight         int    `json:"item_weight"`
	Item_volume         int    `json:"item_volume"`
	Item_instructions   string `json:"item_instructions"`
	Dimensions_verified bool   `json:"dimensions_verified"`
}
