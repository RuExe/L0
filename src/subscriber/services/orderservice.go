package services

import "L0/core"

type OrderService struct {
	orders map[string]core.Order
}

func (s *OrderService) Add(order core.Order) {
	s.orders[order.OrderUid] = order
}

func (s OrderService) Get(id string) core.Order {
	return s.orders[id]
}
