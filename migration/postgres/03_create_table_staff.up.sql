CREATE TABLE "staff"(
    "id" UUID NOT NULL PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
    "tarif_id" UUID REFERENCES "staff_tarif"("id") NOT NULL,
    "type" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    "balace" INT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
);