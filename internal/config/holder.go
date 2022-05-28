package config

import (
	"os"

	"github.com/ghodss/yaml"
)

var (
	Wombot *Config
)

func LoadConfig() error {
	f, err := os.ReadFile("./config.yml")

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, &Wombot)

	if err != nil {
		return err
	}

	return nil
}
