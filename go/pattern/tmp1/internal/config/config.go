package config

import "os"

type Config struct {
	PublicHost string
	Port       string
	DBName     string
	DBAddr     string
	DBUser     string
	DBPass     string
}

func Load() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBName:     getEnv("DB_NAME", "mysql-local"),
		DBAddr:     getEnv("DB_ADDR", "172:0.0.0:3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "abcd"),
	}
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
