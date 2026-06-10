package model

import "time"

type User struct {
	ID int64 `json:"id" db:"tg_id"` 
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Balance float64 `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type NewUserData struct{
	TgID int64 `json:"id" db:"tg_id"` 
	Username string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
}