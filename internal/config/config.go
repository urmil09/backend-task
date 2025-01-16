package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Load loads the config.
func Load() (Config, error) {
	var c Config
	err := env.Parse(&c)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}
	return c, nil
}

// config is the root configuration.
type Config struct {
	HTTPPort string `env:"HTTP_PORT" envDefault:"8080"`
	Database Database
}

// Database contains configurations relevant to our database connection.
type Database struct {
	Host     string `env:"DB_HOST" envDefault:"postgres-db"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"flink"`
	Password string `env:"DB_PASSWORD" envDefault:"flink.8080"`
	Name     string `env:"DB_NAME" envDefault:"flink"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
}

// ConnectionString returns the database connection string.
func (cfg Database) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)
}
