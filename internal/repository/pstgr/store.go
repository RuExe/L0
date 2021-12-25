package pstgr

import (
	"L0/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	config          *config.Config
	db              *sqlx.DB
	orderRepository *OrderRepository
}

func NewStore(config *config.Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sqlx.Connect(s.config.DBConfig.Driver, s.config.DBConfig.Url)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Order() *OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = NewOrderRepository(s)
	}

	return s.orderRepository
}
