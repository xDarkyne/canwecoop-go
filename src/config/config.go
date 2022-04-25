package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB             DatabaseConfig
	Port           int    `env:"APP_PORT" env-default:"8080"`
	TimeZone       string `env:"APP_TZ" env-default:"Europe/Berlin"`
	AuthCookieName string `env:"APP_AUTH_COOKIE_NAME" env-default:"cwc-auth"`
	SteamAPIKey    string `env:"STEAM_API_KEY" env-default:""`
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Username string `env:"DB_USER" env-default:"dev"`
	Password string `env:"DB_PASS" env-default:"dev"`
	DBName   string `env:"DB_NAME" env-default:"dev"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
	SSLMode  string `env:"DB_SSL_MODE" env-default:"disable"`
}

var App AppConfig

func LoadConfig() {
	var err error

	// load env file in development
	err = godotenv.Load("../.env.local")

	if err != nil {
		fmt.Println(err)
	}

	err = cleanenv.ReadEnv(&App)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
