package main

import (
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/delivery"
	"github.com/edwardsuwirya/carCollection/repository"
	"github.com/edwardsuwirya/carCollection/useCase"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	AppName    = "Enigma Car Collection"
	AppTagLine = "Car Collection Sample Project"
	Version    = "0.0.1"
)

type app struct {
}

func (a app) runFakeApi(configFilePath string) {
	cfg := config.NewConfig(configFilePath)
	carrepo := repository.NewFakeAPIRepository(cfg.GetConfigValue("fake_api_url"))
	carusecase := useCase.NewCarUseCase(carrepo)
	delivery.NewCliDelivery(carusecase)
}

func (a app) runTemp() {
	carrepo := repository.NewTempRepository()
	carusecase := useCase.NewCarUseCase(carrepo)
	delivery.NewCliDelivery(carusecase)
}

func newApp() *app {
	return &app{}
}

func main() {

	appConfig := &cli.App{
		Name:    AppName,
		Usage:   AppTagLine,
		Version: Version,
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "runfakeapi",
				Aliases: []string{"f"},
				Usage:   "Run with Fake API",
				Action: func(c *cli.Context) error {
					newApp().runFakeApi(c.String("config"))
					return nil
				},
			},
			{
				Name:    "runtemp",
				Aliases: []string{"t"},
				Usage:   "Run with Temporary Slice",
				Action: func(c *cli.Context) error {
					newApp().runTemp()
					return nil
				},
			},
		},
	}
	err := appConfig.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
