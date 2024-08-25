package rest

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"zakladka/handlers"
)

const port = "8080"

func StartServer() {
	log.Println("Starting BookMarks server...")
	router := mux.NewRouter()
	router.HandleFunc("/bookmarks", handlers.GetAllBookmarks).Methods(http.MethodGet)
	router.HandleFunc("/bookmarks", handlers.PostBookmark).Methods(http.MethodPost)
	router.HandleFunc("/bookmarks/{id}", handlers.DeleteBookmark).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
