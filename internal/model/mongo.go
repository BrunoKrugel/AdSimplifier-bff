package model

import "time"

type MongoSales struct {
	UserId      string    `json:"user_id"`
	ProductId   string    `json:"product_id"`
	StoreId     string    `json:"store_id"`
	Date        time.Time `json:"date"`
	ProductName string    `json:"product_name"`
}

type MongoSalesInfo struct {
	UserId        string    `json:"user_id"`
	ProductName   string    `json:"product_name"`
	StoreId       string    `json:"store_id"`
	Date          time.Time `json:"date"`
	PaymentMethod string    `json:"payment_method"`
	OrderStatus   string    `json:"order_status"`
	Commission    int       `json:"commission"`
}
