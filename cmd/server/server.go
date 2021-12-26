package server

import (
	"L0/internal/config"
	"L0/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	config  *config.Config
	service *service.OrderService
}

func NewServer(config *config.Config, service *service.OrderService) *Server {
	return &Server{
		config:  config,
		service: service,
	}
}

func (s *Server) Start() error {
	s.configureRoutes()

	return http.ListenAndServe(s.config.ServerPort, nil)
}

func (s *Server) configureRoutes() {
	http.HandleFunc("/", s.handleGetOrder)
	http.HandleFunc("/List", s.handleGetOrderList)
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
	} else {
		log.Println(err)
	}
}

func (s *Server) handleGetOrderList(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)

	orders, err := s.service.GetOrderList()
	if err == nil {
		result, _ := json.Marshal(orders)
		res.Write(result)
	} else {
		log.Println(err)
	}
}
