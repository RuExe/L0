package fakefactory

import (
	"L0/internal/domain"
	"fmt"
	"time"
)

func CreateFakeOrder() domain.Order {
	return domain.Order{
		OrderUid:          "b563feb7b2b84b6test2",
		TrackNumber:       "WBILMTESTTRACK2",
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
		DateCreated:       time.Date(2021, 11, 26, 5, 22, 19, 0, time.UTC),
		OofShard:          "1",
	}
}

func createFakeDelivery() domain.Delivery {
	return domain.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
}

func createFakePayment() domain.Payment {
	return domain.Payment{
		Transaction:  "b563feb7b2b84b6test2",
		RequestId:    "",
		Currency:     "RU",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}
}

func createFakeItems(count int) []domain.Item {
	items := make([]domain.Item, count)
	for i := 0; i < len(items); i++ {
		items[i] = domain.Item{
			ChrtId:      9934930,
			TrackNumber: "WBILMTESTTRACK2",
			Price:       453,
			Rid:         fmt.Sprint("ab4219087a764ae0btest22%s", i),
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NmId:        2389212,
			Brand:       "Vivienne Sabo",
			Status:      202,
		}
	}
	return items
}
