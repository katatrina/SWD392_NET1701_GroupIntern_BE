-- name: CreateService :one
INSERT INTO services (name, category_id, unit, cost, warranty_duration)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetService :one
SELECT *
FROM services
WHERE id = $1;

-- name: UpdateService :exec
UPDATE services
SET name              = $2,
    category_id       = $3,
    unit              = $4,
    cost              = $5,
    warranty_duration = $6
WHERE id = $1;

-- name: DeleteService :exec
DELETE FROM services
WHERE id = $1;

-- name: ListServicesByCategory :many
SELECT *
FROM services
WHERE category_id = (SELECT id FROM service_categories WHERE slug = $1)
ORDER BY created_at DESC;

-- name: ListServicesByNameAndCategory :many
SELECT *
FROM services
WHERE name ILIKE '%' || sqlc.arg(name)::text || '%'
AND category_id = (SELECT id FROM service_categories WHERE slug = sqlc.arg(category)::text)
ORDER BY id;
