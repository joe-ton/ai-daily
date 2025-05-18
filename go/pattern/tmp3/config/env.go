package config

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAdress   string
	DBName     string
}

func initConfig() Config {
	return Config{}
}
