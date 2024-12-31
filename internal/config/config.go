package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}
type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true"`
	Storage    string `yaml:"storage" env:"STORAGE" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var config string
	config = os.Getenv("CONFIG")

	if config == "" {
		flags := flag.String("config", "", "config file path")
		flag.Parse()
		config = *flags
		if config == "" {
			log.Fatal("Config File Path Is Not Set")
		}
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatal("Config File Not Exist")
	}

	var cfg Config

	err := cleanenv.ReadConfig(config, &cfg)

	if err != nil {
		log.Fatal("Can Not Read Config File: %s", err.Error())
	}

	return &cfg
}
