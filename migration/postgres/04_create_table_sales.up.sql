

CREATE TABLE sales(
    "id" UUID NOT NULL PRIMARY KEY,
    "branch_id" UUID REFERENCES branch("id") NOT NULL,
    "shop_assistent_id" UUID REFERENCES staff("id"),
    "cashier_id" UUID REFERENCES staff("id") NOT NULL,
    "price" INT NOT NULL,
    "payment_type" VARCHAR NOT NULL,
    "status" VARCHAR NOT NULL DEFAULT 'success',
    "client_name" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
);
