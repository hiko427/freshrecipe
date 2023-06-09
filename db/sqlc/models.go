// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"time"
)

type Recipe struct {
	ID      int64  `json:"id"`
	Cuisine string `json:"cuisine"`
	// must be positive
	NumPeople           int64     `json:"num_people"`
	Ingredients         string    `json:"ingredients"`
	ExcludedIngredients string    `json:"excluded_ingredients"`
	Mealtype            string    `json:"mealtype"`
	Recipe              string    `json:"recipe"`
	CreatedAt           time.Time `json:"created_at"`
}
