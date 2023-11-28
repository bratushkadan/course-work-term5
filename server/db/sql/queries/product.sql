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
WITH p AS (
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
)
SELECT
  *
FROM p
WHERE
  (@lk_name::bool = false OR LOWER(p.name) LIKE ('%' || LOWER(@name) || '%'))
  AND (CASE WHEN @is_store_id::bool THEN p.store_id = @store_id ELSE TRUE END)
  AND (CASE WHEN @is_category_id::bool THEN p.category_id = @category_id ELSE TRUE END)
  AND (CASE WHEN @is_min_height::bool THEN p.min_height >= @min_height ELSE TRUE END)
  AND (CASE WHEN @is_max_height::bool THEN p.max_height >= @max_height ELSE TRUE END)
  AND (CASE WHEN @is_min_price::bool THEN p.price >= @min_price ELSE TRUE END)
  AND (CASE WHEN @is_max_price::bool THEN p.price <= @max_price ELSE TRUE END)
ORDER BY
  CASE WHEN @id_asc::bool THEN id END ASC,
  CASE WHEN @id_desc::bool THEN id END DESC,
  CASE WHEN @max_height_asc::bool THEN id END ASC,
  CASE WHEN @max_height_desc::bool THEN id END DESC,
  CASE WHEN @price_asc::bool THEN id END ASC,
  CASE WHEN @price_desc::bool THEN id END DESC,
  CASE WHEN @store_id_asc::bool THEN id END ASC,
  CASE WHEN @store_id_desc::bool THEN id END DESC;

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

-- name: GetProductsCategory :one
SELECT
  id, name, description
FROM
  "floral"."product_category"
WHERE
  id = $1;

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