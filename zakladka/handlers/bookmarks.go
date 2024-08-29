package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"zakladka/model"
	"zakladka/responses"
	"zakladka/service"
)

type MyServer struct {
	bookMarkService *service.BookmarkService
}

func (ms *MyServer) GetAllBookMarks() (map[string]string, error) {
	return ms.bookMarkService.ViewBookMarks()
}

func (ms *MyServer) AddBookMark(name, text string) error {
	return ms.bookMarkService.AddBookMarks(name, text)
}

func (ms *MyServer) RemoveBookMark(name string) error {
	return ms.bookMarkService.DeleteBookMarks(name)
}

func GetAllBookmarks(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	log.Println("Getting all bookmarks")
	myServer := MyServer{bookMarkService: &service.BookmarkService{}}

	m, err := myServer.GetAllBookMarks()
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

func PostBookmark(writer http.ResponseWriter, r *http.Request) {
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
	err = myServer.AddBookMark(bookmark.Name, bookmark.Text)

	if err != nil {
		log.Println("Закладка с таким названием уже есть")

		responses.Response500(writer, "Закладка с таким названием уже есть")

		return
	}

	responses.Response200(writer, "Закладка добавлена")
}

func DeleteBookmark(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)
	log.Println("Delete bookmark")

	path := r.URL.Path
	bookMarkName := fmt.Sprintf("%s", path[11:])
	myServer := MyServer{bookMarkService: &service.BookmarkService{}}
	err := myServer.RemoveBookMark(bookMarkName)
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
