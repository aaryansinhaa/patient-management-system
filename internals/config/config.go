package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServerConfig struct {
	Host string `yaml:"host"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Env              string           `yaml:"env"`
	Description      string           `yaml:"description"`
	HTTPServerConfig HTTPServerConfig `yaml:"http_server"`
	DatabaseConfig   DatabaseConfig   `yaml:"database"`
}

func MustLoadConfig() *Config {
	var config Config
	configPath := "/home/aaryankumarsinha/internship/tmp/config/config.yaml"

	if configPath == "" {
		flags := flag.String("config", "config/prod.yaml", "Path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Configuration file path is not set. Use the --config flag.")
		}

	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist at " + configPath)
	}
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return &config
}
