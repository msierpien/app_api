-- Usuwamy klucz obcy i kolumnę z 'car_parts'
ALTER TABLE car_parts DROP COLUMN IF EXISTS "car_type_id";

-- Usuwamy tabelę 'car_type'
DROP TABLE IF EXISTS car_type;

-- Usuwamy tabelę 'car_parts'
DROP TABLE IF EXISTS car_parts;

-- Usuwamy tabelę 'images'
DROP TABLE IF EXISTS images;

-- Usuwamy tabelę 'products'
DROP TABLE IF EXISTS products;