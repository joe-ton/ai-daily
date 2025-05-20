package config

import "os"

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "abc"),
		DBAddress:  getEnv("DB_ADRESS", "172.0.0.0:3306"),
		DBName:     getEnv("DB_NAME", "mysql"),
	}
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}
