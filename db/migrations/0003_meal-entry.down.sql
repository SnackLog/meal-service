ALTER TABLE meal_entries
-- Macros: Fats
DROP COLUMN IF EXISTS cholesterol_unit,

-- Macros: Carbohydrates & Sugars
DROP COLUMN IF EXISTS sugars_unit,
DROP COLUMN IF EXISTS sucrose_unit,
DROP COLUMN IF EXISTS glucose_unit,
DROP COLUMN IF EXISTS fructose_unit,
DROP COLUMN IF EXISTS lactose_unit,
DROP COLUMN IF EXISTS maltose_unit,
DROP COLUMN IF EXISTS galactose_unit,
DROP COLUMN IF EXISTS starch_unit,
DROP COLUMN IF EXISTS polyols_unit,
DROP COLUMN IF EXISTS fiber_unit,

-- Macros: Proteins & Others
DROP COLUMN IF EXISTS alcohol_unit,
DROP COLUMN IF EXISTS water_unit,

-- Vitamins
DROP COLUMN IF EXISTS vitamin_a_unit,
DROP COLUMN IF EXISTS vitamin_b1_unit,
DROP COLUMN IF EXISTS vitamin_b2_unit,
DROP COLUMN IF EXISTS vitamin_b6_unit,
DROP COLUMN IF EXISTS vitamin_b9_unit,
DROP COLUMN IF EXISTS vitamin_b12_unit,
DROP COLUMN IF EXISTS vitamin_c_unit,
DROP COLUMN IF EXISTS vitamin_d_unit,
DROP COLUMN IF EXISTS vitamin_e_unit,
DROP COLUMN IF EXISTS vitamin_pp_unit,
DROP COLUMN IF EXISTS pantothenic_acid_unit,
DROP COLUMN IF EXISTS phylloquinone_unit,
DROP COLUMN IF EXISTS beta_carotene_unit,

-- Minerals
DROP COLUMN IF EXISTS calcium_unit,
DROP COLUMN IF EXISTS iron_unit,
DROP COLUMN IF EXISTS magnesium_unit,
DROP COLUMN IF EXISTS phosphorus_unit,
DROP COLUMN IF EXISTS potassium_unit,
DROP COLUMN IF EXISTS zinc_unit,
DROP COLUMN IF EXISTS copper_unit,
DROP COLUMN IF EXISTS manganese_unit,
DROP COLUMN IF EXISTS selenium_unit,
DROP COLUMN IF EXISTS iodine_unit,

-- Scores & Estimates
DROP COLUMN IF EXISTS fruits_veg_legumes_est_unit,
DROP COLUMN IF EXISTS fruits_veg_nuts_est_unit;
