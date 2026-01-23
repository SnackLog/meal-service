package models

import (
	"time"
)

type MealEntry struct {
	ID           int        `json:"id" db:"id"`
	MealID       int        `json:"meal_id" db:"meal_id"`
	Name         string     `json:"name" db:"name"`
	Quantity     float64    `json:"quantity" db:"quantity"`
	QuantityUnit string     `json:"quantity_unit" db:"quantity_unit"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`

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

	// Macros: Carbohydrates & Sugars
	Carbohydrates     *float64 `json:"carbohydrates" db:"carbohydrates"`
	CarbohydratesUnit *string  `json:"carbohydrates_unit" db:"carbohydrates_unit"`
	Sugars            *float64 `json:"sugars" db:"sugars"`
	Sucrose           *float64 `json:"sucrose" db:"sucrose"`
	Glucose           *float64 `json:"glucose" db:"glucose"`
	Fructose          *float64 `json:"fructose" db:"fructose"`
	Lactose           *float64 `json:"lactose" db:"lactose"`
	Maltose           *float64 `json:"maltose" db:"maltose"`
	Galactose         *float64 `json:"galactose" db:"galactose"`
	Starch            *float64 `json:"starch" db:"starch"`
	Polyols           *float64 `json:"polyols" db:"polyols"`
	Fiber             *float64 `json:"fiber" db:"fiber"`

	// Macros: Proteins & Others
	Proteins     *float64 `json:"proteins" db:"proteins"`
	ProteinsUnit *string  `json:"proteins_unit" db:"proteins_unit"`
	Alcohol      *float64 `json:"alcohol" db:"alcohol"`
	Water        *float64 `json:"water" db:"water"`

	// Salt & Sodium
	Salt       *float64 `json:"salt" db:"salt"`
	SaltUnit   *string  `json:"salt_unit" db:"salt_unit"`
	Sodium     *float64 `json:"sodium" db:"sodium"`
	SodiumUnit *string  `json:"sodium_unit" db:"sodium_unit"`

	// Vitamins
	VitaminA        *float64 `json:"vitamin_a" db:"vitamin_a"`
	VitaminB1       *float64 `json:"vitamin_b1" db:"vitamin_b1"`
	VitaminB2       *float64 `json:"vitamin_b2" db:"vitamin_b2"`
	VitaminB6       *float64 `json:"vitamin_b6" db:"vitamin_b6"`
	VitaminB9       *float64 `json:"vitamin_b9" db:"vitamin_b9"`
	VitaminB12      *float64 `json:"vitamin_b12" db:"vitamin_b12"`
	VitaminC        *float64 `json:"vitamin_c" db:"vitamin_c"`
	VitaminD        *float64 `json:"vitamin_d" db:"vitamin_d"`
	VitaminE        *float64 `json:"vitamin_e" db:"vitamin_e"`
	VitaminPP       *float64 `json:"vitamin_pp" db:"vitamin_pp"`
	PantothenicAcid *float64 `json:"pantothenic_acid" db:"pantothenic_acid"`
	Phylloquinone   *float64 `json:"phylloquinone" db:"phylloquinone"`
	BetaCarotene    *float64 `json:"beta_carotene" db:"beta_carotene"`

	// Minerals
	Calcium    *float64 `json:"calcium" db:"calcium"`
	Iron       *float64 `json:"iron" db:"iron"`
	Magnesium  *float64 `json:"magnesium" db:"magnesium"`
	Phosphorus *float64 `json:"phosphorus" db:"phosphorus"`
	Potassium  *float64 `json:"potassium" db:"potassium"`
	Zinc       *float64 `json:"zinc" db:"zinc"`
	Copper     *float64 `json:"copper" db:"copper"`
	Manganese  *float64 `json:"manganese" db:"manganese"`
	Selenium   *float64 `json:"selenium" db:"selenium"`
	Iodine     *float64 `json:"iodine" db:"iodine"`

	// Scores & Estimates
	NovaGroup           *int16   `json:"nova_group" db:"nova_group"`
	NutritionScoreFr    *int16   `json:"nutrition_score_fr" db:"nutrition_score_fr"`
	FruitsVegLegumesEst *float64 `json:"fruits_veg_legumes_est" db:"fruits_veg_legumes_est"`
	FruitsVegNutsEst    *float64 `json:"fruits_veg_nuts_est" db:"fruits_veg_nuts_est"`
}
