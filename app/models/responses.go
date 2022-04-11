package models

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
	// Errors  RequriedFields `json:"errors,omitempty"`
	Errors         []string `json:"errors" bson:"errors"`
	Requiredfields []string `json:"required_fields,omitempty"`
}
