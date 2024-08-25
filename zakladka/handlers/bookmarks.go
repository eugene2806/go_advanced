package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"zakladka/model"
	"zakladka/service"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      bool   `json:"error"`
}

func response200(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 200,
		Message:    message,
		Error:      false,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(msg)
}

func response400(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 400,
		Message:    message,
		Error:      true,
	}

	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(msg)
}

func response500(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 500,
		Message:    message,
		Error:      true,
	}

	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(msg)
}

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAllBookmarks(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	log.Println("Getting all bookmarks")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(model.DB)
}

func PostBookmark(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	log.Println("Post bookmark")
	var bookmark model.Bookmark

	err := json.NewDecoder(r.Body).Decode(&bookmark)
	if err != nil {
		log.Println("Invalid json received from client:", err)

		response400(w, "Invalid json received from client")

		return
	}
	bookmarkService := service.BookmarkService{}

	msg, err := bookmarkService.AddBookMarks(model.DB, bookmark.Name, bookmark.Text)

	if err != nil {
		log.Println(msg)

		response500(w, msg)

		return
	}

	response200(w, msg)
}

func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Delete bookmark")

	path := r.URL.Path
	bookMarkName := fmt.Sprintf("%s", path[11:])
	bookmarkService := service.BookmarkService{}
	msg, err := bookmarkService.DeleteBookMarks(model.DB, bookMarkName)
	if err != nil {
		log.Println(msg)
		response500(w, msg)

		return
	}

	response200(w, msg)
}
