package config

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func (c *AppConfig) GetConfig(configPath string) *AppConfig {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Error().
			Msg("Error reading YAML file.")
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Error().
			Msg("Error ummarshalling YAML file.")
		os.Exit(1)
	}

	return c
}
