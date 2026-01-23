package models

import (
	"time"
)

type Meal struct {
	ID          int         `json:"id" db:"id" binding:"-"`
	Name        string      `json:"name" db:"name" binding:"required"`
	Timestamp   time.Time   `json:"timestamp" db:"timestamp" binding:"required"`
	Username    string      `json:"-" db:"username" binding:"-"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at" binding:"-"`
	MealEntries []MealEntry `json:"meal_entries" db:"-"`
}
