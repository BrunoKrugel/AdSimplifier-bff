package model

import "time"

type MongoSales struct {
	User_id      string    `json:"user_id"`
	Product_id   string    `json:"product_id"`
	Store_id     string    `json:"store_id"`
	Date         time.Time `json:"date"`
	Product_name string    `json:"product_name"`
}

type MongoSalesInfo struct {
	User_id        string    `json:"user_id"`
	Product_name   string    `json:"product_name"`
	Store_id       string    `json:"store_id"`
	Date           time.Time `json:"date"`
	Payment_method string    `json:"payment_method"`
	Order_status   string    `json:"order_status"`
	Commission     int       `json:"commission"`
	Order_ref      string    `json:"order_ref"`
	Src            string    `json:"src"`
	Sck            string    `json:"sck"`
	Utm_medium     string    `json:"utm_medium"`
	Utm_content    string    `json:"utm_content"`
	Utm_campaign   string    `json:"utm_campaign"`
}
