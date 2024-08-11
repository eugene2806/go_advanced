package storage

type Config struct {
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURI: "host=localhost port=5432 user=eugene password=postgres dbname=restapi sslmode=disable",
	}
}
