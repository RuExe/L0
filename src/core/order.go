package core

import "time"

type Order struct {
	OrderUid    string
	TrackNumber string
	Entry       string
	Delivery
	Payment
	Items             []Item
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              int
	DateCreated       time.Time
	OofShard          string
}
