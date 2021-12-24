package repositories

import (
	"L0/core"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config          *core.Config
	db              *sql.DB
	orderRepository *OrderRepository
}

func NewStore(config *core.Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open(s.config.DBConfig.Driver, s.config.DBConfig.Url)
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
