-- name: ListAllServiceCategories :many
SELECT * FROM service_categories;

-- name: ListAllServicesOfACategory :many
SELECT * FROM services WHERE category_id = $1;
