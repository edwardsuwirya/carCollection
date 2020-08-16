package infra

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/storage"
	"fmt"
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/entity"
	guuid "github.com/google/uuid"
	"google.golang.org/api/iterator"
	"io"
)

type FirebaseInfrastucture struct {
	strclient      *storage.Client
	frsclient      *firestore.Client
	collectionName string
	ctx            context.Context
}

func NewFirebaseInfrastructure() CloudInfrastucture {
	return &FirebaseInfrastucture{
		config.AppConfig.ClientStorage,
		config.AppConfig.ClientFirestore,
		config.AppConfig.GetConfigValue("google_collection"),
		context.Background(),
	}
}

func (fb *FirebaseInfrastucture) GetAllDocument() ([]*entity.Car, error) {
	iter := fb.frsclient.Collection(fb.collectionName).Documents(fb.ctx)
	carResults := make([]*entity.Car, 0, 1)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		carData := doc.Data()

		carDetailData := entity.CarDetail{
			Id:           0,
			Car:          carData["Car"].(string),
			CarModel:     carData["CarModel"].(string),
			CarColor:     carData["CarColor"].(string),
			CarModelYear: int(carData["CarModelYear"].(int64)),
			CarVin:       carData["CarVin"].(string),
			Price:        carData["Price"].(string),
			Availability: carData["Availability"].(bool),
		}
		car := entity.Car{carDetailData}
		carResults = append(carResults, &car)
	}
	return carResults, nil
}
func (fb *FirebaseInfrastucture) CreateDocument(car *entity.Car) (*entity.Car, error) {
	id := guuid.New()
	docRef, err := fb.frsclient.Collection(fb.collectionName).Doc(fmt.Sprintf("CC-%s", id.String())).Set(fb.ctx, *car)
	if err != nil {
		config.AppConfig.Logger.Error(fmt.Sprintf("Failed create document: %v", err))
		return nil, err
	}
	config.AppConfig.Logger.Debug("Create timestamp:" + docRef.UpdateTime.String())
	return car, nil
}
func (fb *FirebaseInfrastucture) Upload(fileInput []byte, fileName string) error {
	bucket, err := fb.strclient.DefaultBucket()
	if err != nil {
		return err
	}

	object := bucket.Object(fileName)
	writer := object.NewWriter(context.Background())
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(fileInput)); err != nil {
		return err
	}

	return nil
}
