package main

import (
	"fmt"
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/delivery"
	"github.com/edwardsuwirya/carCollection/repository"
	"github.com/edwardsuwirya/carCollection/useCase"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"strconv"
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
	listeningAddress := fmt.Sprintf("%s:%s", a.appConfig.GetConfigValue("host"), a.appConfig.GetConfigValue("port"))
	config.Logger.Debug(fmt.Sprintf("Server runs on %s", listeningAddress))
	to, _ := strconv.Atoi(a.appConfig.GetConfigValue("fake_api_timeout"))

	carrepo := repository.NewFakeAPIRepository(a.appConfig.GetConfigValue("fake_api_url"), to)
	carusecase := useCase.NewCarUseCase(carrepo)
	r := mux.NewRouter()
	delivery.NewRestServer(r, carusecase)
	if err := http.ListenAndServe(listeningAddress, r); err != nil {
		panic(err)
	}
}

func (a app) runFakeApi() {
	config.Logger.Debug("Run Fake API")
	to, _ := strconv.Atoi(a.appConfig.GetConfigValue("fake_api_timeout"))

	carrepo := repository.NewFakeAPIRepository(a.appConfig.GetConfigValue("fake_api_url"), to)
	carusecase := useCase.NewCarUseCase(carrepo)
	delivery.NewCliDelivery(carusecase)
}

func (a app) runTemp() {
	carrepo := repository.NewTempRepository()
	carusecase := useCase.NewCarUseCase(carrepo)
	delivery.NewCliDelivery(carusecase)
}

func newApp(configPath string) *app {
	cfg := config.NewConfig(configPath)
	return &app{
		cfg,
	}
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
