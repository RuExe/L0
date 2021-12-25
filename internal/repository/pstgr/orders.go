package pstgr

import (
	"L0/internal/domain"
	"log"
)

type OrderRepository struct {
	store *Store
}

func NewOrderRepository(store *Store) *OrderRepository {
	return &OrderRepository{
		store: store,
	}
}

func (r *OrderRepository) Add(order domain.Order) {
	tx, err := r.store.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO orders 
              VALUES (:id, :track_number, :entry, :delivery, :payment, :items, :locale, :internal_signature, :customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard)`
	if _, err = tx.NamedExec(query, &order); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	/*	delivery := order.Delivery
		query := `INSERT INTO deliveries (name, phone, zip, city, address, region, email)
				  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
		if err := r.store.db.QueryRow(query,
			delivery.Name,
			delivery.Phone,
			delivery.Zip,
			delivery.City,
			delivery.Address,
			delivery.Region,
			delivery.Email,
		).Scan(
			&delivery.Id,
		); err != nil {
			log.Fatal(err)
		}*/
}

func (r *OrderRepository) All() ([]domain.Order, error) {
	orders := make([]domain.Order, 0)
	query := "SELECT * FROM orders"
	if err := r.store.db.Select(&orders, query); err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *OrderRepository) GetById(id int) (domain.Order, error) {
	var order domain.Order
	query := "SELECT * FROM orders WHERE id = $1"
	if err := r.store.db.Get(&order, query, id); err != nil {
		return order, err
	}

	return order, nil
}
