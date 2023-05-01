-- name: CreateUser :one
INSERT INTO users (
  email,
  encrypted_password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
  set encrypted_password = $2
WHERE id = $1
RETURNING *;
