SELECT
    COUNT(*) OVER(),
    id,
    branch_id,
    tarif_id,
    type,
    name,
    balace,
    created_at,
    updated_at,
    deleted,
    deleted_at
FROM staff
WHERE 
    balace BETWEEN 0 AND 10000000;

SELECT
    name,
    branch_id,
    balace
FROM
    staff
ORDER BY balace desc;


SELECT
        name,
        branch_id,
        balace
FROM
        staff
WHERE deleted = false
ORDER BY desc