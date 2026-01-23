package models

import (
	"time"
)

type Meal struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
