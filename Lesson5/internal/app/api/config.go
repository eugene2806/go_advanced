package api

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	//Logger level
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
	}
}

func ConfigEnvFile(config *Config, path *string) *Config {
	config = NewConfig()

	if err := godotenv.Load(*path); err != nil {
		log.Println("Error loading .env file! Using default values:", err)

		return config
	}

	config.BindAddr = os.Getenv("BIND_ADDR")
	config.LoggerLevel = os.Getenv("LOGGER_LEVEL")
	log.Println("Successfully loaded .env file!")

	return config
}

func ConfigTomlFile(config *Config, path *string) *Config {
	config = NewConfig()

	if _, err := toml.DecodeFile(*path, config); err != nil {
		log.Println("Error loading .toml file! Using default values:", err)

		return config
	}

	log.Println("Successfully loaded .toml file!")

	return config
}
