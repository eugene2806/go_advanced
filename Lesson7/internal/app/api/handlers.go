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

func (api *API) GetAllArticles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	api.logger.Info("Get All Articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()

	if err != nil {
		api.logger.Info("Error while Articles.SelectAll :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)

		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/article")

	var article models.Article
	err := json.NewDecoder(request.Body).Decode(&article)

	if err != nil {
		api.logger.Info("Invalid json received from client:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a, err := api.storage.Article().Create(&article)

	if err != nil {
		api.logger.Info("Trouble while creating new article:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article By ID /api/v1/articles/{id}")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing [id] param :", err)

		msg := Message{
			StatusCode: 400,
			Message:    "Don't use id as casting to int value",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) width id:", err)

		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	if !ok {
		api.logger.Info("Cant find article by id in database:", id)
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that id not found in database",
			IsError:    false,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	api.logger.Info("Get Article By ID /api/v1/articles/{id}")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing [id] param :", err)

		msg := Message{
			StatusCode: 400,
			Message:    "Don't use id as casting to int value",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) width id:", err)

		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	writer.WriteHeader(200)
	msg := Message{
		StatusCode: 200,
		Message:    fmt.Sprintf("Article width ID %d deleted successfully", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)

}

func (api *API) PostUserRegister(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register POST /api/v1/user/register")

	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		api.logger.Info("Invalid json received from client:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	// Ищем пользователя в бд
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (user) width id:", err)

		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	// Если такой пользователь уже есть то не делаем регистрацию
	if ok {
		api.logger.Info("User already registered with username:", user.Login)
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login already exists ",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}
	// Теперь пытаемся добавить в бд
	userAdd, err := api.storage.User().Create(&user)

	if err != nil {
		api.logger.Info("Troubles while accessing database table (user) width id:", err)

		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User registered with username: %s", userAdd.Login),
		IsError:    false,
	}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
