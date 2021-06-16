package models

import "time"

//Order datos del carro de compras
type Order struct {
	ID          		string 		`json:"id"`
	CreatedAt 			time.Time	`json:"created_at"`
	UpdatedAt			time.Time	`json:"updated_at,omitempty"`
}

//Orders lista de carros de compras
type Orders []Order