package models

import (
	"github.com/kamva/mgm/v3"
)

type MainItemAssociationData struct {
	Data ItemAssociation `json:"data" bson:"data"`
}

type ItemAssociation struct {
	mgm.DefaultModel `bson:",inline"`
	Account_id       string `json:"account_id" bson:"account_id,omitempty"`
	Parent_sku_id    string `json:"parent_sku_id" bson:"parent_sku_id,omitempty"`
	Sku_id           string `json:"sku_id" bson:"sku_id,omitempty"`
	Title            string `json:"title" bson:"title,omitempty"`
	Front_view       string `json:"front_view" bson:"front_view,omitempty"`
	Quantity         int    `json:"quantity" bson:"quantity,omitempty"`
	Status           string `json:"status" bson:"status,omitempty"`
	Is_archive       bool   `json:"is_archive" bson:"is_archive,omitempty"`
}
