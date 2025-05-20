package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBAddr     string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddr: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_Name", "ecom"),
	}
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
