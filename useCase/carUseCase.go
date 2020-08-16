package useCase

import (
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/entity"
	"github.com/edwardsuwirya/carCollection/repository"
)

type CarUseCase interface {
	RegisterCar(car *entity.Car) (*entity.Car, error)
	GetCarCollection() ([]*entity.Car, error)
	Upload(file []byte) error
}

type CarUseCaseImplementation struct {
	repo repository.CarRepository
}

func NewCarUseCase(repo repository.CarRepository) CarUseCase {
	return &CarUseCaseImplementation{repo}
}
func (c *CarUseCaseImplementation) RegisterCar(car *entity.Car) (*entity.Car, error) {
	if err := car.Validate(); err != nil {
		config.AppConfig.Logger.WithField("Car", "Registration").Error(err)
		return nil, err
	}
	car, err := c.repo.Create(car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (c *CarUseCaseImplementation) GetCarCollection() ([]*entity.Car, error) {
	coll, err := c.repo.FindAll()
	return coll, err
}

func (c *CarUseCaseImplementation) Upload(file []byte) error {
	return c.repo.Upload(file, "File_001.png")
}
