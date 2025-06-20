package models

type Payment struct {
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
}