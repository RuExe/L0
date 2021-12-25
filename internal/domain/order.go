package domain

import "time"

type Order struct {
	OrderUid          int       `db:"id"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	Delivery          string    `db:"delivery"`
	Payment           string    `db:"payment"`
	Items             string    `db:"items"`
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	CustomerId        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	Shardkey          string    `db:"shardkey"`
	SmId              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}
