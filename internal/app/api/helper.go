package api

import (
	"StandartWebServer/store"
	"net/http"

	"github.com/sirupsen/logrus"
)

//пытаемся откофигугрировать наш api instance
func (a *API) configureLoggerFiled() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterFiled() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})
}
func (a *API) configureStoreField() error {
	store := store.New(a.config.Store)
	if err := store.Open(); err != nil {
		return err
	}
	a.store = store
	return nil
}
