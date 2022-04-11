package utils

import (
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() {
	err := mgm.SetDefaultConfig(nil, "Products", options.Client().ApplyURI(Config("MONGOURI", "")))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connected to db")
}
