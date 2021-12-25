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
	_, err := r.store.Db.Exec(
		`INSERT INTO orders (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature,customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
			   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.Delivery,
		order.Payment,
		order.Items,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func (r *OrderRepository) All() []domain.Order {
	rows, err := r.store.Db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orders := make([]domain.Order, 0)
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(
			&order.OrderUid,
			&order.TrackNumber,
			&order.Entry,
			&order.Delivery,
			&order.Payment,
			&order.Items,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
		)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return orders
}

func (r *OrderRepository) GetById(id string) (*domain.Order, error) {
	order := new(domain.Order)

	if err := r.store.Db.QueryRow(
		"SELECT * FROM orders WHERE order_uid = $1",
		id,
	).Scan(
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&order.Delivery,
		&order.Payment,
		&order.Items,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
	); err != nil {
		return nil, err
	}

	return order, nil
}
