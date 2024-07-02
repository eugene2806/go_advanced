package main

import (
	"Lesson4/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const apiPrefix string = "/api/v1"

var (
	port                    string
	bookResourcePrefix      string = apiPrefix + "/book"
	manyBooksResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server... on port:", port)

	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)
	log.Println("Router configured successfully")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
