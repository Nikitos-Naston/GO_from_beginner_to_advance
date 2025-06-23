package api

type Config struct {
	BinAddr string `toml:"bind_addr"`

	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr:     "8080",
		LoggerLevel: "debug",
	}
}
