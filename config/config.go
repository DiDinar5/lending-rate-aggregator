package config

type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	HTTP HTTP
}

func LoadConfig() (Config, error) {
	return Config{
		HTTP: HTTP{
			Host: "localhost",
			Port: ":8080",
		},
	}, nil
}
