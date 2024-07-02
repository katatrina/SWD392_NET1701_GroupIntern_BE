-- name: ListServiceCategories :many
SELECT *
FROM service_categories
ORDER BY created_at DESC;

-- name: CreateServiceCategory :one
INSERT INTO service_categories (name, icon_url, banner_url, slug, description)
VALUES (sqlc.arg(name)::text, sqlc.arg(icon_url), sqlc.arg(banner_url), sqlc.arg(slug),
        sqlc.arg(description)) RETURNING *;

-- name: GetServiceCategoryBySlug :one
SELECT *
FROM service_categories
WHERE slug = $1;

-- name: GetServiceCategoryByID :one
SELECT id,
       name::text,
       icon_url,
       banner_url,
       description,
       slug,
       created_at
FROM service_categories
WHERE id = $1;

-- name: UpdateServiceCategory :exec
UPDATE service_categories
SET name        = sqlc.arg(name)::text,
    icon_url    = sqlc.arg(icon_url),
    banner_url  = sqlc.arg(banner_url),
    slug        = sqlc.arg(slug),
    description = sqlc.arg(description)
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
