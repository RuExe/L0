package repositories

import "L0/core"

type OrderStorage interface {
	Add(order core.Order)
	All() []core.Order
	GetById(id string) (*core.Order, error)
}
