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
RETURNING *;

-- name: RemoveCartPosition :one
DELETE FROM
  "floral"."cart_position"
WHERE
  user_id = $1 AND
  product_id = $2
RETURNING *;

-- name: ClearCart :many
DELETE FROM
  "floral"."cart_position"
WHERE
  user_id = $1
RETURNING *;
