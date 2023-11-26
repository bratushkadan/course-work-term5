-- name: GetProduct :one
SELECT
  p.id, 
  p.store_id,
  p.name,
  p.description,
  p.image_url,
  p.price,
  p.min_height,
  p.max_height,
  pc.id AS "category_id",
  pc.name AS "category_name",
  pc.description AS "category_description",
  p.created
FROM
  "floral"."product" p
JOIN
  "floral"."product_category" pc
ON
  p.category_id = pc.id
WHERE
  p.id = $1;

-- name: AddProduct :one
INSERT INTO
  "floral"."product" (
    store_id,
    name, 
    description,
    image_url,
    price,
    min_height,
    max_height,
    category_id
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateProduct :one
UPDATE
  "floral"."product"
SET
  name = $1,
  description = $2,
  image_url = $3,
  price = $4,
  min_height = $5,
  max_height = $6,
  category_id = $7,
  modified = NOW()
WHERE
  id = $8 AND
  store_id = $9
RETURNING *;

-- name: DeleteProduct :one
DELETE FROM
  "floral"."product"
WHERE
  id = $1 AND
  store_id = $2
RETURNING *;


-- name: GetProducts :many
SELECT
  p.id, 
  p.store_id,
  p.name,
  p.description,
  p.image_url,
  p.price,
  p.min_height,
  p.max_height,
  pc.id AS "category_id",
  pc.name AS "category_name",
  pc.description AS "category_description",
  p.created
FROM
  "floral"."product" p
JOIN
  "floral"."product_category" pc
ON
  p.category_id = pc.id;

-- name: GetStoreProducts :many
SELECT
  p.id, 
  p.store_id,
  p.name,
  p.description,
  p.image_url,
  p.price,
  p.min_height,
  p.max_height,
  pc.id AS "category_id",
  pc.name AS "category_name",
  pc.description AS "category_description",
  p.created
FROM
  "floral"."product" p
JOIN
  "floral"."product_category" pc
ON
  p.category_id = pc.id
WHERE
  p.store_id = $1;

-- name: GetProductsCategories :many
SELECT
  id, name, description
FROM
  "floral"."product_category";

-- name: AddCategory :one
INSERT INTO
  "floral"."product_category" (
    name, description
  )
VALUES
  ($1, $2)
RETURNING *;

-- name: RemoveCategory :one
DELETE FROM
  "floral"."product_category"
WHERE
  id = $1
RETURNING *;