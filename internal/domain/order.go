package domain

import "time"

type Order struct {
	OrderUid          string
	TrackNumber       string
	Entry             string
	Delivery          string
	Payment           string
	Items             string
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              int
	DateCreated       time.Time
	OofShard          string
}
