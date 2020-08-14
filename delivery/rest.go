package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/repository"
	"github.com/edwardsuwirya/carCollection/useCase"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Server struct {
	useCase          useCase.CarUseCase
	router           *mux.Router
	listeningAddress string
}

func (s *Server) initRouter() error {
	s.router.HandleFunc("/car-collection", s.carCollectionHandler).Methods(http.MethodGet)
	return nil
}

func (s *Server) carCollectionHandler(w http.ResponseWriter, r *http.Request) {
	coll, _ := s.useCase.GetCarCollection()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(coll)
	if err != nil {
		config.Logger.WithField("Rest-Server", "Status").Fatal("Can not parse json")
	}
}
func NewRestServer(c *config.Config) CarDelivery {
	listeningAddress := fmt.Sprintf("%s:%s", c.GetConfigValue("host"), c.GetConfigValue("port"))
	to, _ := strconv.Atoi(c.GetConfigValue("fake_api_timeout"))
	carrepo := repository.NewFakeAPIRepository(c.GetConfigValue("fake_api_url"), to)
	carusecase := useCase.NewCarUseCase(carrepo)
	return &Server{
		useCase:          carusecase,
		router:           mux.NewRouter(),
		listeningAddress: listeningAddress,
	}
}

func (s *Server) Run() {
	err := s.initRouter()
	if err != nil {
		panic(err)
	}
	config.Logger.Debug(fmt.Sprintf("Server runs on %s", s.listeningAddress))
	if err := http.ListenAndServe(s.listeningAddress, s.router); err != nil {
		panic(err)
	}
}
