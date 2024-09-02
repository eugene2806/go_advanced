package main

import (
	"flag"
	"go_advanced/Lesson8/internal/app/api/apiserver"
	"log"
)

var (
	config *apiserver.Config
)

func init() {
	format := flag.String("format", ".toml", "Configuration file format .toml")
	path := flag.String("path", "configs/api.toml", "The path to the configuration file")

	flag.Parse()

	switch *format {
	case ".env":
		config = apiserver.ConfigEnvFile(config, path)

	case ".toml":
		config = apiserver.ConfigTomlFile(config, path)

	default:
		log.Println("Successfully loaded default configuration!")
		config = apiserver.NewConfig()
	}
}

func main() {
	log.Println("It's running")

	server := apiserver.New(config)

	log.Fatal(server.Start())
}
