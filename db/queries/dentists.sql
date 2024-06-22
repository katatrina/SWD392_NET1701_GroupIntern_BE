-- name: CreateDentistDetail :one
INSERT INTO dentist_detail (dentist_id, date_of_birth, gender, specialty_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: ListDentists :many
SELECT users.id,
       users.full_name,
       users.email,
       users.phone_number,
       users.created_at,
       dentist_detail.date_of_birth,
       dentist_detail.gender,
       specialties.name AS specialty
FROM users
         JOIN dentist_detail ON users.id = dentist_detail.dentist_id
         JOIN specialties ON dentist_detail.specialty_id = specialties.id
WHERE users.role = 'Dentist'
ORDER BY users.created_at DESC;

-- name: ListDentistsByName :many
SELECT users.id,
       users.full_name,
       users.email,
       users.phone_number,
       users.created_at,
       dentist_detail.date_of_birth,
       dentist_detail.gender,
       specialties.name AS specialty
FROM users
         JOIN dentist_detail ON users.id = dentist_detail.dentist_id
         JOIN specialties ON dentist_detail.specialty_id = specialties.id
WHERE users.role = 'Dentist'
  AND users.full_name ILIKE '%' || sqlc.arg(name)::text || '%'
ORDER BY users.created_at DESC;

-- name: GetDentist :one
SELECT users.id,
       users.full_name,
       users.email,
       users.phone_number,
       users.created_at,
       dentist_detail.date_of_birth,
       dentist_detail.gender,
       dentist_detail.specialty_id,
       specialties.name AS specialty
FROM users
         JOIN dentist_detail ON users.id = dentist_detail.dentist_id
         JOIN specialties ON dentist_detail.specialty_id = specialties.id
WHERE users.id = $1
  AND users.role = 'Dentist';

