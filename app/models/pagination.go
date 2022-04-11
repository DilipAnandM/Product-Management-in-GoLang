package models

type Query struct {
	PageNo        int
	Limit         int
	Status        string
	Sku_id        string
	Account_id    string
	Parent_sku_id string
}
