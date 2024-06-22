-- name: CreateSpecialty :one
INSERT INTO specialties (name)
VALUES ($1) RETURNING *;

-- name: GetSpecialty :one
SELECT *
FROM specialties
WHERE id = $1;
