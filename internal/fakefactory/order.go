package fakefactory

import (
	"L0/internal/domain"
	"time"
)

func CreateFakeOrder() domain.Order {
	return domain.Order{
		OrderUid:          1,
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Delivery:          "createFakeDelivery()",
		Payment:           "createFakePayment()",
		Items:             "createFakeItems(1)",
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

func createFakeItems(count int) []domain.Item {
	items := make([]domain.Item, count)
	for i := 0; i < len(items); i++ {
		items[i] = createFakeItem()
	}

	return items
}

func createFakeItem() domain.Item {
	return domain.Item{
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
