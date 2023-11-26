-- name: GetOrder :one
SELECT
  id, user_id, status, created, status_modified
FROM
  "floral"."order"
WHERE id = $1;

-- name: AddOrder :one
INSERT INTO
  "floral"."order" (user_id, status)
VALUES
  ($1, $2)
RETURNING id, user_id, status, created, status_modified;

-- name: UpdateOrderStatus :one
UPDATE
  "floral"."order"
SET
  status = $1,
  status_modified = NOW()
WHERE
  id = $2
RETURNING id, user_id, status, created, status_modified;

-- name: GetUserOrders :many
SELECT
  id, user_id, status, created, status_modified
FROM
  "floral"."order"
WHERE
  user_id = $1;

-- name: GetUserOrderPositions :many
SELECT
  op.quantity,
  o.id AS "order_id",
  p.id AS "product_id",
  p.name AS "product_name",
  p.image_url AS "product_image_url",
  p.price AS "product_price"
FROM
  "floral"."order_position" op
JOIN
  "floral"."order" o
ON
  op.order_id = o.id
JOIN
  "floral"."product" p
ON
  op.product_id = p.id
WHERE
  order_id = $1;

-- name: AddUserOrderPositions :copyfrom
INSERT INTO
  "floral"."order_position" (order_id, product_id, quantity)
VALUES
  ($1, $2, $3);