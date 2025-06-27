package api

import (
	"learning_GO/RAMEN_API_base/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port", api.config.BinAddr)

	api.configureRouterField()

	if err := api.configurStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BinAddr, api.router)
}
