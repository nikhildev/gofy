package configs

type AppConfig struct {
	Port        int    `yaml:"port"`
	Environment string `yaml:"environment"`
}
