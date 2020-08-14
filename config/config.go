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
	c.Config()
	return &c
}
func (c *Config) Config() *viper.Viper {
	if _, err := os.Stat(c.configFilePath); err != nil {
		Logger.WithField("Application", "Status").Fatal("Config file path is not found")
		panic(err)
	}
	viper.SetConfigFile(c.configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	CfgViper = viper.GetViper()
	c.Logger()
	return CfgViper
}

func (c *Config) Logger() {
	level := CfgViper.GetString("log_level")
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
	return c.Config().GetString(key)
}
