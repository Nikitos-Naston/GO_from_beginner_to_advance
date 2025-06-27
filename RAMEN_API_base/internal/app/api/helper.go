package api

import (
	"learning_GO/RAMEN_API_base/storage"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is REST api!"))
	})
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
