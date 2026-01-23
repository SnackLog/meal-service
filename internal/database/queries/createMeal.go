package queries

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/meal-service/internal/database/models"
)

func CreateMeal(db *sql.DB, username string, meal models.Meal) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	mealID, err := CreateMealTx(tx, username, meal)
	if err != nil {
		return 0, fmt.Errorf("failed to create meal: %v", err)
	}

	for _, entry := range meal.MealEntries {
		err = InsertMealEntryTx(tx, mealID, entry)
		if err != nil {
			return 0, fmt.Errorf("failed to insert meal entry: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return mealID, nil
}

func CreateMealTx(tx *sql.Tx, username string, meal models.Meal) (int, error) {
	var mealID int
	err := tx.QueryRow(`
		INSERT INTO meals (name, timestamp, username)
		VALUES ($1, $2, $3)
		RETURNING id
	`, meal.Name, meal.Timestamp, username).Scan(&mealID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert meal: %v", err)
	}
	return mealID, nil
}

func InsertMealEntryTx(tx *sql.Tx, mealId int, entry models.MealEntry) error {
	query := `
		INSERT INTO meal_entries (
			meal_id, name, quantity, quantity_unit, created_at,
			energy_kcal, energy_kcal_unit, energy_kj, energy_kj_unit,
			fat, fat_unit, saturated_fat, saturated_fat_unit, monounsaturated_fat, monounsaturated_fat_unit, polyunsaturated_fat, polyunsaturated_fat_unit, trans_fat, trans_fat_unit, cholesterol,
			carbohydrates, carbohydrates_unit, sugars, sucrose, glucose, fructose, lactose, maltose, galactose, starch, polyols, fiber,
			proteins, proteins_unit, alcohol, water,
			salt, salt_unit, sodium, sodium_unit,
			vitamin_a, vitamin_b1, vitamin_b2, vitamin_b6, vitamin_b9, vitamin_b12, vitamin_c, vitamin_d, vitamin_e, vitamin_pp, pantothenic_acid, phylloquinone, beta_carotene,
			calcium, iron, magnesium, phosphorus, potassium, zinc, copper, manganese, selenium, iodine,
			nova_group, nutrition_score_fr, fruits_veg_legumes_est, fruits_veg_nuts_est
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32,
			$33, $34, $35, $36,
			$37, $38, $39, $40,
			$41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53,
			$54, $55, $56, $57, $58, $59, $60, $61, $62, $63,
			$64, $65, $66, $67
		)`

	_, err := tx.Exec(query,
		mealId, entry.Name, entry.Quantity, entry.QuantityUnit, entry.CreatedAt,
		entry.EnergyKcal, entry.EnergyKcalUnit, entry.EnergyKj, entry.EnergyKjUnit,
		entry.Fat, entry.FatUnit, entry.SaturatedFat, entry.SaturatedFatUnit, entry.MonounsaturatedFat, entry.MonounsaturatedFatUnit, entry.PolyunsaturatedFat, entry.PolyunsaturatedFatUnit, entry.TransFat, entry.TransFatUnit, entry.Cholesterol,
		entry.Carbohydrates, entry.CarbohydratesUnit, entry.Sugars, entry.Sucrose, entry.Glucose, entry.Fructose, entry.Lactose, entry.Maltose, entry.Galactose, entry.Starch, entry.Polyols, entry.Fiber,
		entry.Proteins, entry.ProteinsUnit, entry.Alcohol, entry.Water,
		entry.Salt, entry.SaltUnit, entry.Sodium, entry.SodiumUnit,
		entry.VitaminA, entry.VitaminB1, entry.VitaminB2, entry.VitaminB6, entry.VitaminB9, entry.VitaminB12, entry.VitaminC, entry.VitaminD, entry.VitaminE, entry.VitaminPP, entry.PantothenicAcid, entry.Phylloquinone, entry.BetaCarotene,
		entry.Calcium, entry.Iron, entry.Magnesium, entry.Phosphorus, entry.Potassium, entry.Zinc, entry.Copper, entry.Manganese, entry.Selenium, entry.Iodine,
		entry.NovaGroup, entry.NutritionScoreFr, entry.FruitsVegLegumesEst, entry.FruitsVegNutsEst,
	)

	if err != nil {
		return fmt.Errorf("failed to insert meal entry: %w", err)
	}

	return nil
}
