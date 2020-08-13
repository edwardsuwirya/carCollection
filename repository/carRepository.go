package repository

import "github.com/edwardsuwirya/carCollection/entity"

type CarRepository interface {
	FindAll() ([]*entity.Car, error)
	Find(id int) (*entity.Car, error)
	Create(car *entity.Car) (*entity.Car, error)
}
