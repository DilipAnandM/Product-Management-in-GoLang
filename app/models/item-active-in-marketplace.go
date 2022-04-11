package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type MainMarketplaceData struct {
	Data Marketplace `json:"data" bson:"data"`
}

type Marketplace struct {
	mgm.DefaultModel       `bson:",inline"`
	Account_id             string    `json:"account_id" bson:"account_id" validation:"required,omitempty"`
	Sku_id                 string    `json:"sku_id" bson:"sku_id" validation:"required,omitempty"`
	Channel_id             string    `json:"channel_id" bson:"channel_id" validation:"required,omitempty"`
	Store_id               string    `json:"store_id" bson:"store_id" validation:"required,omitempty"`
	Product_url            string    `json:"product_url" bson:"product_url" validation:"required,omitempty"`
	Marketplace_product_id string    `json:"marketplace_product_id" bson:"marketplace_product_id" validation:"required,omitempty"`
	Location_ids           string    `json:"location_ids" bson:"location_ids" validation:"required,omitempty"`
	Status                 string    `json:"status" bson:"status" validation:"required,omitempty"`
	Last_sync_timestamp    time.Time `json:"last_sync_timestamp" bson:"last_sync_timestamp" validation:"required,omitempty"`
}
