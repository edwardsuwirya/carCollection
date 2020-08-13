package delivery

import (
	"fmt"
	"github.com/edwardsuwirya/carCollection/entity"
	"github.com/edwardsuwirya/carCollection/useCase"
	"strings"
)

type Cli struct {
	useCase useCase.CarUseCase
}

func (c *Cli) init(uc useCase.CarUseCase) error {
	fmt.Println("Application Started")
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	c.useCase = uc

	car01 := entity.Car{
		CarDetail: entity.CarDetail{
			Car:      "Nissan",
			CarModel: "Livina",
			CarColor: "Dark Blue",
		},
	}
	car02 := entity.Car{
		CarDetail: entity.CarDetail{
			Car:      "Honda",
			CarModel: "Brio",
			CarColor: "Black",
		},
	}
	c.RegisterCar(&car01)
	c.RegisterCar(&car02)
	c.PrintAllCarCollection()

	return nil
}

func (c *Cli) RegisterCar(car *entity.Car) {
	newCar, err := c.useCase.RegisterCar(car)
	if err != nil {
		fmt.Println(err)
	}
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

func NewCliDelivery(useCase useCase.CarUseCase) CarDelivery {
	cli := new(Cli)
	err := cli.init(useCase)
	if err != nil {
		panic("Application failed")
	}
	return cli
}
