package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var (
	Logger   logrus.Logger
	cfgViper *viper.Viper
)

type config struct {
	configFilePath string
}

func NewConfig(configFilePath string) *config {
	return &config{configFilePath: configFilePath}
}
func (c *config) Config() *viper.Viper {
	if _, err := os.Stat(c.configFilePath); err != nil {
		Logger.WithField("Server", "Status").Fatal("Config file path is not found")
		panic(err)
	}
	viper.SetConfigFile(c.configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cfgViper = viper.GetViper()
	c.Logger()
	return cfgViper
}

func (c *config) Logger() {
	level := cfgViper.GetString("log_level")
	logFormat := new(logrus.JSONFormatter)
	var logLevel, err = logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	Logger = logrus.Logger{
		Out:       os.Stderr,
		Formatter: logFormat,
		Level:     logLevel,
	}
}

func (c *config) GetConfigValue(key string) string {
	return c.Config().GetString(key)
}
