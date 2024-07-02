-- name: ListServiceCategories :many
SELECT *
FROM service_categories
ORDER BY created_at DESC;

-- name: CreateServiceCategory :one
INSERT INTO service_categories (name, icon_url, banner_url, slug, description)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetServiceCategoryBySlug :one
SELECT *
FROM service_categories
WHERE slug = $1;

-- name: GetServiceCategoryByID :one
SELECT *
FROM service_categories
WHERE id = $1;

-- name: UpdateServiceCategory :exec
UPDATE service_categories
SET name        = $2,
    icon_url    = $3,
    banner_url  = $4,
    slug        = $5,
    description = $6
WHERE id = $1;

-- name: DeleteServiceCategory :exec
DELETE
FROM service_categories
WHERE id = $1;

-- name: ListServiceCategoriesByName :many
SELECT *
FROM service_categories
WHERE name ILIKE '%' || sqlc.arg(name)::text || '%'
ORDER BY created_at DESC;
