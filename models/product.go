package models

import "time"

//Product datos del producto
type Product struct {
	ID          		string 		`json:"id"`
	Name        		string 		`json:"name"`
	ShortDescription 	string 		`json:"short_description"`
	LongDescription		string		`json:"long_description"`
	Price       		string 		`json:"Price"`
	CreatedAt 			time.Time	`json:"created_at"`
	UpdatedAt			time.Time	`json:"updated_at,omitempty"`
}

//Products lista de productos
type Products []Product