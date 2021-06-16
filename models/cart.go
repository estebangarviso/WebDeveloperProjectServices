package models

import "time"

//Cart datos del carro de compras
type Cart struct {
	ID          		string 		`json:"id"`
	CreatedAt 			time.Time	`json:"created_at"`
	UpdatedAt			time.Time	`json:"updated_at,omitempty"`
}

//Carts lista de carros de compras
type Carts []Cart