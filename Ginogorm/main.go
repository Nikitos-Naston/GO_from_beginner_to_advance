package main

import (
	"Ginogorm/models"
	"Ginogorm/routers"
	"Ginogorm/storage"
	"log"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host=... user = ... password = ... dbname = ...")
	if err != nil {
		log.Println("error while accessing", err)
	}

	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Arcticle{})

	r := routers.SetupRouter()

	r.Run()

}
