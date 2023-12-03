-- name: GetOrder :one
SELECT
  id, user_id, status, created, status_modified
FROM
  "floral"."order"
WHERE
  id = $1
  AND user_id = $2;

-- name: AddOrder :one
INSERT INTO
  "floral"."order" (user_id, status)
VALUES
  ($1, 'created')
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

-- name: GetOrders :many
SELECT
  id,
  user_id,
  status,
  created,
  status_modified
FROM
  "floral"."order"
WHERE
  user_id = $1
ORDER BY created DESC;

-- name: GetOrderPositions :many
SELECT
  p.id,
  op.quantity,
  p.name,
  p.description,
  p.price,
  p.image_url,
  p.category_id,
  c.name AS "category_name",
  s.id AS "store_id",
  s.name AS "store_name"
FROM
  "floral"."order_position" op
JOIN "floral"."order" o ON op.order_id = o.id
JOIN "floral"."product" p ON op.product_id = p.id
JOIN "floral"."store" s ON p.store_id = s.id
JOIN "floral"."product_category" c ON p.category_id = c.id
WHERE
  order_id = $1
  AND user_id = $2;

-- name: GetUserPurchasedProduct :one
SELECT
  COUNT(user_id) > 0 AS "purchased"
FROM 
  "floral"."order" o
JOIN
  "floral"."order_position" op
ON
  op.order_id = o.id
WHERE
  o.user_id = $1
  AND product_id = $2
  AND o.status = 'completed';

-- name: AddOrderPositions :copyfrom
INSERT INTO
  "floral"."order_position" (order_id, product_id, quantity)
VALUES
  ($1, $2, $3);