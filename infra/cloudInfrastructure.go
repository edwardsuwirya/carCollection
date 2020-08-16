package infra

import "github.com/edwardsuwirya/carCollection/entity"

type CloudInfrastucture interface {
	CreateDocument(car *entity.Car) (*entity.Car, error)
	Upload(fileInput []byte, fileName string) error
	GetAllDocument() ([]*entity.Car, error)
}
