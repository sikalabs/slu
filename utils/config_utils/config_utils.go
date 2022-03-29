package config_utils

import (
	"github.com/sikalabs/slu/config"
)

func GetCurrentDigitalOceanAccount() *config.SluSecretsDigitalOcean {
	s := config.ReadSecrets()
	c := config.ReadState()

	for _, do := range s.DigitalOcean {
		if do.Alias == c.DigitalOcean.CurrentContext {
			return &do
		}
	}
	return nil
}

func GetDigitalOceanAccountByAlias(alias string) *config.SluSecretsDigitalOcean {
	s := config.ReadSecrets()

	for _, do := range s.DigitalOcean {
		if do.Alias == alias {
			return &do
		}
	}
	return nil
}
