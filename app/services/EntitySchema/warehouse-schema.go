package EntitySchema

var SchemaData = []byte(`{
    "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "data":{
			"type":"object",
			"properties":{
				"account_id":{
					"type":"string",
					"minimum": 4
				},
				"warehouse_id":{
					"type":"string"
				},
				"sku_id":{
					"type":"string"
				},
				"Status":{
					"type":"string",
					"oneof":"active inactive"
				}
			}
		}
    }
  }`)
