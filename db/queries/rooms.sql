-- name: CreateRoom :one
INSERT INTO rooms (name)
VALUES ($1) RETURNING *;

-- name: ListRooms :many
SELECT *
FROM rooms
ORDER BY created_at DESC;

-- name: UpdateRoom :exec
UPDATE rooms
SET name = $2
WHERE id = $1;
