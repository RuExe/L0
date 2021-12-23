package services

import (
	core "L0/core"
	"L0/subscriber/repositories"
	"L0/subscriber/storages"
)

type OrderService struct {
	orderRepository *repositories.OrderRepository
	cacheStorage    *storages.CacheOrderStorage
}

func NewOrderService(repository *repositories.OrderRepository) *OrderService {
	test, _ := repository.GetById("1")

	storage := &storages.CacheOrderStorage{
		Orders: map[string]core.Order{
			"1": *test,
		},
	}

	return &OrderService{
		orderRepository: repository,
		cacheStorage:    storage,
	}
}

func (s *OrderService) AddOrder(order core.Order) {
	s.cacheStorage.Add(order)
}

func (s OrderService) GetOrder(id string) (core.Order, error) {
	return s.cacheStorage.Get(id)
}
