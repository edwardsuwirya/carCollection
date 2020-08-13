package repository

import (
	"encoding/json"
	"fmt"
	"github.com/edwardsuwirya/carCollection/entity"
	"math/rand"
	"net/http"
)

type FakeApiRepository struct {
	url string
}

func NewFakeAPIRepository(url string) CarRepository {
	return &FakeApiRepository{
		url,
	}
}

func (t *FakeApiRepository) FindAll() ([]*entity.Car, error) {
	var defaultClient = &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cars", t.url), nil)
	resp, err := (*defaultClient).Do(req)
	if err != nil {
		return nil, err
	}
	//data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var carResponse entity.Cars
	err = json.NewDecoder(resp.Body).Decode(&carResponse)
	if err != nil {
		return nil, err
	}
	var result []*entity.Car
	for _, cm := range carResponse.Cars {
		result = append(result, &entity.Car{CarDetail: entity.CarDetail{
			Id:       cm.Id,
			Car:      cm.Car,
			CarModel: cm.CarModel,
			CarColor: cm.CarColor,
		}})
	}
	return result, nil
}

func (t *FakeApiRepository) Find(id int) (*entity.Car, error) {
	panic("implement me")
}

func (t *FakeApiRepository) Create(car *entity.Car) (*entity.Car, error) {
	car.CarDetail.Id = rand.Intn(100)
	return car, nil
}
