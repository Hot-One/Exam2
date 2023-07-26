ALTER TABLE staff_transaction
ADD COLUMN caisher_id UUID REFERENCES staff("id") NOT NULL;