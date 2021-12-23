package classes

import (
	"L0/core"
	"L0/subscriber/repositories"
	"L0/subscriber/services"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	config  *core.Config
	store   *repositories.Store
	service *services.OrderService
}

func NewServer(config *core.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() error {
	if err := s.configureStore(); err != nil {
		log.Fatal(err)
	}

	s.service = services.NewOrderService(s.store.Order())

	s.configureRoutes()

	return http.ListenAndServe(s.config.ServerPort, nil)
}

func (s *Server) configureStore() error {
	st := repositories.NewStore(s.config)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *Server) configureRoutes() {
	http.HandleFunc("/", s.handleGetOrder)
}

func (s *Server) handleGetOrder(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)

	query := req.URL.Query()
	id := query.Get("id")

	order, err := s.service.GetOrder(id)
	if err == nil {
		result, _ := json.Marshal(order)
		res.Write(result)
	}
}
