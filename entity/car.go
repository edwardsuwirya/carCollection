package entity

import "fmt"

type Car struct {
	CarDetail CarDetail
}

func (c *Car) ToString() string {
	return fmt.Sprintf("%d-%s-%s-%s",
		c.CarDetail.Id,
		c.CarDetail.Car,
		c.CarDetail.CarModel,
		c.CarDetail.CarColor)
}

type CarDetail struct {
	Id       int
	Car      string
	CarModel string
	CarColor string
}
