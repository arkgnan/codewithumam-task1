package models

import "time"

type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

type TransactionReport struct {
	TotalRevenue      int               `json:"total_revenue"`
	TotalTransactions int               `json:"total_transactions"`
	BestSeller        BestSellerProduct `json:"best_seller"`
}

type BestSellerProduct struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SoldQuantity int    `json:"sold_quantity"`
}
