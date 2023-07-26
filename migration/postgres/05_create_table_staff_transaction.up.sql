CREATE TABLE "staff_transaction"(
    "id" UUID NOT NULL PRIMARY KEY,
    "sales_id" UUID REFERENCES sales("id") NOT NULL,
    "type" VARCHAR NOT NULL,
    "source_type" VARCHAR NOT NULL,
    "text" TEXT NOT NULL,
    "amount" INT NOT NULL,
    "staff_id" UUID REFERENCES staff("id") NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
);

