package handlers

import (
	"bookmarks/model"
	"bookmarks/responses"
	"bookmarks/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	bookMarkService *service.BookmarkService
}

func (ms *MyServer) GetAllBookmarks(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	log.Println("Getting all bookmarks")
	myServer := MyServer{bookMarkService: &service.BookmarkService{}}

	m, err := myServer.bookMarkService.ViewBookMarks()
	if err != nil {
		responses.Response500(writer, err.Error())

		return
	}

	if len(m) <= 0 {
		responses.Response200(writer, "Закладок нет")

		return
	}

	responses.Response200(writer, m)
}

func (ms *MyServer) PostBookmark(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	log.Println("Post bookmark")
	var bookmark model.Bookmark

	err := json.NewDecoder(r.Body).Decode(&bookmark)
	if err != nil {
		log.Println("Invalid json received from client:", err)

		responses.Response400(writer, "Invalid json received from client")

		return
	}

	myServer := MyServer{bookMarkService: &service.BookmarkService{}}
	err = myServer.bookMarkService.AddBookMarks(bookmark.Name, bookmark.Text)

	if err != nil {
		log.Println("Закладка с таким названием уже есть")

		responses.Response500(writer, "Закладка с таким названием уже есть")

		return
	}

	responses.Response200(writer, "Закладка добавлена")
}

func (ms *MyServer) DeleteBookmark(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)
	log.Println("Delete bookmark")

	path := r.URL.Path
	bookMarkName := fmt.Sprintf("%s", path[11:])
	myServer := MyServer{bookMarkService: &service.BookmarkService{}}
	err := myServer.bookMarkService.DeleteBookMarks(bookMarkName)
	if err != nil {
		log.Println("Закладки с таким названием нет")
		responses.Response500(writer, "Закладки с таким названием нет")

		return
	}

	responses.Response200(writer, "Закладка удалена")
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
