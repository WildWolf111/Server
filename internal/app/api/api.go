package api

import (
	"net/http"

	"StandartWebServer/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {

	if err := api.configureLoggerFiled(); err != nil {

		return err
	}
	api.logger.Debug("starting api server at port:", api.config.BindAddr)

	api.configureRouterFiled()

	api.configureStoreField()
	if err := api.configureStoreField(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.BindAddr, api.router)

}
