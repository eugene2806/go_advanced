package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_advanced/Lesson7/internal/app/models"
	"net/http"
	"strconv"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("content-type", "application/json")
}

// Responses------------------

func response200(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 200,
		Message:    message,
		IsError:    false,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(msg)
}

func response400(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 400,
		Message:    message,
		IsError:    true,
	}

	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(msg)
}

func response500(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 500,
		Message:    message,
		IsError:    true,
	}

	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(msg)
}

//Handlers------------------

func (api *API) GetAllArticles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	api.logger.Info("Get All Articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()

	if err != nil {
		api.logger.Info("Error while Articles.SelectAll :", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/article")

	var article models.Article
	err := json.NewDecoder(request.Body).Decode(&article)

	if err != nil {
		api.logger.Info("Invalid json received from client:", err)

		response400(writer, "Provided json is invalid")

		return
	}

	a, err := api.storage.Article().Create(&article)

	if err != nil {
		api.logger.Info("Trouble while creating new article:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article By ID /api/v1/articles/{id}")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing [id] param :", err)

		response400(writer, "Don't use id as casting to int value")

		return
	}

	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	if !ok {
		api.logger.Info("Cant find article by id in database:", id)
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that id not found in database",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	api.logger.Info("Get Article By ID /api/v1/articles/{id}")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing [id] param :", err)

		response400(writer, "Don't use id as casting to int value")

		return
	}

	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	response200(writer, fmt.Sprintf("Article width ID %d deleted successfully", id))
}

func (api *API) PostUserRegister(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register POST /api/v1/user/register")

	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		api.logger.Info("Invalid json received from client:", err)

		response400(writer, "Provided json is invalid")

		return
	}

	// Ищем пользователя в бд
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (user) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	// Если такой пользователь уже есть то не делаем регистрацию
	if ok {
		api.logger.Info("User already registered with username:", user.Login)

		response400(writer, "User with that login already exists")

		return
	}
	// Теперь пытаемся добавить в бд
	userAdd, err := api.storage.User().Create(&user)

	if err != nil {
		api.logger.Info("Troubles while accessing database table (user) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	response200(writer, fmt.Sprintf("User registered with username: %s", userAdd.Login))
}
