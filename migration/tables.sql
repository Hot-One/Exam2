CREATE TABLE "branch"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(50),
    "address" VARCHAR(65),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
)

CREATE TABLE staff_tarif(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(35) NOT NULL,
    "type" VARCHAR(35) NOT NULL DEFAULT "fixed",
    "amountForCash" NUMERIC NOT NULL,
    "amountForCard" NUMERIC NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
)

CREATE TABLE "staff"(
    "id" UUID NOT NULL PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
    "tarif_id" UUID REFERENCES "staff_tarif"("id") NOT NULL,
    "type" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    "balace" NUMERIC NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
)

CREATE TABLE "staff_transaction"(
    "id" UUID NOT NULL PRIMARY KEY,
    "sales_id" UUID REFERENCES sales("id") NOT NULL,
    "type" VARCHAR NOT NULL,
    "source_type" VARCHAR NOT NULL,
    "text" TEXT NOT NULL,
    "amount" NUMERIC NOT NULL
    "staff_id" UUID REFERENCES staff("id") NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
)

CREATE TABLE sales(
    "id" UUID NOT NULL PRIMARY KEY,
    "branch_id" UUID REFERENCES branch("id") NOT NULL,
    "shop_assistent_id" UUID REFERENCES staff("id"),
    "cashier_id" UUID REFERENCES staff("id") NOT NULL,
    "price" NUMERIC NOT NULL,
    "payment_type" VARCHAR NOT NULL,
    "status" VARCHAR NOT NULL DEFAULT "success",
    "client_name" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
);


