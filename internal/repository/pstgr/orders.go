package pstgr

import (
	"L0/internal/domain"
	"errors"
	"github.com/jmoiron/sqlx"
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
	query := `SELECT 
			  o.order_uid, o.track_number, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id, o.delivery_service,
			  d.id, d."name", d.phone, d.zip, d.city, d.address, d.region, d.email,
			  pay."transaction", pay.request_id, pay.currency, pay.provider, pay.amount, pay.payment_dt, pay.bank, pay.delivery_cost, pay.goods_total, pay.custom_fee
			  FROM orders o 
			  JOIN deliveries d ON o.delivery_id = d.id
			  JOIN payments pay ON pay.transaction = o.order_uid`
	if err := r.store.db.Select(&orders, query); err != nil {
		return orders, err
	}

	tracks := make([]string, len(orders))
	for _, v := range orders {
		tracks = append(tracks, v.TrackNumber)
	}

	query, args, err := sqlx.In("SELECT * FROM items WHERE track_number IN (?)", tracks)
	if err != nil {
		log.Fatal(err)
	}
	query = r.store.db.Rebind(query)

	items := make([]domain.Item, 0)
	if err := r.store.db.Select(&items, query, args...); err != nil {
		log.Fatal(err)
	}

	ordersMap := make(map[string]*domain.Order, len(orders))
	for i := range orders {
		ordersMap[orders[i].TrackNumber] = &orders[i]
	}

	for _, v := range items {
		ordersMap[v.TrackNumber].Items = append(ordersMap[v.TrackNumber].Items, v)
	}

	return orders, nil
}

func (r *OrderRepository) GetById(id string) (domain.Order, error) {
	var order domain.Order
	return order, errors.New("not implemented")
}
