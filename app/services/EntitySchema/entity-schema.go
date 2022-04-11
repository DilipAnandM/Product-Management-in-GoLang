package EntitySchema

var CreateSchema = []byte(`{
	"$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Data",
	"type": "object",
	"data":{
		"type": "object",
		"properties":{
			"account_id": {
				"type": "string"
			},
			"client_code":{
				"type":"string"
			},
			"product_identifiers": {
				"type": "object",
				"properties": {
					"sku_id": {
						"type": "string",
						"min": 2,
						"max": 26,
					},
					"product_id_type": {
						"type": "string"
					},
					"product_id_value": {
						"type": "string"
					}
				},
				"required": ["sku_id"]
			},
			"product_details":{
				"type":"object",
				"properties":{
					"title":{
						"type":"string",
						"min": 1,
						"max": 250,
					},
					"description":{
						"type":"string",
						"min":1,
						"max":2000
					},
					"size":{
						"type":"integer"
					},
					"color":{
						"type":"string"
					},
					"condition_type":{
						"type":"string",
						"oneof":"New UsedLikeNew UsedVeryGood UsedGood UsedAcceptable CollectibleLikeNew CollectibleVeryGood CollectibleGood CollectibleAcceptable Club"
					}
				},
				"required":["title", "description"]
			},
			"category_details":{
				"type":"object",
				"properties":{
					"category_id":{
						"type":"integer"
					},
					"category_name":{
						"type":"string"
					},
					"sub_category_id":{
						"type":"integer"
					},
					"sub_category_name":{
						"type":"string"
					}
				},
				"required": ["category_id", "sub_category_id"]
			},
			"package_details":{
				"type":"object",
				"properties":{
					"package_length":{
						"type":"integer"
					},
					"package_width":{
						"type":"integer"
					},
					"package_volume":{
						"type":"integer"
					},
					"dimension_verified":{
						"type":"bool"
					}
				}
			},
			"fulfillment_details":{
				"type":"object",
				"properties":{
					"storage_type":{
						"type":"string",
						"oneof":"Normal HighValued Freeze COld Hazard Hanger Hold Damage"
					},
					"fulfillment_type":{
						"type":"string",
						"oneof":"MFN WFN VSKU"
					},
					"lot_requried":{
						"type":"integer",
						"oneof":"0 1"
					},
					"serial_no_requried":{
						"type":"integer",
						"oneof":"0 1"
					},
					"instruction":{
						"type":"string",
						"min":1,
						"max":2000
					}
				},
				"required":["storage_type", fulfillment_type]
			},
			"price_details":{
				"type":"object",
				"properties":{
					"retail_price":{
						"type":"integer",
						"min":0
					},
					"selling_price":{
						"type":"integer",
						"min":0
					},
					"sourcing_price":{
						"type":"integer",
						"min":0
					}
				}
				"required": ["max_retail_price","sourcing_price"]
			},
			"image":{
				"type":"obejct",
				"properties":{
					"front_view":{
						"type": "string"
					},
					"left_view": {
						"type": "string"
					},
					"right_view": {
						"type": "string"
					},
					"top_view": {
						"type": "string"
					},
					"back_view": {
						"type": "string"
					}
				}
			},
			"item_dimension":{
				"type":"object",
				"properties":{
					"item_length":{
						"type":"integer",
						"min": 0.01
					},
					"item_width": {
						"type": "number",
						"min": 0.01
					},
					"item_height": {
						"type": "number",
						"min": 0.01
					},
					"item_weight": {
						"type": "number",
						"min": 0.01
					},
					"item_volume": {
						"type": "number"
					}
				}
			}
		},
		"required":["data"]
	}	
}`)
