package api

import (
	"learning_GO/RAMEN_API_base/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/arcticles", a.GetAllArcticles).Methods("GET")
	a.router.HandleFunc(prefix+"/arcticles/{id}", a.GetArcticalById).Methods("GET")
	a.router.HandleFunc(prefix+"/arcticles/{id}", a.DeleteArcticalById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/arcticles", a.PostArcticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/registr", a.PostUserRegistr).Methods("POST")
}

func (a *API) configurStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}

func (c *Config) ConfigurateWithEnvFile(conf string) error {
	err := godotenv.Load(conf)
	if err != nil {
		log.Println("Cound not find .env file:", err)
		return err
	}
	port := os.Getenv("app_port")
	logger_level := os.Getenv("logger_level")
	database_uri := os.Getenv("database_uri")

	c.BinAddr = ":" + port
	c.LoggerLevel = logger_level
	c.Storage.DatabaseURI = database_uri
	return nil

}
