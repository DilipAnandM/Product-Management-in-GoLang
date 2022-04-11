package ErrorCodes

func ErrorCodes() map[string]map[string]interface{} {

	ErrorCodes := map[string]map[string]interface{}{
		"FAILED_TO_CREATE": {
			"error_code":  "VDI_SLSA_PRODUCTS_0001",
			"description": "DB error or required fields might be missing",
		},
		"FAILED_TO_UPDATE": {
			"error_code":  "VDI_SLSA_PRODUCTS_0002",
			"description": "Record might not exist to perform update operation",
		},
		"FAILED_TO_REMOVE_ASSOCIATION": {
			"error_code":  "VDI_SLSA_PRODUCTS_0003",
			"description": "Invalid association or can't remove",
		},
		"FAILED_TO_CREATE_ASSOCIATION": {
			"error_code":  "VDI_SLSA_PRODUCTS_0004",
			"description": "Duplicate Parent SKU. It should be unique",
		},
		"FAILED_TO_FETCH_DATA": {
			"error_code":  "VDI_SLSA_PRODUCTS_0005",
			"description": "Failed to fetch data or data not found.",
		},
		"FAILED_TO_UPDATE_ITEM_DIMENSIONS": {
			"error_code":  "VDI_SLSA_PRODUCTS_0006",
			"description": "There is a problem in updating item dimensions.",
		},
		"FAILED_TO_UPDATE_DIMENSIONS": {
			"error_code":  "VDI_SLSA_PRODUCTS_0007",
			"description": "There is a problem in updating dimensions.",
		},
		"NODE_INTERNAL": {
			"error_code":  "NODE_INT_000",
			"description": "Internal Error, Unable to process the request",
		},
	}
	return ErrorCodes
}

var Error_codes map[string]map[string]interface{} = ErrorCodes()
