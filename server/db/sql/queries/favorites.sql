-- name: GetFavorites :many
SELECT
  p.id, 
  p.name,
  p.description,
  p.image_url,
  p.price,
  pc.id AS "category_id",
  pc.name AS "category_name",
  p.store_id,
  s.name AS "store_name",
  uf.created AS "added_favorite"
FROM
  "floral"."user_favorite" uf
JOIN "floral"."product" p ON uf.product_id =  p.id
JOIN "floral"."product_category" pc ON p.category_id = pc.id
JOIN "floral"."store" s ON p.store_id = s.id
WHERE
  uf.user_id = $1;

-- name: AddFavorite :one
INSERT INTO
  "floral"."user_favorite" (user_id, product_id)
VALUES
  ($1, $2)
RETURNING product_id;

-- name: DeleteFavorite :one
DELETE FROM
  "floral"."user_favorite"
WHERE
  user_id = $1 AND
  product_id = $2
RETURNING product_id;
