package api

import (
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type API struct {
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

	if err := api.configureLoggerFiled(); err != nil {

		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	api.configureRouterFiled()
	return nil
}
