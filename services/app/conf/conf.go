package conf

type Config struct {
	Port string
}

var config Config

func GetConfig() Config {
	config = Config{
		Port: "8448",
	}
	return config
}
