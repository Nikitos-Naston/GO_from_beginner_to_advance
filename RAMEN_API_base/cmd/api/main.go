package main

import (
	"flag"
	"learning_GO/RAMEN_API_base/internal/app/api"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath     string
	configFileType string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file")
	flag.StringVar(&configFileType, "type", "toml", "type of config file(env,toml)")

}

func main() {
	flag.Parse()
	config := api.NewConfig()
	if configFileType == "toml" {

		_, err := toml.DecodeFile(configPath, config)
		log.Println("File config:", config.Storage)
		if err != nil {
			log.Println("can not find configs file, using default values", err)
		}
	} else if configFileType == "env" {
		err := config.ConfigurateWithEnvFile(configPath)
		if err != nil {
			log.Println("Error with env file, take the basic values")
		}
	} else {
		log.Println("undefied type of config file take basic values")
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
