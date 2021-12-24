package repositories

import (
	"L0/core"
	"errors"
)

type CacheOrderStorage struct {
	storage *OrderStorage
	orders  map[string]core.Order
}

func NewCacheOrderStorage(storage *OrderStorage) *CacheOrderStorage {
	orders := make(map[string]core.Order)
	for _, v := range (*storage).All() {
		orders[v.OrderUid] = v
	}

	return &CacheOrderStorage{
		storage: storage,
		orders:  orders,
	}
}

func (s *CacheOrderStorage) Add(order core.Order) {
	(*s.storage).Add(order)
	s.orders[order.OrderUid] = order
}

func (s *CacheOrderStorage) All() []core.Order {
	orders := make([]core.Order, 0)
	for _, v := range s.orders {
		orders = append(orders, v)
	}

	return orders
}

func (s CacheOrderStorage) GetById(id string) (*core.Order, error) {
	val, ok := s.orders[id]

	err := error(nil)
	if !ok {
		err = errors.New("order doesn't exist")
	}

	return &val, err
}
