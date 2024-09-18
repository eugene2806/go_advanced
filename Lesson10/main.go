package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go_advanced/Lesson10/models"
	"go_advanced/Lesson10/routers"
	"go_advanced/Lesson10/storage"
	"log"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=eugene password=postgres dbname=restapi sslmode=disable")

	if err != nil {
		log.Println("error while accessing database connection", err)
	}
	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Article{}) // орм сама сгенерит все запросы, миграции и их применит

	r := routers.SetupRouter()
	r.Run(":8080")
}
