-- name: GetUserFavorites :one
SELECT
  p.id, 
  p.store_id,
  p.name,
  p.description,
  p.image_url,
  p.price,
  pc.id AS "category_id",
  pc.name AS "category_name",
  uf.created AS "added_favorite"
FROM
  "floral"."user_favorite" uf
JOIN
  "floral"."product" p
ON
  uf.product_id =  p.id
JOIN
  "floral"."product_category" pc
ON
  p.category_id = pc.id
WHERE
  uf.user_id = $1;

-- name: AddUserFavoriteProduct :one
INSERT INTO
  "floral"."user_favorite" (user_id, product_id)
VALUES
  ($1, $2)
RETURNING product_id;

-- name: RemoveUserFavoriteProduct :one
DELETE FROM
  "floral"."user_favorite"
WHERE
  user_id = $1 AND
  product_id = $2
RETURNING product_id;
