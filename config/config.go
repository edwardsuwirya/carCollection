package config

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"log"
	"os"
)

var (
	AppConfig *Config
)

type Config struct {
	configFilePath  string
	CfgViper        *viper.Viper
	Logger          logrus.Logger
	ClientStorage   *storage.Client
	ClientFirestore *firestore.Client
}

func NewConfig(configFilePath string) *Config {
	c := Config{configFilePath: configFilePath}
	c.init()
	AppConfig = &c
	return AppConfig
}

func (c *Config) init() {
	viper.SetConfigFile(c.configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	c.CfgViper = viper.GetViper()
	level := c.CfgViper.GetString("log_level")
	c.logger(level)
	c.initFirebaseStorage()
}

func (c *Config) logger(level string) {
	logFormat := new(logrus.JSONFormatter)
	var logLevel, err = logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	c.Logger = logrus.Logger{
		Out:       os.Stderr,
		Formatter: logFormat,
		Level:     logLevel,
	}
}

func (c *Config) GetConfigValue(key string) string {
	return c.CfgViper.GetString(key)
}

func (c *Config) initFirebaseStorage() {
	fbConfig := &firebase.Config{
		StorageBucket: c.CfgViper.GetString("google_storage_bucket"),
	}
	opt := option.WithCredentialsFile(c.CfgViper.GetString("google_credential_path"))
	app, err := firebase.NewApp(context.Background(), fbConfig, opt)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	clientStorage, err := app.Storage(ctx)
	clientFirestore, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	c.ClientFirestore = clientFirestore
	c.ClientStorage = clientStorage
	//bucket, err := client.DefaultBucket()
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
