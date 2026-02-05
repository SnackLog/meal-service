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
            cholesterol_unit,
            carbohydrates,
            carbohydrates_unit,
            sugars,
            sugars_unit,
            sucrose,
            sucrose_unit,
            glucose,
            glucose_unit,
            fructose,
            fructose_unit,
            lactose,
            lactose_unit,
            maltose,
            maltose_unit,
            galactose,
            galactose_unit,
            starch,
            starch_unit,
            polyols,
            polyols_unit,
            fiber,
            fiber_unit,
            proteins,
            proteins_unit,
            alcohol,
            alcohol_unit,
            water,
            water_unit,
            salt,
            salt_unit,
            sodium,
            sodium_unit,
            vitamin_a,
            vitamin_a_unit,
            vitamin_b1,
            vitamin_b1_unit,
            vitamin_b2,
            vitamin_b2_unit,
            vitamin_b6,
            vitamin_b6_unit,
            vitamin_b9,
            vitamin_b9_unit,
            vitamin_b12,
            vitamin_b12_unit,
            vitamin_c,
            vitamin_c_unit,
            vitamin_d,
            vitamin_d_unit,
            vitamin_e,
            vitamin_e_unit,
            vitamin_pp,
            vitamin_pp_unit,
            pantothenic_acid,
            pantothenic_acid_unit,
            phylloquinone,
            phylloquinone_unit,
            beta_carotene,
            beta_carotene_unit,
            calcium,
            calcium_unit,
            iron,
            iron_unit,
            magnesium,
            magnesium_unit,
            phosphorus,
            phosphorus_unit,
            potassium,
            potassium_unit,
            zinc,
            zinc_unit,
            copper,
            copper_unit,
            manganese,
            manganese_unit,
            selenium,
            selenium_unit,
            iodine,
            iodine_unit,
            nova_group,
            nutrition_score_fr,
            fruits_veg_legumes_est,
            fruits_veg_legumes_est_unit,
            fruits_veg_nuts_est,
            fruits_veg_nuts_est_unit
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
			&m.CholesterolUnit,
			&m.Carbohydrates,
			&m.CarbohydratesUnit,
			&m.Sugars,
			&m.SugarsUnit,
			&m.Sucrose,
			&m.SucroseUnit,
			&m.Glucose,
			&m.GlucoseUnit,
			&m.Fructose,
			&m.FructoseUnit,
			&m.Lactose,
			&m.LactoseUnit,
			&m.Maltose,
			&m.MaltoseUnit,
			&m.Galactose,
			&m.GalactoseUnit,
			&m.Starch,
			&m.StarchUnit,
			&m.Polyols,
			&m.PolyolsUnit,
			&m.Fiber,
			&m.FiberUnit,
			&m.Proteins,
			&m.ProteinsUnit,
			&m.Alcohol,
			&m.AlcoholUnit,
			&m.Water,
			&m.WaterUnit,
			&m.Salt,
			&m.SaltUnit,
			&m.Sodium,
			&m.SodiumUnit,
			&m.VitaminA,
			&m.VitaminAUnit,
			&m.VitaminB1,
			&m.VitaminB1Unit,
			&m.VitaminB2,
			&m.VitaminB2Unit,
			&m.VitaminB6,
			&m.VitaminB6Unit,
			&m.VitaminB9,
			&m.VitaminB9Unit,
			&m.VitaminB12,
			&m.VitaminB12Unit,
			&m.VitaminC,
			&m.VitaminCUnit,
			&m.VitaminD,
			&m.VitaminDUnit,
			&m.VitaminE,
			&m.VitaminEUnit,
			&m.VitaminPP,
			&m.VitaminPPUnit,
			&m.PantothenicAcid,
			&m.PantothenicAcidUnit,
			&m.Phylloquinone,
			&m.PhylloquinoneUnit,
			&m.BetaCarotene,
			&m.BetaCaroteneUnit,
			&m.Calcium,
			&m.CalciumUnit,
			&m.Iron,
			&m.IronUnit,
			&m.Magnesium,
			&m.MagnesiumUnit,
			&m.Phosphorus,
			&m.PhosphorusUnit,
			&m.Potassium,
			&m.PotassiumUnit,
			&m.Zinc,
			&m.ZincUnit,
			&m.Copper,
			&m.CopperUnit,
			&m.Manganese,
			&m.ManganeseUnit,
			&m.Selenium,
			&m.SeleniumUnit,
			&m.Iodine,
			&m.IodineUnit,
			&m.NovaGroup,
			&m.NutritionScoreFr,
			&m.FruitsVegLegumesEst,
			&m.FruitsVegLegumesEstUnit,
			&m.FruitsVegNutsEst,
			&m.FruitsVegNutsEstUnit,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan meal entry: %w", err)
		}
		mealEntries = append(mealEntries, m)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over meal entries rows: %w", err)
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
