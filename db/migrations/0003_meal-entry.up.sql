ALTER TABLE meal_entries
-- Macros: Fats
ADD COLUMN cholesterol_unit VARCHAR(10),

-- Macros: Carbohydrates & Sugars
ADD COLUMN sugars_unit VARCHAR(10),
ADD COLUMN sucrose_unit VARCHAR(10),
ADD COLUMN glucose_unit VARCHAR(10),
ADD COLUMN fructose_unit VARCHAR(10),
ADD COLUMN lactose_unit VARCHAR(10),
ADD COLUMN maltose_unit VARCHAR(10),
ADD COLUMN galactose_unit VARCHAR(10),
ADD COLUMN starch_unit VARCHAR(10),
ADD COLUMN polyols_unit VARCHAR(10),
ADD COLUMN fiber_unit VARCHAR(10),

-- Macros: Proteins & Others
ADD COLUMN alcohol_unit VARCHAR(10),
ADD COLUMN water_unit VARCHAR(10),

-- Vitamins
ADD COLUMN vitamin_a_unit VARCHAR(10),
ADD COLUMN vitamin_b1_unit VARCHAR(10),
ADD COLUMN vitamin_b2_unit VARCHAR(10),
ADD COLUMN vitamin_b6_unit VARCHAR(10),
ADD COLUMN vitamin_b9_unit VARCHAR(10),
ADD COLUMN vitamin_b12_unit VARCHAR(10),
ADD COLUMN vitamin_c_unit VARCHAR(10),
ADD COLUMN vitamin_d_unit VARCHAR(10),
ADD COLUMN vitamin_e_unit VARCHAR(10),
ADD COLUMN vitamin_pp_unit VARCHAR(10),
ADD COLUMN pantothenic_acid_unit VARCHAR(10),
ADD COLUMN phylloquinone_unit VARCHAR(10),
ADD COLUMN beta_carotene_unit VARCHAR(10),

-- Minerals
ADD COLUMN calcium_unit VARCHAR(10),
ADD COLUMN iron_unit VARCHAR(10),
ADD COLUMN magnesium_unit VARCHAR(10),
ADD COLUMN phosphorus_unit VARCHAR(10),
ADD COLUMN potassium_unit VARCHAR(10),
ADD COLUMN zinc_unit VARCHAR(10),
ADD COLUMN copper_unit VARCHAR(10),
ADD COLUMN manganese_unit VARCHAR(10),
ADD COLUMN selenium_unit VARCHAR(10),
ADD COLUMN iodine_unit VARCHAR(10),

-- Scores & Estimates
ADD COLUMN fruits_veg_legumes_est_unit VARCHAR(10),
ADD COLUMN fruits_veg_nuts_est_unit VARCHAR(10);
