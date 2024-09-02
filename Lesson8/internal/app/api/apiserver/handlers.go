package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"go_advanced/Lesson8/internal/app/api/middleware"
	"go_advanced/Lesson8/internal/app/models"
	"net/http"
	"strconv"
	"time"
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

func (api *API) PutArticle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Put Article PUT /api/v1/articles/{id}")

	var article models.Article
	err := json.NewDecoder(request.Body).Decode(&article)

	if err != nil {
		api.logger.Info("Invalid json received from client:", err)

		response400(writer, "Provided json is invalid")

		return
	}

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing [id] param :", err)

		response400(writer, "Don't use id as casting to int value")

		return
	}

	articleId, ok, err := api.storage.Article().UpdateArticle(&article, id)
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

	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	article.ID = articleId
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)

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

	// 1. Ищем пользователя в бд
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (user) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	// 2. Если такой пользователь уже есть то не делаем регистрацию
	if ok {
		api.logger.Info("User already registered with username:", user.Login)

		response400(writer, "User with that login already exists")

		return
	}
	// 3. Теперь пытаемся добавить в бд
	userAdd, err := api.storage.User().Create(&user)

	if err != nil {
		api.logger.Info("Troubles while accessing database table (user) width id:", err)

		response500(writer, "We have some troubles to accessing database")

		return
	}

	response200(writer, fmt.Sprintf("User registered with username: %s", userAdd.Login))
}

func (api *API) PostToAuth(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post To Auth POST /api/v1/auth")
	var userFromJson models.User
	err := json.NewDecoder(request.Body).Decode(&userFromJson)
	// 1. Обрабатываем JSON если он вовсе не JSON или с ним какие либо проблемы
	if err != nil {
		api.logger.Info("Invalid json received from client:", err)
		response400(writer, "Provided json is invalid")

		return
	}
	// 2. Необходимо попытаться обнаружить пользователя с таким логином в бд
	userInDb, ok, err := api.storage.User().FindByLogin(userFromJson.Login)
	// Проблема доступа к бд
	if err != nil {
		api.logger.Info("Can not make user search in database", err)
		response500(writer, "We have some troubles to accessing database")

		return
	}
	// 3. Если подключение удалось но пользователя с таким логином нет
	if !ok {
		api.logger.Info("User with that login does not exists")
		response400(writer, "User with that login does not exist database. Try registering first")

		return
	}

	// 4. Если пользователь с таким логином есть в бд - проверяем совпадает ли пароль
	if userInDb.Password != userFromJson.Password {
		api.logger.Info("Invalid password")
		response400(writer, "Your password is invalid")

		return
	}

	// 5. Теперь выбиваем токен как знак успешной операции
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)               // Дополнительные действия в (формате мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Время жизни токена 2ч
	claims["admin"] = true
	claims["user"] = userInDb.Login

	tokenString, err := token.SignedString(middleware.SecretKey)

	// Если не удалось выдать токен
	if err != nil {
		api.logger.Info("Can not claim jwt-token", err)
		response500(writer, "We have some troubles. Some again")

		return
	}
	// В случае если токен успешно выбит - выдаем его клиенту
	response200(writer, tokenString)

}
