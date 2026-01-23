TRUNCATE TABLE meals CASCADE;
ALTER SEQUENCE meals_id_seq RESTART WITH 3;

-- Seed data for 'meals'
INSERT INTO meals (id, name, timestamp, username) VALUES
(1, 'Quinoa Power Bowl', '2026-01-23 12:30:00', 'foo'),
(2, 'Salmon Dinner', '2026-01-23 19:00:00', 'bar');

-- Seed data for 'meal_entries'

-- Entries for Meal 1: Quinoa Power Bowl
INSERT INTO meal_entries (
    meal_id, name, quantity, quantity_unit,
    energy_kcal, energy_kj,
    fat, saturated_fat, monounsaturated_fat,
    carbohydrates, sugars, fiber,
    proteins, salt,
    vitamin_c, iron, magnesium
) VALUES (
    1, 'Cooked Quinoa', 185, 'g',
    222, 929,
    3.6, 0.4, 1.0,
    39.4, 1.6, 5.2,
    8.1, 0.01,
    0, 2.8, 118
), (
    1, 'Avocado', 50, 'g',
    80, 335,
    7.3, 1.0, 4.9,
    4.3, 0.3, 3.4,
    1.0, 0.005,
    5.0, 0.3, 14.5
);

-- Entries for Meal 2: Salmon Dinner
INSERT INTO meal_entries (
    meal_id, name, quantity, quantity_unit,
    energy_kcal,
    fat, saturated_fat, polyunsaturated_fat, cholesterol,
    proteins, sodium,
    vitamin_d, vitamin_b12, selenium, potassium,
    nova_group, nutrition_score_fr
) VALUES (
    2, 'Atlantic Salmon (Grilled)', 170, 'g',
    350,
    22.0, 4.5, 8.2, 94.0,
    38.0, 105,
    15.0, 4.8, 55.0, 620,
    1, 1
), (
    2, 'Steamed Broccoli', 150, 'g',
    50,
    0.5, 0.1, 0.0, 0.0,
    4.0, 50,
    0.0, 0.0, 4.0, 450,
    1, 1
), (
    2, 'Extra Virgin Olive Oil', 10, 'ml',
    88,
    10.0, 1.4, 0.8, 0.0,
    0.0, 0,
    0.0, 0.0, 0.0, 0,
    2, 2
);
