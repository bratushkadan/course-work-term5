-- https://docs.sqlc.dev/en/stable/howto/query_count.html

-- name: GetUsers :many
SELECT
  id, first_name, last_name, email, phone_number, created
FROM
  "floral"."user";

-- name: GetUser :one
SELECT
  id, first_name, last_name, email, phone_number, created
FROM
  "floral"."user"
WHERE
  id = $1;

-- name: GetUserCreds :one
SELECT
  id, password
FROM
  "floral"."user"
WHERE
  email = $1;

-- name: CreateUser :one
INSERT INTO
  "floral"."user" (first_name, last_name, email, password, phone_number)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, email, phone_number, created;

-- name: UpdateUserPassword :one
UPDATE
  "floral"."user"
SET
  password = $1
WHERE id = $2
RETURNING id, first_name, last_name, email, phone_number, created;
