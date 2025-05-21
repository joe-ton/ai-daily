package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBName     string
	DBAddr     string
}

func (c Config) Load() (Config, error) {
	instance := Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "abc"),
		DBName:     getEnv("DB_NAME", "mysql"),
		DBAddr:     getEnv("DB_ADDR", "172.0.0.0:3306"),
	}
	if c.DBPass == "" {
		return Config{}, errors.New("Invalid DBPass")
	}
	return instance, nil
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
