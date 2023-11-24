-- https://docs.sqlc.dev/en/stable/howto/query_count.html

-- name: GetStore :one
SELECT
  id, name, email, phone_number, created
FROM
  "floral"."store"
WHERE
  id = $1;

-- name: CountStoreByCreds :one
SELECT
  count(id) as "count"
FROM
  "floral"."store"
WHERE
  email = $1 AND
  password = $2;

-- name: CreateStore :one
INSERT INTO
  "floral"."store" (name, email, password, phone_number)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING id, name, email, phone_number, created;

-- name: DeleteStore :one
DELETE FROM
  "floral"."store"
WHERE
  id = $1 AND password = $2;
