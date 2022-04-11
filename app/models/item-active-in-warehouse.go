package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type WarehouseMain struct {
	Data Warehouse `json:"data" bson:"data"`
}
type Warehouse struct {
	mgm.DefaultModel    `bson:",inline"`
	Account_id          string            `json:"account_id" bson:"account_id" validation:"required,omitempty"`
	Warehouse_id        string            `json:"warehouse_id" bson:"warehouse_id" validation:"required,omitempty"`
	Sku_id              string            `json:"sku_id" bson:"sku_id" validation:"required,omitempty"`
	Status              string            `json:"status" bson:"status,omitempty"`
	Last_sync_timestamp time.Time         `json:"last_sync_timestamp" bson:"last_sync_timestamp,omitempty"`
	Entity_errors       map[string]string `json:"entity_errors" bson:"entity_errors,omitempty"`
}
