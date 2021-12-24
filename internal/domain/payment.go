package domain

type Payment struct {
	Transaction  string
	RequestId    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDt    int
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}
