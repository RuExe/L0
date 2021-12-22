package factories

import (
	"L0/core"
	"time"
)

func CreateFakeOrder() core.Order {
	dateTime, _ := time.Parse("2006-01-02T15:04:05-0700", "2021-11-26T06:22:19Z")

	order := core.Order{
		OrderUid:          "123321",
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Delivery:          createFakeDelivery(),
		Payment:           createFakePayment(),
		Items:             createFakeItems(1),
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       dateTime,
		OofShard:          "1",
	}

	return order
}

func createFakeDelivery() core.Delivery {
	return core.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
}

func createFakePayment() core.Payment {
	return core.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}
}

func createFakeItems(count int) []core.Item {
	var items []core.Item

	for i := 0; i < count; i++ {
		items = append(items, createFakeItem())
	}

	return items
}

func createFakeItem() core.Item {
	return core.Item{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}
}
