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
