package config

type AppConfig struct {
	Apps    []int `yaml:"apps"`
	Discord struct {
		Enabled bool   `yaml:"enabled"`
		Webhook string `yaml:"webhook"`
	} `yaml:"discord"`
	Interval int `yaml:"interval"`
}
