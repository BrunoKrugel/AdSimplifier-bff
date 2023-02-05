package model

import "time"

type MongoRequest struct {
	UserId      string    `json:"user_id"`
	ProductId   string    `json:"product_id"`
	StoreId     string    `json:"store_id"`
	Date        time.Time `json:"date"`
	ProductName string    `json:"product_name"`
}

type NewRequest struct {
	User_id      string    `json:"user_id"`
	Product_id   string    `json:"product_id"`
	Store_id     string    `json:"store_id"`
	Date         time.Time `json:"date"`
	Product_name string    `json:"product_name"`
	Sales_number int       `json:"sales_number"`
}

type OldRequest struct {
	Date         string `json:"date"`
	Product_id   string `json:"product_id"`
	Product_name string `json:"product_name"`
	Store_id     string `json:"store_id"`
	Sales_number int    `json:"sales_number"`
}
