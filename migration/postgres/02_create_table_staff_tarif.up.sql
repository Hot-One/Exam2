CREATE TABLE staff_tarif(
    "id" UUID NOT NULL PRIMARY KEY, 
    "name" VARCHAR(35) NOT NULL, 
    "type" VARCHAR(35) NOT NULl DEFAULT 'fixed', 
    "amountforcash" INT NOT NULL, 
    "amountforcard" INT NOT NULL, 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    "updated_at" TIMESTAMP, 
    "deleted" BOOLEAN DEFAULT false, 
    "deleted_at" TIMESTAMP 
);