package store

type Config struct{
	databaseURL string `toml:"database_url"`
}

func NewConfig() *Config{
	return &Config{
		databaseURL: "host=localhost port=5322 user=admin dbname=kuznecov_messenger password=admin sslmode=disable",
	}
}