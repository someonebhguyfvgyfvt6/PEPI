// Package config carga la configuración de la aplicación desde variables de entorno.
package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv          string
	Port            string
	DatabaseURL     string
	ShutdownTimeout time.Duration
}

// Load lee el archivo .env (si existe) y construye la configuración.
// En producción las variables vienen del entorno, por eso el .env es opcional.
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no se encontró .env, se usan las variables del entorno")
	}

	cfg := &Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		Port:            getEnv("PORT", "8080"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		ShutdownTimeout: 10 * time.Second,
	}

	if cfg.DatabaseURL == "" {
		return nil, errors.New("config: DATABASE_URL es obligatoria")
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
