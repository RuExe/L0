package services

import (
	"L0/core"
	"L0/subscriber/repositories"
)

type OrderService struct {
	storage *repositories.OrderStorage
}

func NewOrderService(storage *repositories.OrderStorage) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) AddOrder(order core.Order) {
	(*s.storage).Add(order)
}

func (s *OrderService) GetOrder(id string) (core.Order, error) {
	res, err := (*s.storage).GetById(id)
	return *res, err
}

func (s *OrderService) GetOrderList() ([]core.Order, error) {
	return (*s.storage).All(), nil
}
