package storage

type Config struct {
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURI: "host=localhost port=5433 user=postgres password=postgres dbname=restapi sslmode=disable",
	}
}
