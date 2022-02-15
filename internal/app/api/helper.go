package api

import (
	"StandartWebServer/store"

	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
}

func (a *API) configureStoreField() error {
	store := store.New(a.config.Store)
	if err := store.Open(); err != nil {

		return err
	}
	a.store = store
	return nil
}
