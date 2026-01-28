package queries

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SnackLog/meal-service/internal/database/models"
)

func GetMealsByDate(db *sql.DB, username string, dateNotAfter time.Time, dateNotBefore time.Time) ([]models.Meal, error) {
	query := `SELECT id, name, timestamp, created_at
			  FROM meals
			  WHERE timestamp < $1 AND timestamp > $2 AND username = $3
			  ORDER BY timestamp DESC
			  LIMIT 100`
	rows, err := db.Query(query, dateNotAfter, dateNotBefore, username)
	if err != nil {
		return nil, fmt.Errorf("error querying meals by date: %w", err)
	}
	defer rows.Close()

	var meals []models.Meal = []models.Meal{}
	for rows.Next() {
		var meal models.Meal
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Timestamp, &meal.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning meal row: %w", err)
		}
		meals = append(meals, meal)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over meal rows: %w", err)
	}

	err = PopulateMealsWithEntries(db, meals)
	if err != nil {
		return nil, fmt.Errorf("error populating meals with entries: %w", err)
	}

	return meals, nil
}

func GetMealsByDateNotBefore(db *sql.DB, username string, dateNotBefore time.Time) ([]models.Meal, error) {
	query := `SELECT id, name, timestamp, created_at
			  FROM meals
			  WHERE timestamp > $1 AND username = $2
			  ORDER BY timestamp DESC
			  LIMIT 100`
	rows, err := db.Query(query, dateNotBefore, username)
	if err != nil {
		return nil, fmt.Errorf("error querying meals by date not before: %w", err)
	}
	defer rows.Close()

	var meals []models.Meal = []models.Meal{}
	for rows.Next() {
		var meal models.Meal
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Timestamp, &meal.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning meal row: %w", err)
		}
		meals = append(meals, meal)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over meal rows: %w", err)
	}

	err = PopulateMealsWithEntries(db, meals)
	if err != nil {
		return nil, fmt.Errorf("error populating meals with entries: %w", err)
	}

	return meals, nil
}

func GetMealsByDateNotAfter(db *sql.DB, username string, dateNotAfter time.Time) ([]models.Meal, error) {
	query := `SELECT id, name, timestamp, created_at
			  FROM meals
			  WHERE timestamp < $1 AND username = $2
			  ORDER BY timestamp DESC
			  LIMIT 100`
	rows, err := db.Query(query, dateNotAfter, username)
	if err != nil {
		return nil, fmt.Errorf("error querying meals by date not after: %w", err)
	}
	defer rows.Close()

	var meals []models.Meal = []models.Meal{}
	for rows.Next() {
		var meal models.Meal
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Timestamp, &meal.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning meal row: %w", err)
		}
		meals = append(meals, meal)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over meal rows: %w", err)
	}

	err = PopulateMealsWithEntries(db, meals)
	if err != nil {
		return nil, fmt.Errorf("error populating meals with entries: %w", err)
	}

	return meals, nil
}

func PopulateMealsWithEntries(db *sql.DB, meals []models.Meal) error {
	for i := range meals {
		mealEntries, err := getMealEntries(db, meals[i].ID)
		if err != nil {
			return fmt.Errorf("error getting meal entries for meal ID %d: %w", meals[i].ID, err)
		}
		meals[i].MealEntries = *mealEntries
	}
	return nil
}
