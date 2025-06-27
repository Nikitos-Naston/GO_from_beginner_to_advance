package api

import "learning_GO/RAMEN_API_base/storage"

type Config struct {
	BinAddr string `toml:"bind_addr"`

	LoggerLevel string `toml:"logger_level"`

	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:     "8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
