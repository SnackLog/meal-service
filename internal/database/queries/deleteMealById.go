package queries

import (
	"database/sql"
	"fmt"
)

func DeleteMealById(db *sql.DB, mealId int, username string) (sql.Result, error) {
	result, err := db.Exec(`
		DELETE FROM meals
		WHERE id = $1 AND username = $2
	`, mealId, username)
	if err != nil {
		return nil, fmt.Errorf("failed to delete meal: %v", err)
	}

	return result, nil
}
