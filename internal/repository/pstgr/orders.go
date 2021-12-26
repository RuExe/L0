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

	var deliveryId int
	delivery := order.Delivery
	query := `INSERT INTO deliveries (name, phone, zip, city, address, region, email)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	if err = tx.QueryRowx(query,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	).Scan(&deliveryId); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	query = `INSERT INTO orders 
             VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	if _, err = tx.Exec(query,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		deliveryId,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	payment := order.Payment
	query = `INSERT INTO payments 
             VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err = tx.Exec(query,
		payment.Transaction,
		payment.RequestId,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	query = `INSERT INTO items 
             VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	for _, item := range order.Items {
		if _, err = tx.Exec(query,
			item.Rid,
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
		); err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (r *OrderRepository) All() ([]domain.Order, error) {
	orders := make([]domain.Order, 0)
	query := "SELECT * FROM orders"
	if err := r.store.db.Select(&orders, query); err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *OrderRepository) GetById(id string) (domain.Order, error) {
	var order domain.Order
	query := "SELECT * FROM orders WHERE order_uid = $1"
	if err := r.store.db.Get(&order, query, id); err != nil {
		return order, err
	}

	return order, nil
}
