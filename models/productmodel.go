package models

type Product struct {
	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
	Description string `json:"description"`
}