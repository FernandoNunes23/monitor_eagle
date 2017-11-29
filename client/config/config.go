package config

type Config struct {
	Hostname string
}

func InitConfig() Config {
	config := Config{}
	config.Hostname =  "localhost"

	return config
}