CREATE TABLE "branch"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(50),
    "address" VARCHAR(65),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN DEFAULT false,
    "deleted_at" TIMESTAMP
);