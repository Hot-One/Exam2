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
    type,
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
WHERE deleted = false AND type = 'Shop Asistent' 
AND created_at BETWEEN 2023-07-25 AND 2023-07-27 
ORDER BY balace desc


SELECT
    b.name as branch,
    SUM(s.price) as total_sum,
    CURRENT_DATE as DAY
FROM branch as b
JOIN sales as s ON s.branch_id = b.id
GROUP BY b.name
ORDER BY total_sum desc

SELECT
        b.name as branch,
        SUM(s.price) as total_sum,
        CURRENT_DATE as DAY
FROM branch as b
JOIN sales as s ON s.branch_id = b.id
WHERE deleted = false 
GROUP BY b.name
ORDER BY total_sum desc


where += `AND name ILIKE '%' || ` + req.Search + ` || '%' OR branch_id ILIKE '%' || ` + req.Search + ` || '%' OR tarif_id ILIKE '%' || ` + req.Search + ` || '%' OR type ILIKE '%' || ` + req.Search + ` || '%'`
