package model

type MongoRequest struct {
	UserId      string `json:"user_id"`
	ProductId   string `json:"product_id"`
	StoreId     string `json:"store_id"`
	Date        string `json:"date"`
	ProductName string `json:"product_name"`
}
