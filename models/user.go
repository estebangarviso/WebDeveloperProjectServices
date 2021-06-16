package models

import "time"

//Customer datos del usuario
type User struct {
	ID      		string		`json:"id"`
	FirstName		string		`json:"firstname"`
	LastName		string		`json:"lastname"`
	Email   		string		`json:"email"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at,omitempty"`
}

//Customers lista de usuarios
type Users []User