package api

import (
	"github.com/sirupsen/logrus"
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
