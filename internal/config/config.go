package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerPort       string `env:"PORT" envDefault:"8080"`
	Target           string `env:"TARGET" envDefault:"https://api.nasa.gov/planetary/apod"`
	ApiKey           string `env:"API_KEY" envDefault:"7mrYwiNGfVBzDgLEcGvsL3j88hRd57g9iHKLAFRg"`
	PostgresURL      string `env:"POSTGRES_URL" envDefault:"postgres://user:password@postgres/db?sslmode=disable"`
	ImageDestenation string `env:"IMAGE_DEST" envDefault:"./www/html"`
	NginxURL         string `env:"NGINX_PORT" envDefault:"http://localhost:8089/"`
}

func New() (Config, error) {
	cnf := Config{}
	err := env.Parse(&cnf)

	return cnf, err
}
