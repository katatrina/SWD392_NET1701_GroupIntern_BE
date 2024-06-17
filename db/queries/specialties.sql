-- name: CreateSpecialty :one
INSERT INTO specialties (name)
VALUES ($1) RETURNING *;
