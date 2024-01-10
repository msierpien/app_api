CREATE TYPE user_type AS ENUM ('CUSTOMER', 'SUPPLIER', 'EMPLOYEE', 'SELLER', 'ADMIN');


CREATE TABLE IF NOT EXISTS "users" (

    "id"               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id_fakturownia"   INTEGER,
    "user_type"             "user_type" NOT NULL,
    "first_name"            VARCHAR(100),
    "last_name"             VARCHAR(100),
    "name"                  VARCHAR(255),
    "email"                 VARCHAR(255),
    "phone"                 VARCHAR(20),
    "website"               VARCHAR(255),
    "note"                  TEXT,
    "created_at"            TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"            TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "users_corporate" (

    "id_corporate"  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "company_name"  VARCHAR(255),
    "tax_no"        VARCHAR(20),
    "bank_account"  VARCHAR(30),
    "bank_name"     VARCHAR(100),
    "tax_no_kind"   VARCHAR(50),
    "country"       VARCHAR(50)
)INHERITS (users);

CREATE TABLE IF NOT EXISTS "addresses" (
    "address_id"    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id"       UUID REFERENCES users("id"),
    "street"        VARCHAR(255),
    "post_code"     VARCHAR(10),
    "city"          VARCHAR(100),
    "country"       VARCHAR(50),
    "address_1"     VARCHAR(150),
    "address_2"     VARCHAR(150)
);

ALTER TABLE "product_suppliers"
ADD FOREIGN KEY ("supplier_id") REFERENCES "users_corporate"(id_corporate);