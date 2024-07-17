package api

import (
	"github.com/sirupsen/logrus"
	"go_advanced/Lesson6/storage"
	"net/http"
)

func (api *API) configLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}

	api.logger.SetLevel(logLevel)

	return nil
}

func (api *API) configRouterField() {
	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is rest API"))
	})
}

func (api *API) configStorageField() error {
	storageConf := storage.New(api.config.Storage)
	if err := storageConf.Open(); err != nil {

		return err
	}

	api.storage = storageConf

	return nil
}
