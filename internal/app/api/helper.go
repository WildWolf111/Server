package api

import (
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
