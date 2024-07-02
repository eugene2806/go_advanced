package handlers

import (
	"Lesson4/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error converting id to int => Error:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Do not use ID not supported int casting"}
		json.NewEncoder(writer).Encode(msg)

		return
	}

	book, ok := models.FindBookByID(id)
	log.Println("Get book with id:", id)

	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book not found"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Create new book...")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)

	if err != nil {
		msg := models.Message{Message: "Provided JSON file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	var newBookID int = len(models.DB) + 1
	book.ID = newBookID
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Update book...")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error converting id to int => Error:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Do not use ID not supported int casting"}
		json.NewEncoder(writer).Encode(msg)

		return
	}

	oldBook, ok := models.FindBookByID(id)
	var newBook models.Book

	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book not found"}
		json.NewEncoder(writer).Encode(msg)

		return
	}

	err = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "Provided JSON file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	for i, b := range models.DB {
		if b.ID == oldBook.ID {
			newBook.ID = id
			models.DB[i] = newBook
			break
		}
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(newBook)

	log.Println("information about the book has been updated")
}

func DeleteBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Delete book...")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error converting id to int => Error:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Do not use ID not supported int casting"}
		json.NewEncoder(writer).Encode(msg)

		return
	}

	book, ok := models.FindBookByID(id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book not found"}
		json.NewEncoder(writer).Encode(msg)

		return
	}

	for i, b := range models.DB {
		if b.ID == book.ID {
			models.DB = append(models.DB[:i], models.DB[i+1:]...)
			break
		}
	}

	writer.WriteHeader(200)
	msg := models.Message{Message: "Book deleted"}
	json.NewEncoder(writer).Encode(msg)
}
