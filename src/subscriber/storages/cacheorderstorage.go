package storages

import (
	"L0/core"
	"errors"
)

type CacheOrderStorage struct {
	Orders map[string]core.Order
}

func (s *CacheOrderStorage) Add(order core.Order) {
	s.Orders[order.OrderUid] = order
}

func (s CacheOrderStorage) Get(id string) (core.Order, error) {
	val, ok := s.Orders[id]

	err := error(nil)
	if !ok {
		err = errors.New("order doesn't exist")
	}

	return val, err
}
