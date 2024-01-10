
CREATE TABLE IF NOT EXISTS "product_suppliers" (
    "id"                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "product_id"        UUID,
    "car_part_id"       UUID,
    "supplier_id"       UUID,
    "index_suppliers"     VARCHAR(25) NOT NULL,
    "created_at"        TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"        TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY ("product_id") REFERENCES "products"(id),
    FOREIGN KEY ("car_part_id") REFERENCES "car_parts"(id_car_part)
);

CREATE TABLE IF NOT EXISTS "product_prices" (
    "id"                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "product_suppliers_id"        UUID NOT NULL,
    "price"             DOUBLE PRECISION NOT NULL,
    "effective_date"    TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY ("product_suppliers_id") REFERENCES "product_suppliers"("id")
);