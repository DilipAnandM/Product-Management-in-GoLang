package models

import (
	"github.com/kamva/mgm/v3"
)

type MainAssociationData struct {
	Data Association `json:"data" bson:"data"`
}

type Association struct {
	mgm.DefaultModel `bson:",inline"`
	Account_id       string    `json:"account_id" bson:"account_id,omitempty"`
	Parent_sku_id    string    `json:"parent_skU_id" bson:"parent_sku_id,omitempty"`
	Association_type string    `json:"association_type" bson:"association_type,omitempty"`
	Title            string    `json:"title" bson:"title,omitempty"`
	Images           AssoImage `json:"images" bson:"images,omitempty"`
	Description      string    `json:"description" bson:"description,omitempty"`
	Instruction      string    `json:"instruction" bson:"instruction,omitempty"`
	Selling_price    int       `json:"selling_price" bson:"selling_price,omitempty"`
	Currency_code    string    `json:"currency_code" bson:"currency_code,omitempty"`
	Status           string    `json:"status" bson:"status,omitempty"`
	Is_archive       bool      `json:"is_archive" bson:"is_archive,omitempty"`
}

type AssoImage struct {
	Front_view    string `json:"front_view" bson:"front_view,omitempty"`
	Left_view     string `json:"left_view" bson:"left_view,omitempty"`
	Right_view    string `json:"right_view" bson:"right_view,omitempty"`
	Top_view      string `json:"top_view" bson:"top_view,omitempty"`
	Back_view     string `json:"back_view" bson:"back_view,omitempty"`
	Zoom_out_view string `json:"zoom_out_view" bson:"zoom_out_view,omitempty"`
	Zoom_in_view  string `json:"zoom_in_view" bson:"zoom_in_view,omitempty"`
}
