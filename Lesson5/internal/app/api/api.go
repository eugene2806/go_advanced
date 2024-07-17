package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Base API Server instance description

type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configLoggerField(); err != nil {
		return err
	}

	api.logger.Info("Starting API Server at port:", api.config.BindAddr)

	api.configRouterField()

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
