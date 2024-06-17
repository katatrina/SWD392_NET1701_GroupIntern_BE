-- name: CreateService :one
INSERT INTO services (name, category_id, unit, cost, warranty_duration)
VALUES ($1, $2, $3, $4, $5) RETURNING *;
