package models

import (
	"time"
)

type MealEntry struct {
	ID           int       `json:"-" db:"id" binding:"-"`
	MealID       int       `json:"-" db:"meal_id" binding:"-"`
	Name         string    `json:"name" db:"name" binding:"required"`
	Quantity     float64   `json:"quantity" db:"quantity" binding:"required"`
	QuantityUnit string    `json:"quantity_unit" db:"quantity_unit" binding:"required"`
	CreatedAt    time.Time `json:"created_at" db:"created_at" binding:"-"`

	// Energy
	EnergyKcal     *float64 `json:"energy_kcal" db:"energy_kcal"`
	EnergyKcalUnit *string  `json:"energy_kcal_unit" db:"energy_kcal_unit"`
	EnergyKj       *float64 `json:"energy_kj" db:"energy_kj"`
	EnergyKjUnit   *string  `json:"energy_kj_unit" db:"energy_kj_unit"`

	// Macros: Fats
	Fat                    *float64 `json:"fat" db:"fat"`
	FatUnit                *string  `json:"fat_unit" db:"fat_unit"`
	SaturatedFat           *float64 `json:"saturated_fat" db:"saturated_fat"`
	SaturatedFatUnit       *string  `json:"saturated_fat_unit" db:"saturated_fat_unit"`
	MonounsaturatedFat     *float64 `json:"monounsaturated_fat" db:"monounsaturated_fat"`
	MonounsaturatedFatUnit *string  `json:"monounsaturated_fat_unit" db:"monounsaturated_fat_unit"`
	PolyunsaturatedFat     *float64 `json:"polyunsaturated_fat" db:"polyunsaturated_fat"`
	PolyunsaturatedFatUnit *string  `json:"polyunsaturated_fat_unit" db:"polyunsaturated_fat_unit"`
	TransFat               *float64 `json:"trans_fat" db:"trans_fat"`
	TransFatUnit           *string  `json:"trans_fat_unit" db:"trans_fat_unit"`
	Cholesterol            *float64 `json:"cholesterol" db:"cholesterol"`
	CholesterolUnit        *string  `json:"cholesterol_unit" db:"cholesterol_unit"`

	// Macros: Carbohydrates & Sugars
	Carbohydrates     *float64 `json:"carbohydrates" db:"carbohydrates"`
	CarbohydratesUnit *string  `json:"carbohydrates_unit" db:"carbohydrates_unit"`
	Sugars            *float64 `json:"sugars" db:"sugars"`
	SugarsUnit        *string  `json:"sugars_unit" db:"sugars_unit"`
	Sucrose           *float64 `json:"sucrose" db:"sucrose"`
	SucroseUnit        *string  `json:"sucrose_unit" db:"sucrose_unit"`
	Glucose           *float64 `json:"glucose" db:"glucose"`
	GlucoseUnit        *string  `json:"glucose_unit" db:"glucose_unit"`
	Fructose          *float64 `json:"fructose" db:"fructose"`
	FructoseUnit       *string  `json:"fructose_unit" db:"fructose_unit"`
	Lactose           *float64 `json:"lactose" db:"lactose"`
	LactoseUnit        *string  `json:"lactose_unit" db:"lactose_unit"`
	Maltose           *float64 `json:"maltose" db:"maltose"`
	MaltoseUnit        *string  `json:"maltose_unit" db:"maltose_unit"`
	Galactose         *float64 `json:"galactose" db:"galactose"`
	GalactoseUnit      *string  `json:"galactose_unit" db:"galactose_unit"`
	Starch            *float64 `json:"starch" db:"starch"`
	StarchUnit         *string  `json:"starch_unit" db:"starch_unit"`
	Polyols           *float64 `json:"polyols" db:"polyols"`
	PolyolsUnit        *string  `json:"polyols_unit" db:"polyols_unit"`
	Fiber             *float64 `json:"fiber" db:"fiber"`
	FiberUnit          *string  `json:"fiber_unit" db:"fiber_unit"`

	// Macros: Proteins & Others
	Proteins     *float64 `json:"proteins" db:"proteins"`
	ProteinsUnit *string  `json:"proteins_unit" db:"proteins_unit"`
	Alcohol      *float64 `json:"alcohol" db:"alcohol"`
	AlcoholUnit  *string  `json:"alcohol_unit" db:"alcohol_unit"`
	Water        *float64 `json:"water" db:"water"`
	WaterUnit    *string  `json:"water_unit" db:"water_unit"`

	// Salt & Sodium
	Salt       *float64 `json:"salt" db:"salt"`
	SaltUnit   *string  `json:"salt_unit" db:"salt_unit"`
	Sodium     *float64 `json:"sodium" db:"sodium"`
	SodiumUnit *string  `json:"sodium_unit" db:"sodium_unit"`

	// Vitamins
	VitaminA            *float64 `json:"vitamin_a" db:"vitamin_a"`
	VitaminAUnit        *string  `json:"vitamin_a_unit" db:"vitamin_a_unit"`
	VitaminB1           *float64 `json:"vitamin_b1" db:"vitamin_b1"`
	VitaminB1Unit       *string  `json:"vitamin_b1_unit" db:"vitamin_b1_unit"`
	VitaminB2           *float64 `json:"vitamin_b2" db:"vitamin_b2"`
	VitaminB2Unit       *string  `json:"vitamin_b2_unit" db:"vitamin_b2_unit"`
	VitaminB6           *float64 `json:"vitamin_b6" db:"vitamin_b6"`
	VitaminB6Unit       *string  `json:"vitamin_b6_unit" db:"vitamin_b6_unit"`
	VitaminB9           *float64 `json:"vitamin_b9" db:"vitamin_b9"`
	VitaminB9Unit       *string  `json:"vitamin_b9_unit" db:"vitamin_b9_unit"`
	VitaminB12          *float64 `json:"vitamin_b12" db:"vitamin_b12"`
	VitaminB12Unit      *string  `json:"vitamin_b12_unit" db:"vitamin_b12_unit"`
	VitaminC            *float64 `json:"vitamin_c" db:"vitamin_c"`
	VitaminCUnit        *string  `json:"vitamin_c_unit" db:"vitamin_c_unit"`
	VitaminD            *float64 `json:"vitamin_d" db:"vitamin_d"`
	VitaminDUnit        *string  `json:"vitamin_d_unit" db:"vitamin_d_unit"`
	VitaminE            *float64 `json:"vitamin_e" db:"vitamin_e"`
	VitaminEUnit        *string  `json:"vitamin_e_unit" db:"vitamin_e_unit"`
	VitaminPP           *float64 `json:"vitamin_pp" db:"vitamin_pp"`
	VitaminPPUnit       *string  `json:"vitamin_pp_unit" db:"vitamin_pp_unit"`
	PantothenicAcid     *float64 `json:"pantothenic_acid" db:"pantothenic_acid"`
	PantothenicAcidUnit *string  `json:"pantothenic_acid_unit" db:"pantothenic_acid_unit"`
	Phylloquinone       *float64 `json:"phylloquinone" db:"phylloquinone"`
	PhylloquinoneUnit   *string  `json:"phylloquinone_unit" db:"phylloquinone_unit"`
	BetaCarotene        *float64 `json:"beta_carotene" db:"beta_carotene"`
	BetaCaroteneUnit    *string  `json:"beta_carotene_unit" db:"beta_carotene_unit"`

	// Minerals
	Calcium           *float64 `json:"calcium" db:"calcium"`
	CalciumUnit       *string  `json:"calcium_unit" db:"calcium_unit"`
	Iron              *float64 `json:"iron" db:"iron"`
	IronUnit          *string  `json:"iron_unit" db:"iron_unit"`
	Magnesium         *float64 `json:"magnesium" db:"magnesium"`
	MagnesiumUnit     *string  `json:"magnesium_unit" db:"magnesium_unit"`
	Phosphorus        *float64 `json:"phosphorus" db:"phosphorus"`
	PhosphorusUnit     *string  `json:"phosphorus_unit" db:"phosphorus_unit"`
	Potassium         *float64 `json:"potassium" db:"potassium"`
	PotassiumUnit     *string  `json:"potassium_unit" db:"potassium_unit"`
	Zinc              *float64 `json:"zinc" db:"zinc"`
	ZincUnit          *string  `json:"zinc_unit" db:"zinc_unit"`
	Copper            *float64 `json:"copper" db:"copper"`
	CopperUnit        *string  `json:"copper_unit" db:"copper_unit"`
	Manganese         *float64 `json:"manganese" db:"manganese"`
	ManganeseUnit     *string  `json:"manganese_unit" db:"manganese_unit"`
	Selenium          *float64 `json:"selenium" db:"selenium"`
	SeleniumUnit      *string  `json:"selenium_unit" db:"selenium_unit"`
	Iodine            *float64 `json:"iodine" db:"iodine"`
	IodineUnit        *string  `json:"iodine_unit" db:"iodine_unit"`

	// Scores & Estimates
	NovaGroup              *int16   `json:"nova_group" db:"nova_group"`
	NutritionScoreFr       *int16   `json:"nutrition_score_fr" db:"nutrition_score_fr"`
	FruitsVegLegumesEst    *float64 `json:"fruits_veg_legumes_est" db:"fruits_veg_legumes_est"`
	FruitsVegLegumesEstUnit *string  `json:"fruits_veg_legumes_est_unit" db:"fruits_veg_legumes_est_unit"`
	FruitsVegNutsEst       *float64 `json:"fruits_veg_nuts_est" db:"fruits_veg_nuts_est"`
	FruitsVegNutsEstUnit   *string  `json:"fruits_veg_nuts_est_unit" db:"fruits_veg_nuts_est_unit"`
}
