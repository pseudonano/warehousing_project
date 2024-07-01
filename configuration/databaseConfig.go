package configuration

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetConfig creates a new database configuration from environment variables
func GetConfig() *Config {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432" // default port for PostgreSQL
	}

	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}
