package main

import (
	"flag"
	"learning_GO/RAMEN_API_base/internal/app/api"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file")

}

func main() {
	flag.Parse()

	//server inti
	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	log.Println("File config:", config.Storage)
	if err != nil {
		log.Println("can not find configs file, using default values", err)
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
