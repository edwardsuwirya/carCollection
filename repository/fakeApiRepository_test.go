package repository

import (
	"github.com/edwardsuwirya/carCollection/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeApiRepositoryTestSuite struct {
	suite.Suite
	dummyUrl string
	timeout  int
	Repo     FakeApiRepository
}

func (suite *FakeApiRepositoryTestSuite) SetupTest() {
	suite.dummyUrl = "http://example.com"
	suite.timeout = 0
}

func (suite *FakeApiRepositoryTestSuite) TestBuildFakeApiRepository() {
	resultTest := NewFakeAPIRepository(suite.dummyUrl, suite.timeout)
	assert.Implements(suite.T(), (*CarRepository)(nil), resultTest)
	assert.Equal(suite.T(), resultTest.(*FakeApiRepository).url, suite.dummyUrl)
}

func (suite *FakeApiRepositoryTestSuite) TestCreate() {
	suite.Repo = FakeApiRepository{
		url:        suite.dummyUrl,
		httpClient: nil,
	}
	newCar := entity.Car{entity.CarDetail{Car: "Car Test 3", CarModel: "Car Model Test 3", CarColor: "Car Color Test 3", CarModelYear: 1950, CarVin: "", Price: "Rp.100", Availability: true}}
	resultTest, err := suite.Repo.Create(&newCar)
	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), resultTest.Id)
}

func (suite *FakeApiRepositoryTestSuite) TestFindAll() {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte(`{"cars": [{"id": 1, "car": "Mitsubishi", "car_model": "Montero", "car_color": "Yellow", "car_model_year": 2002, "car_vin": "SAJWJ0FF3F8321657", "price": "$2814.46", "availability": false}]}`))
	}))
	defer server.Close()
	suite.Repo = FakeApiRepository{
		url:        server.URL,
		httpClient: server.Client(),
	}
	resultTest, err := suite.Repo.FindAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(resultTest), 1)
}

func TestFakeApiRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FakeApiRepositoryTestSuite))
}
