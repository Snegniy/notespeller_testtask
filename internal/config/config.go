package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	DebugMode  string `env:"SERVER_DEBUG_MODE" env-description:"Debug mode logger" env-default:"yes"`
	ServerPort string `env:"SERVER_PORT" env-default:"8000"`
	Postgres   Postgres
	Names      Names
}

type Postgres struct {
	Username   string `env:"POSTGRES_USER" env-default:"postgres"`
	Password   string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Host       string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port       string `env:"POSTGRES_PORT" env-default:"5432"`
	ConnString string
}

type Names struct {
	App string `env:"APP_NAME" env-default:"note-speller"`
	DB  string `env:"DB_NAME" env-default:"note-speller-db"`
}

var path = ".env"

func NewConfig() Config {
	log.Println("\t\tRead application configuration...")
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		help, _ := cleanenv.GetDescription(&cfg, nil)
		log.Println(help)
		log.Fatalf("%s", err)
	}

	cfg.Postgres.ConnString = fmt.Sprintf("postgres://%s:%s@%s:%s/notes?sslmode=disable", cfg.Postgres.Username, cfg.Postgres.Password, cfg.Names.DB, cfg.Postgres.Port)

	log.Println("\t\tGet configuration - OK!")

	return cfg
}
