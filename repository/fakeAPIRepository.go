package repository

import (
	"encoding/json"
	"fmt"
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/entity"
	"math/rand"
	"net/http"
	"time"
)

type FakeApiRepository struct {
	url        string
	httpClient *http.Client
}

func NewFakeAPIRepository(url string, timeout int) CarRepository {
	return &FakeApiRepository{
		url, &http.Client{
			Timeout: time.Second * time.Duration(timeout),
		},
	}
}

func (t *FakeApiRepository) FindAll() ([]*entity.Car, error) {
	urlPath := fmt.Sprintf("%s/cars", t.url)
	config.Logger.Debug(urlPath)
	req, err := http.NewRequest("GET", urlPath, nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		config.Logger.WithField("HTTP-Client", "Status").Fatal("Can not make GET /cars request")
		return nil, err
	}
	//data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var carResponse entity.Cars
	err = json.NewDecoder(resp.Body).Decode(&carResponse)
	if err != nil {
		config.Logger.WithField("HTTP-Client", "Status").Fatal("Failed to parse JSON")
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
