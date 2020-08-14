package main

import (
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/delivery"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	AppName    = "Enigma Car Collection"
	AppTagLine = "Car Collection Sample Project"
	Version    = "0.0.1"
)

type app struct {
	appConfig *config.Config
}

func (a app) runServer() {
	a.run(delivery.NewRestServer(a.appConfig))
}

func (a app) runFakeApi() {
	a.run(delivery.NewCliDelivery(a.appConfig))
}

func (a app) runTemp() {
	a.run(delivery.NewCliDeliveryTemp(a.appConfig))
}

func newApp(configPath string) *app {
	cfg := config.NewConfig(configPath)
	cfg.Config()
	return &app{appConfig: cfg}
}

func (a app) run(delivery delivery.CarDelivery) {
	delivery.Run()
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
					newApp(c.String("config")).runFakeApi()
					return nil
				},
			},
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "Run server mode",
				Action: func(c *cli.Context) error {
					newApp(c.String("config")).runServer()
					return nil
				},
			},
			{
				Name:    "runtemp",
				Aliases: []string{"t"},
				Usage:   "Run with Temporary Slice",
				Action: func(c *cli.Context) error {
					newApp(c.String("config")).runTemp()
					return nil
				},
			},
		},
	}
	err := appConfig.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
