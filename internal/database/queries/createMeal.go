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
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
            $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
            $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
            $31, $32, $33, $34, $35, $36, $37, $38, $39, $40,
            $41, $42, $43, $44, $45, $46, $47, $48, $49, $50,
            $51, $52, $53, $54, $55, $56, $57, $58, $59, $60,
            $61, $62, $63, $64, $65, $66, $67, $68, $69, $70,
            $71, $72, $73, $74, $75, $76, $77, $78, $79, $80,
            $81, $82, $83, $84, $85, $86, $87, $88, $89, $90,
            $91, $92, $93, $94, $95, $96, $97, $98, $99, $100,
            $101, $102, $103, $104, $105, $106
        )`

	_, err := tx.Exec(query,
		mealId,
		entry.Name,
		entry.Quantity,
		entry.QuantityUnit,
		entry.CreatedAt,
		entry.EnergyKcal,
		entry.EnergyKcalUnit,
		entry.EnergyKj,
		entry.EnergyKjUnit,
		entry.Fat,
		entry.FatUnit,
		entry.SaturatedFat,
		entry.SaturatedFatUnit,
		entry.MonounsaturatedFat,
		entry.MonounsaturatedFatUnit,
		entry.PolyunsaturatedFat,
		entry.PolyunsaturatedFatUnit,
		entry.TransFat,
		entry.TransFatUnit,
		entry.Cholesterol,
		entry.CholesterolUnit,
		entry.Carbohydrates,
		entry.CarbohydratesUnit,
		entry.Sugars,
		entry.SugarsUnit,
		entry.Sucrose,
		entry.SucroseUnit,
		entry.Glucose,
		entry.GlucoseUnit,
		entry.Fructose,
		entry.FructoseUnit,
		entry.Lactose,
		entry.LactoseUnit,
		entry.Maltose,
		entry.MaltoseUnit,
		entry.Galactose,
		entry.GalactoseUnit,
		entry.Starch,
		entry.StarchUnit,
		entry.Polyols,
		entry.PolyolsUnit,
		entry.Fiber,
		entry.FiberUnit,
		entry.Proteins,
		entry.ProteinsUnit,
		entry.Alcohol,
		entry.AlcoholUnit,
		entry.Water,
		entry.WaterUnit,
		entry.Salt,
		entry.SaltUnit,
		entry.Sodium,
		entry.SodiumUnit,
		entry.VitaminA,
		entry.VitaminAUnit,
		entry.VitaminB1,
		entry.VitaminB1Unit,
		entry.VitaminB2,
		entry.VitaminB2Unit,
		entry.VitaminB6,
		entry.VitaminB6Unit,
		entry.VitaminB9,
		entry.VitaminB9Unit,
		entry.VitaminB12,
		entry.VitaminB12Unit,
		entry.VitaminC,
		entry.VitaminCUnit,
		entry.VitaminD,
		entry.VitaminDUnit,
		entry.VitaminE,
		entry.VitaminEUnit,
		entry.VitaminPP,
		entry.VitaminPPUnit,
		entry.PantothenicAcid,
		entry.PantothenicAcidUnit,
		entry.Phylloquinone,
		entry.PhylloquinoneUnit,
		entry.BetaCarotene,
		entry.BetaCaroteneUnit,
		entry.Calcium,
		entry.CalciumUnit,
		entry.Iron,
		entry.IronUnit,
		entry.Magnesium,
		entry.MagnesiumUnit,
		entry.Phosphorus,
		entry.PhosphorusUnit,
		entry.Potassium,
		entry.PotassiumUnit,
		entry.Zinc,
		entry.ZincUnit,
		entry.Copper,
		entry.CopperUnit,
		entry.Manganese,
		entry.ManganeseUnit,
		entry.Selenium,
		entry.SeleniumUnit,
		entry.Iodine,
		entry.IodineUnit,
		entry.NovaGroup,
		entry.NutritionScoreFr,
		entry.FruitsVegLegumesEst,
		entry.FruitsVegLegumesEstUnit,
		entry.FruitsVegNutsEst,
		entry.FruitsVegNutsEstUnit,
	)

	if err != nil {
		return fmt.Errorf("failed to insert meal entry: %w", err)
	}

	return nil
}
