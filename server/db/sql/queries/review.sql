-- name: GetProductReviews :many
SELECT
  r.id,
  r.user_id,
  r.product_id,
  r.rating,
  r.review_text,
  u.first_name AS "user_first_name",
  u.last_name AS "user_last_name"
FROM
  "floral"."review" r
JOIN
  "floral"."user" u
ON
  r.user_id = u.id
WHERE r.product_id = $1;

-- name: AddProductReview :one
INSERT INTO
  "floral"."review" (rating, review_text, modified)
VALUES
  ($1, $2, NOW())
RETURNING *;

-- name: UpdateProductReview :one
UPDATE
  "floral"."review"
SET
  rating = $1,
  review_text = $2,
  modified = NOW()
WHERE
  id = $3
RETURNING *;

-- name: DeleteProductReview :one
DELETE FROM
  "floral"."review"
WHERE
  id = $1
RETURNING *;

