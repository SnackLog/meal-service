package meal

import "database/sql"

// MealController holds the database connection for meal handlers
type MealController struct {
	DB *sql.DB
}
