package main

import (
	"github.com/edwardsuwirya/carCollection/delivery"
	"github.com/edwardsuwirya/carCollection/repository"
	"github.com/edwardsuwirya/carCollection/useCase"
)

type app struct {
}

func (a app) run() {
	carrepo := repository.NewTempRepository()
	carusecase := useCase.NewCarUseCase(carrepo)
	delivery.NewCliDelivery(carusecase)
}
func newApp() *app {
	return &app{}
}

func main() {
	newApp().run()
}
