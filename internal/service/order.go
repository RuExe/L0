package service

import (
	"L0/internal/domain"
	"L0/internal/repository"
)

type OrderService struct {
	storage *repository.OrderStorage
}

func NewOrderService(storage *repository.OrderStorage) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) AddOrder(order domain.Order) {
	(*s.storage).Add(order)
}

func (s *OrderService) GetOrderList() ([]domain.Order, error) {
	return (*s.storage).All()
}

func (s *OrderService) GetOrder(id string) (domain.Order, error) {
	return (*s.storage).GetById(id)
}
