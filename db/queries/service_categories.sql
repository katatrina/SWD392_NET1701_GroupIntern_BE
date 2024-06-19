-- name: ListServiceCategories :many
SELECT * FROM service_categories;

-- name: ListServicesOfOneCategory :many
SELECT * FROM services WHERE category_id = (SELECT id FROM service_categories WHERE slug = $1);

-- name: CreateServiceCategory :one
INSERT INTO service_categories (name, icon_url, banner_url, slug, description)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetServiceCategoryBySlug :one
SELECT * FROM service_categories WHERE slug = $1;
