package main

import (
	"flag"
	"go_advanced/Lesson5/internal/app/api"
	"log"
)

var (
	config *api.Config
)

func init() {
	format := flag.String("format", ".toml", "Configuration file format .toml")
	path := flag.String("path", "configs/api.toml", "The path to the configuration file")

	flag.Parse()

	switch *format {
	case ".env":
		config = api.ConfigEnvFile(config, path)

	case ".toml":
		config = api.ConfigTomlFile(config, path)

	default:
		log.Println("Successfully loaded default configuration!")
		config = api.NewConfig()
	}
}

func main() {
	log.Println("It's running")

	server := api.New(config)

	log.Fatal(server.Start())
}
