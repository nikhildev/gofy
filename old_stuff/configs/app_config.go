package configs

type AppConfig struct {
	ApiPort     int    `yaml:"port"`
	Environment string `yaml:"environment"`
}
