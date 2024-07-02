package utils

import (
	"Lesson4/handlers"
	"github.com/gorilla/mux"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookByID).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookByID).Methods("DELETE")
}

func BuildManyBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
