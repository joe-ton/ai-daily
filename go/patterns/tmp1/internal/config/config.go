package config

import "os"

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string
}

func InitConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "https://localhost"),
	}
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
