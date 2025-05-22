package config

import "os"

type Config struct {
	PublicHost string
	Port       string
	DBName     string
	DBUser     string
	DBPass     string
	DBAddr     string
}

func NewConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBName:     getEnv("DB_NAME", "mysql"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "abc"),
		DBAddr:     getEnv("DB_ADDR", "172.0.0.0:3306"),
	}
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
