package repository

import (
	"github.com/edwardsuwirya/carCollection/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Native golang's way
//============================
//var cars []*entity.Car
//var repo TempRepository
//
//func init() {
//	cars = []*entity.Car{
//		&entity.Car{entity.CarDetail{Id: 1, Car: "Car Test 1", CarModel: "Car Model Test 1", CarColor: "Car Color Test 1", CarModelYear: 1900, CarVin: "", Price: "Rp.0", Availability: true}},
//		&entity.Car{entity.CarDetail{Id: 2, Car: "Car Test 2", CarModel: "Car Model Test 2", CarColor: "Car Color Test 2", CarModelYear: 1900, CarVin: "", Price: "Rp.0", Availability: false}},
//	}
//	repo = TempRepository{
//		repo: cars,
//	}
//
//}
//func TestTempRepository_FindAll(t *testing.T) {
//	t.Run("It should return all car", func(t *testing.T) {
//		resultTest, err := repo.FindAll()
//		assert.Nil(t, err)
//		assert.Equal(t, len(resultTest), 2)
//	})
//}

type TempRepositoryTestSuite struct {
	suite.Suite
	Cars []*entity.Car
	Repo TempRepository
}

func (suite *TempRepositoryTestSuite) SetupTest() {
	suite.Cars = []*entity.Car{
		&entity.Car{entity.CarDetail{Id: 1, Car: "Car Test 1", CarModel: "Car Model Test 1", CarColor: "Car Color Test 1", CarModelYear: 1900, CarVin: "", Price: "Rp.0", Availability: true}},
		&entity.Car{entity.CarDetail{Id: 2, Car: "Car Test 2", CarModel: "Car Model Test 2", CarColor: "Car Color Test 2", CarModelYear: 1900, CarVin: "", Price: "Rp.0", Availability: false}},
	}
	suite.Repo = TempRepository{
		repo: suite.Cars,
	}
}

func (suite *TempRepositoryTestSuite) TestBuildTempRepository() {
	resultTest := NewTempRepository()
	assert.Implements(suite.T(), (*CarRepository)(nil), resultTest)
}
func (suite *TempRepositoryTestSuite) TestFind() {
	resultTest, err := suite.Repo.Find(1)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resultTest)

	resultNegativeTest, err := suite.Repo.Find(1000)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), resultNegativeTest)
}
func (suite *TempRepositoryTestSuite) TestFindAll() {
	resultTest, err := suite.Repo.FindAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(resultTest), 2)
}
func (suite *TempRepositoryTestSuite) TestCreate() {
	newCar := entity.Car{entity.CarDetail{Car: "Car Test 3", CarModel: "Car Model Test 3", CarColor: "Car Color Test 3", CarModelYear: 1950, CarVin: "", Price: "Rp.100", Availability: true}}
	resultTest, err := suite.Repo.Create(&newCar)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(suite.Repo.repo), 3)
	assert.NotZero(suite.T(), resultTest.Id)
}

func TestTempRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TempRepositoryTestSuite))
}
