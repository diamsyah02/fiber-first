package models

type Product struct {
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Price int    `json:"price"`
}
