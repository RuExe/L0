package pstgr

import (
	"L0/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config          *config.Config
	Db              *sql.DB
	orderRepository *OrderRepository
}

func NewStore(config *config.Config) *Store {
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

	s.Db = db

	return nil
}

func (s *Store) Close() {
	s.Db.Close()
}

func (s *Store) Order() *OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = NewOrderRepository(s)
	}

	return s.orderRepository
}
