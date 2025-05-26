package config

import "os"

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBName     string
	DBAddr     string
}

func Load() *Config {
	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "abc"),
		DBName:     getEnv("DB_NAME", "mysql-local"),
		DBAddr:     getEnv("DB_ADDR", "172.0.0.1:3306"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
