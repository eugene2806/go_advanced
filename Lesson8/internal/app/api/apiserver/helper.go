package apiserver

import (
	"github.com/sirupsen/logrus"
	"go_advanced/Lesson8/internal/app/api/middleware"
	"go_advanced/Lesson8/storage"
	"net/http"
)

var prefix string = "/api/v1"

func (api *API) configLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}

	api.logger.SetLevel(logLevel)

	return nil
}

func (api *API) configRouterField() {
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods(http.MethodGet)
	// Было до jwt
	//api.router.HandleFunc(prefix+"/articles/{id}", api.GetArticleById).Methods(http.MethodGet)
	// Теперь требует наличия jwt
	api.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleWare.Handler(
		http.HandlerFunc(api.GetArticleById),
	)).Methods(http.MethodGet)
	api.router.HandleFunc(prefix+"/articles/{id}", api.DeleteArticleById).Methods(http.MethodDelete)
	api.router.HandleFunc(prefix+"/articles", api.PostArticle).Methods(http.MethodPost)
	api.router.HandleFunc(prefix+"/articles/{id}", api.PutArticle).Methods(http.MethodPut)
	api.router.HandleFunc(prefix+"/user/register", api.PostUserRegister).Methods(http.MethodPost)
	// New pair for auth
	api.router.HandleFunc(prefix+"/user/auth", api.PostToAuth).Methods(http.MethodPost)
}

func (api *API) configStorageField() error {
	storageConf := storage.New(api.config.Storage)
	if err := storageConf.Open(); err != nil {

		return err
	}

	api.storage = storageConf

	return nil
}
