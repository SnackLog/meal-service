CREATE TABLE meal_entries (
    id SERIAL PRIMARY KEY,
    meal_id INTEGER NOT NULL REFERENCES meals(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    quantity DOUBLE PRECISION NOT NULL,
    quantity_unit VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Energy
    energy_kcal             DOUBLE PRECISION,
    energy_kcal_unit        VARCHAR(10),
    energy_kj               DOUBLE PRECISION,
    energy_kj_unit          VARCHAR(10),

    -- Macros: Fats
    fat                     DOUBLE PRECISION,
    fat_unit                VARCHAR(10),
    saturated_fat           DOUBLE PRECISION,
    saturated_fat_unit      VARCHAR(10),
    monounsaturated_fat     DOUBLE PRECISION,
    monounsaturated_fat_unit    VARCHAR(10),
    polyunsaturated_fat     DOUBLE PRECISION,
    polyunsaturated_fat_unit    VARCHAR(10),
    trans_fat               DOUBLE PRECISION,
    trans_fat_unit          VARCHAR(10),
    cholesterol             DOUBLE PRECISION,

    -- Macros: Carbohydrates & Sugars
    carbohydrates           DOUBLE PRECISION,
    carbohydrates_unit      VARCHAR(10),
    sugars                  DOUBLE PRECISION,
    sucrose                 DOUBLE PRECISION,
    glucose                 DOUBLE PRECISION,
    fructose                DOUBLE PRECISION,
    lactose                 DOUBLE PRECISION,
    maltose                 DOUBLE PRECISION,
    galactose               DOUBLE PRECISION,
    starch                  DOUBLE PRECISION,
    polyols                 DOUBLE PRECISION,
    fiber                   DOUBLE PRECISION,

    -- Macros: Proteins & Others
    proteins                DOUBLE PRECISION,
    proteins_unit           VARCHAR(10),
    alcohol                 DOUBLE PRECISION,
    water                   DOUBLE PRECISION,

    -- Salt & Sodium
    salt                    DOUBLE PRECISION,
    salt_unit               VARCHAR(10),
    sodium                  DOUBLE PRECISION,
    sodium_unit             VARCHAR(10),

    -- Vitamins
    vitamin_a               DOUBLE PRECISION,
    vitamin_b1              DOUBLE PRECISION,
    vitamin_b2              DOUBLE PRECISION,
    vitamin_b6              DOUBLE PRECISION,
    vitamin_b9              DOUBLE PRECISION,
    vitamin_b12             DOUBLE PRECISION,
    vitamin_c               DOUBLE PRECISION,
    vitamin_d               DOUBLE PRECISION,
    vitamin_e               DOUBLE PRECISION,
    vitamin_pp              DOUBLE PRECISION,
    pantothenic_acid        DOUBLE PRECISION,
    phylloquinone           DOUBLE PRECISION,
    beta_carotene           DOUBLE PRECISION,

    -- Minerals
    calcium                 DOUBLE PRECISION,
    iron                    DOUBLE PRECISION,
    magnesium               DOUBLE PRECISION,
    phosphorus              DOUBLE PRECISION,
    potassium               DOUBLE PRECISION,
    zinc                    DOUBLE PRECISION,
    copper                  DOUBLE PRECISION,
    manganese               DOUBLE PRECISION,
    selenium                DOUBLE PRECISION,
    iodine                  DOUBLE PRECISION,

    -- Scores & Estimates
    nova_group              SMALLINT,
    nutrition_score_fr      SMALLINT,
    fruits_veg_legumes_est  DOUBLE PRECISION,
    fruits_veg_nuts_est     DOUBLE PRECISION
);

CREATE INDEX idx_meal_entries_meal_id ON meal_entries(meal_id);
