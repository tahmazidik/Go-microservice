package config

import "github.com/tahmazidik/Go-microservice/internal/config/database"

type Config struct {
	Database *database.Config
}

func NewConfig() *Config {
	return &Config{
		Database: &database.Config{
			User:     getEnv("DB_CONFIG_USER", "root"),
			Password: getEnv("DB_CONFIG_PASSWORD", "root"),
			Host:     getEnv("DB_CONFIG_HOST", "localhost"),
			Port:     getEnvAsInt("DB_CONFIG_PORT", 3306),
			Dbname:   getEnv("DB_CONFIG_DBNAME", ""),
		},
	}
}
