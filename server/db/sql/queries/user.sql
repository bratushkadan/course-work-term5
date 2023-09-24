-- name: GetUsers :many
SELECT
  *
FROM
  "floral"."user";

-- name: GetUser :one
SELECT
  *
FROM
  "floral"."user"
WHERE
  id = $1;