package config

type Config struct {
	Port   uint16
	DbName string
}

func InitConfig() *Config {
	return &Config{
		Port:   3000,
		DbName: "sql",
	}
}
