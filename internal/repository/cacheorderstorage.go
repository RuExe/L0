package repository

import (
	"L0/internal/domain"
	"errors"
)

type CacheOrderStorage struct {
	storage *OrderStorage
	orders  map[string]domain.Order
}

func NewCacheOrderStorage(storage *OrderStorage) *CacheOrderStorage {
	orders := make(map[string]domain.Order)
	for _, v := range (*storage).All() {
		orders[v.OrderUid] = v
	}

	return &CacheOrderStorage{
		storage: storage,
		orders:  orders,
	}
}

func (s *CacheOrderStorage) Add(order domain.Order) {
	(*s.storage).Add(order)
	s.orders[order.OrderUid] = order
}

func (s *CacheOrderStorage) All() []domain.Order {
	orders := make([]domain.Order, 0)
	for _, v := range s.orders {
		orders = append(orders, v)
	}

	return orders
}

func (s *CacheOrderStorage) GetById(id string) (*domain.Order, error) {
	val, ok := s.orders[id]

	err := error(nil)
	if !ok {
		err = errors.New("order doesn't exist")
	}

	return &val, err
}
