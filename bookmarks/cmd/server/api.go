package server

import (
	"bookmarks/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = "8080"

var myServer = handlers.MyServer{}

func StartServer() {
	log.Println("Starting BookMarks server...")
	router := mux.NewRouter()

	router.HandleFunc("/bookmarks", myServer.GetAllBookmarks).Methods(http.MethodGet)
	router.HandleFunc("/bookmarks", myServer.PostBookmark).Methods(http.MethodPost)
	router.HandleFunc("/bookmarks/{id}", myServer.DeleteBookmark).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
