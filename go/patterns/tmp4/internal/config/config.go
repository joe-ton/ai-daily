package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBName     string
	DBUsername string
	DBPassword string
	DBAddress  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "https://localhost"),
		Port:       getEnv("Port", "8080"),
		DBName:     getEnv("DB_NAME", "mysql"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", "adbc"),
		DBAddress:  getEnv("DB_ADDRESS", "172.0.0.1:3306"),
	}
}

func Load() (Config, error) {
	c := Config{
		PublicHost: getEnv("PUBLIC_HOST", "https://localhost"),
		Port:       getEnv("Port", "8080"),
		DBName:     getEnv("DB_NAME", "mysql"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", "adbc"),
		DBAddress:  getEnv("DB_ADDRESS", "172.0.0.1:3306"),
	}

	if c.DBPassword == "" {
		return Config{}, errors.New("Invalid")
	}
	return c, nil
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
