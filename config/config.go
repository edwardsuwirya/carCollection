package config

import (
	"github.com/spf13/viper"
)

type config struct {
	configFilePath string
}

func NewConfig(configFilePath string) *config {
	return &config{configFilePath: configFilePath}
}
func (c *config) Config() *viper.Viper {
	viper.SetConfigFile(c.configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg := viper.GetViper()
	return cfg
}

func (c *config) GetConfigValue(key string) string {
	return c.Config().GetString(key)
}
