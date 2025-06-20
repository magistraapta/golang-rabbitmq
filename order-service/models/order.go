package models

type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}
