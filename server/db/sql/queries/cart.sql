-- name: GetCart :many
SELECT
  c.product_id,
  c.quantity,
  p.name,
  p.description,
  p.price,
  p.image_url,
  p.category_id,
  pc.name AS "category_name"
FROM
  "floral"."cart_position" c
JOIN
  "floral"."product" p
ON
  c.product_id = p.id
JOIN
  "floral"."product_category" pc
ON
  p.category_id = pc.id
WHERE
  c.user_id = $1;
  

-- name: UpsertCartPosition :one
INSERT INTO
  "floral"."cart_position" (
    user_id,
    product_id,
    quantity
  )
VALUES
  ($1, $2, $3)
ON CONFLICT (user_id, product_id)
DO UPDATE
SET
  quantity = $3
WHERE
  "floral"."cart_position".user_id = $1 AND
  "floral"."cart_position".product_id = $2
RETURNING product_id, quantity;

-- name: RemoveCartPosition :one
DELETE FROM
  "floral"."cart_position"
WHERE
  user_id = $1 AND
  product_id = $2
RETURNING product_id, quantity;

-- name: ClearCart :many
DELETE FROM
  "floral"."cart_position"
WHERE
  user_id = $1
RETURNING product_id, quantity;
