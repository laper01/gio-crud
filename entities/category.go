package entities

import "time"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// Slug string `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}