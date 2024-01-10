CREATE TYPE PRODUCT_TYPE AS ENUM ('CARPARTS', 'REGULAR', 'VIRTUAL', 'BUNDLES', 'OTHER');
CREATE TYPE PRODUCT_STATUS AS ENUM ('AVAILABLE', 'OUT_OF_STOCK', 'PREORDER', 'DISCONTINUED', 'DELETED', 'AWAITING_SHIPMENT');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS "brand" (
    "id"                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"                VARCHAR(255),
    "created_at"           TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"           TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "products" (
    "id"                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"              JSONB,
    "slug"              JSONB,
    "ean"               text[],
    "sku"               VARCHAR(255),
    "index"             VARCHAR(255) UNIQUE,
    "status"            VARCHAR(255),
    "description"       JSONB,
    "short_description" JSONB,
    "price"             DOUBLE PRECISION,
    "type"              PRODUCT_TYPE,
    "visibility"        BOOLEAN,
    "weight"            DOUBLE PRECISION,
    "length"            DOUBLE PRECISION,
    "width"             DOUBLE PRECISION,
    "height"            DOUBLE PRECISION,
    "code_ce"           INTEGER,
    "brand_id"          UUID,
    "created_at"         TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"         TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY ("brand_id") REFERENCES "brand"(id)
);

CREATE TABLE IF NOT EXISTS "images" (
    "id"                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "product_id"        UUID NOT NULL,
    "url"               TEXT NOT NULL,
    "created_at"        TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"        TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY ("product_id") REFERENCES "products"(id)
);

CREATE TABLE IF NOT EXISTS "car_parts" (

    "id_car_part"       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "side"              JSONB,
    "ic_index"          VARCHAR(255),
    "tec_doc_art"       VARCHAR(255),
    "tec_doc_brand"     VARCHAR(255),
    "tow_kod"           VARCHAR(255),      
    "history"           JSONB
) INHERITS (products);



CREATE TABLE IF NOT EXISTS "car_type" (
    "id"                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "id_doc"              VARCHAR(255) UNIQUE,
    "cylinders"           INTEGER,
    "cylinders_valves"    INTEGER,
    "power_Kw"            INTEGER,
    "power_Hp"            INTEGER,
    "engine"              VARCHAR(255),
    "code"                VARCHAR(255),
    "displacement"        INTEGER,
    "year_start"          TIMESTAMP,
    "year_end"            TIMESTAMP,
    "created_at"           TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"           TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS OEMNumber (
    "id"                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "part_number"         VARCHAR(255) NOT NULL,
    "brand_product_id"    UUID NOT NULL,
    "product_id"          UUID NOT NULL,
    FOREIGN KEY ("brand_product_id") REFERENCES "brand"(id),
    FOREIGN KEY ("product_id") REFERENCES "car_parts"(id_car_part)
);




ALTER TABLE car_parts
ADD COLUMN "car_type_id" UUID,
ADD FOREIGN KEY ("car_type_id") REFERENCES car_type("id");