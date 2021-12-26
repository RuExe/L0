package repository

import (
	"L0/internal/domain"
)

type OrderStorage interface {
	Add(order domain.Order)
	All() ([]domain.Order, error)
	GetById(id string) (domain.Order, error)
}
