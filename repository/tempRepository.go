package repository

import (
	"github.com/edwardsuwirya/carCollection/entity"
	"math/rand"
)

type TempRepository struct {
	repo []*entity.Car
}

func NewTempRepository() CarRepository {
	return &TempRepository{}
}

func (t *TempRepository) FindAll() ([]*entity.Car, error) {
	return t.repo, nil
}

func (t *TempRepository) Find(id int) (*entity.Car, error) {
	for _, c := range t.repo {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, nil
}

func (t *TempRepository) Create(car *entity.Car) (*entity.Car, error) {
	car.CarDetail.Id = rand.Intn(100)
	t.repo = append(t.repo, car)
	return car, nil
}
