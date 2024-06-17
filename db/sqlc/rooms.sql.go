// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: rooms.sql

package db

import (
	"context"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO rooms (name)
VALUES ($1) RETURNING id, name, created_at
`

func (q *Queries) CreateRoom(ctx context.Context, name string) (Room, error) {
	row := q.db.QueryRowContext(ctx, createRoom, name)
	var i Room
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
