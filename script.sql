SELECT
    COUNT(*) OVER(),
    id,
    sales_id,
    type,
    source_type,
    text,
    amount,
    staff_id,
    created_at,
    updated_at,
    deleted,
    deleted_at
FROM staff_transaction