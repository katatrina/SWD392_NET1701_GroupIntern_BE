-- name: CreateDentist :one
INSERT INTO users (full_name, hashed_password, email, phone_number, role)
VALUES ($1, $2, $3, $4, 'Dentist') RETURNING *;

-- name: CreateDentistDetail :one
INSERT INTO dentist_detail (dentist_id, date_of_birth, sex, specialty_id)
VALUES ($1, $2, $3, $4) RETURNING *;
