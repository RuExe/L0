package repositories

import "L0/core"

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(order *core.Order) {

}

func (r *OrderRepository) GetById(id string) (*core.Order, error) {
	order := core.Order{}

	if err := r.store.db.QueryRow(
		"SELECT order_uid, track_number FROM orders WHERE order_uid = $1",
		id,
	).Scan(
		&order.OrderUid,
		&order.TrackNumber,
	); err != nil {
		return nil, err
	}

	return &order, nil
}
