package repository

import (
	"github.com/edwardsuwirya/carCollection/entity"
	"github.com/edwardsuwirya/carCollection/infra"
)

type CloudRepository struct {
	cloud infra.CloudInfrastucture
}

func NewCloudRepository() CarRepository {
	return &CloudRepository{
		cloud: infra.NewFirebaseInfrastructure(),
	}
}

func (t *CloudRepository) FindAll() ([]*entity.Car, error) {
	return t.cloud.GetAllDocument()
}

func (t *CloudRepository) Find(id int) (*entity.Car, error) {
	panic("Implement Me")
}

func (t *CloudRepository) Create(car *entity.Car) (*entity.Car, error) {
	c, err := t.cloud.CreateDocument(car)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (t *CloudRepository) Upload(file []byte, fileName string) error {
	return t.cloud.Upload(file, fileName)
}
