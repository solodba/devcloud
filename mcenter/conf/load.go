package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

func LoadConfigFromToml(filePath string) error {
	c = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, c)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfigFromEnv() error {
	c = NewDefaultConfig()
	return env.Parse(c)
}
