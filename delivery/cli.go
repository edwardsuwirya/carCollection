package delivery

import (
	"fmt"
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/entity"
	"github.com/edwardsuwirya/carCollection/repository"
	"github.com/edwardsuwirya/carCollection/useCase"
	"io/ioutil"
	"strconv"
	"strings"
)

type Cli struct {
	useCase useCase.CarUseCase
}

func (c *Cli) RegisterCar(car *entity.Car) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	newCar, _ := c.useCase.RegisterCar(car)
	fmt.Printf("Success register : %v\n", *newCar)
}
func (c *Cli) PrintAllCarCollection() {
	coll, err := c.useCase.GetCarCollection()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", strings.Repeat("=", 55))
	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "No.", "ID", "Car", "Model", "Colour")
	fmt.Printf("%s\n", strings.Repeat("-", 55))
	for idx, c := range coll {
		carResult := (*c).ToString()
		car := strings.Split(carResult, "-")
		fmt.Printf("%-5d %-10s %-10s %-10s %-10s\n", idx+1, car[0], car[1], car[2], car[3])
		fmt.Printf("%s\n", strings.Repeat("-", 55))
	}
}

func (c *Cli) Upload() {
	f, _ := ioutil.ReadFile("/Users/edwardsuwirya/Pictures/GO_BUILD.png")
	c.useCase.Upload(f)
}

func NewCliDelivery(c *config.Config) CarDelivery {
	c.Logger.Debug("Run Fake API")
	to, _ := strconv.Atoi(c.GetConfigValue("fake_api_timeout"))
	carrepo := repository.NewFakeAPIRepository(c.GetConfigValue("fake_api_url"), to)
	carusecase := useCase.NewCarUseCase(carrepo)
	return &Cli{
		useCase: carusecase,
	}
}

func NewCliDeliveryTemp(c *config.Config) CarDelivery {
	c.Logger.Debug("Run CLI Temp")
	carrepo := repository.NewTempRepository()
	carusecase := useCase.NewCarUseCase(carrepo)
	cli := &Cli{
		useCase: carusecase,
	}
	car01 := entity.Car{CarDetail: entity.CarDetail{
		Car:          "Honda",
		CarModel:     "Brio",
		CarColor:     "",
		CarModelYear: 1900,
		CarVin:       "",
		Price:        "",
		Availability: false,
	}}
	cli.RegisterCar(&car01)
	cli.Upload()
	return cli
}

func NewCliDeliveryCloud(c *config.Config) CarDelivery {
	c.Logger.Debug("Run CLI Cloud")
	carrepo := repository.NewCloudRepository()
	carusecase := useCase.NewCarUseCase(carrepo)
	cli := &Cli{
		useCase: carusecase,
	}
	car01 := entity.Car{CarDetail: entity.CarDetail{
		Car:          "Toyota",
		CarModel:     "Kijang",
		CarColor:     "Grey",
		CarModelYear: 1900,
		CarVin:       "",
		Price:        "",
		Availability: false,
	}}
	cli.RegisterCar(&car01)
	return cli
}

func (c *Cli) Run() {
	c.PrintAllCarCollection()
}
