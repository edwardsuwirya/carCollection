package delivery

import (
	"encoding/json"
	"github.com/edwardsuwirya/carCollection/config"
	"github.com/edwardsuwirya/carCollection/useCase"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	useCase useCase.CarUseCase
	router  *mux.Router
}

func (s *Server) init(uc useCase.CarUseCase) error {
	s.useCase = uc
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
func NewRestServer(r *mux.Router, useCase useCase.CarUseCase) CarDelivery {
	srv := new(Server)
	srv.router = r
	err := srv.init(useCase)

	if err != nil {
		panic("Application failed")
	}
	return srv
}
