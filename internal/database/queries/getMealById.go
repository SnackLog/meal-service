package queries

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/meal-service/internal/database/models"
)

func GetMealById(db *sql.DB, id int, username string) (*models.Meal, error) {
	meal, err := getMealById(db, id, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get meal by id: %w", err)
	}
	if meal == nil {
		return nil, nil
	}

	mealEntries, err := getMealEntries(db, meal.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get meal entries: %w", err)
	}
	if mealEntries == nil {
		meal.MealEntries = []models.MealEntry{}
	} else {
		meal.MealEntries = *mealEntries
	}

	return meal, nil
}

func getMealEntries(db *sql.DB, mealId int) (*[]models.MealEntry, error) {
	query := `
		SELECT 
			id,
			meal_id,
			name,
			quantity,
			quantity_unit,
			created_at,
			energy_kcal,
			energy_kcal_unit,
			energy_kj,
			energy_kj_unit,
			fat,
			fat_unit,
			saturated_fat,
			saturated_fat_unit,
			monounsaturated_fat,
			monounsaturated_fat_unit,
			polyunsaturated_fat,
			polyunsaturated_fat_unit,
			trans_fat,
			trans_fat_unit,
			cholesterol,
			carbohydrates,
			carbohydrates_unit,
			sugars,
			sucrose,
			glucose,
			fructose,
			lactose,
			maltose,
			galactose,
			starch,
			polyols,
			fiber,
			proteins,
			proteins_unit,
			alcohol,
			water,
			salt,
			salt_unit,
			sodium,
			sodium_unit,
			vitamin_a,
			vitamin_b1,
			vitamin_b2,
			vitamin_b6,
			vitamin_b9,
			vitamin_b12,
			vitamin_c,
			vitamin_d,
			vitamin_e,
			vitamin_pp,
			pantothenic_acid,
			phylloquinone,
			beta_carotene,
			calcium,
			iron,
			magnesium,
			phosphorus,
			potassium,
			zinc,
			copper,
			manganese,
			selenium,
			iodine,
			nova_group,
			nutrition_score_fr,
			fruits_veg_legumes_est,
			fruits_veg_nuts_est
		FROM meal_entries 
		WHERE meal_id = $1`
	rows, err := db.Query(query, mealId)
	if err != nil {
		return nil, fmt.Errorf("failed to query meal entries: %w", err)
	}
	defer rows.Close()

	var mealEntries []models.MealEntry
	for rows.Next() {
		var m models.MealEntry
		err := rows.Scan(
			&m.ID,
			&m.MealID,
			&m.Name,
			&m.Quantity,
			&m.QuantityUnit,
			&m.CreatedAt,
			&m.EnergyKcal,
			&m.EnergyKcalUnit,
			&m.EnergyKj,
			&m.EnergyKjUnit,
			&m.Fat,
			&m.FatUnit,
			&m.SaturatedFat,
			&m.SaturatedFatUnit,
			&m.MonounsaturatedFat,
			&m.MonounsaturatedFatUnit,
			&m.PolyunsaturatedFat,
			&m.PolyunsaturatedFatUnit,
			&m.TransFat,
			&m.TransFatUnit,
			&m.Cholesterol,
			&m.Carbohydrates,
			&m.CarbohydratesUnit,
			&m.Sugars,
			&m.Sucrose,
			&m.Glucose,
			&m.Fructose,
			&m.Lactose,
			&m.Maltose,
			&m.Galactose,
			&m.Starch,
			&m.Polyols,
			&m.Fiber,
			&m.Proteins,
			&m.ProteinsUnit,
			&m.Alcohol,
			&m.Water,
			&m.Salt,
			&m.SaltUnit,
			&m.Sodium,
			&m.SodiumUnit,
			&m.VitaminA,
			&m.VitaminB1,
			&m.VitaminB2,
			&m.VitaminB6,
			&m.VitaminB9,
			&m.VitaminB12,
			&m.VitaminC,
			&m.VitaminD,
			&m.VitaminE,
			&m.VitaminPP,
			&m.PantothenicAcid,
			&m.Phylloquinone,
			&m.BetaCarotene,
			&m.Calcium,
			&m.Iron,
			&m.Magnesium,
			&m.Phosphorus,
			&m.Potassium,
			&m.Zinc,
			&m.Copper,
			&m.Manganese,
			&m.Selenium,
			&m.Iodine,
			&m.NovaGroup,
			&m.NutritionScoreFr,
			&m.FruitsVegLegumesEst,
			&m.FruitsVegNutsEst,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan meal entry: %w", err)
		}
		mealEntries = append(mealEntries, m)
	}
	return &mealEntries, nil
}

func getMealById(db *sql.DB, id int, username string) (*models.Meal, error) {
	var meal models.Meal
	query := `SELECT id, name, username, timestamp, created_at FROM meals WHERE id = $1 AND username = $2`
	row := db.QueryRow(query, id, username)
	err := row.Scan(&meal.ID, &meal.Name, &meal.Username, &meal.Timestamp, &meal.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to scan meal: %w", err)
	}
	return &meal, nil
}
