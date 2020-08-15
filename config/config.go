package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var (
	Logger   logrus.Logger
	CfgViper *viper.Viper
)

type Config struct {
	configFilePath string
}

func NewConfig(configFilePath string) *Config {
	c := Config{configFilePath: configFilePath}
	c.init()
	return &c
}

func (c *Config) init() *viper.Viper {
	viper.SetConfigFile(c.configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	CfgViper = viper.GetViper()
	level := CfgViper.GetString("log_level")
	c.logger(level)
	return CfgViper
}

func (c *Config) logger(level string) {
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

func (c *Config) GetConfigValue(key string) string {
	return CfgViper.GetString(key)
}
