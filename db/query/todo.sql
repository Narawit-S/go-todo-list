-- name: CreateTodo :one
INSERT INTO todos (
  user_id,
  title,
  finished,
  deadline
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: ListTodos :many
SELECT * FROM todos
WHERE user_id = $1
ORDER BY id;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET
  title = $2,
  finished = $3
WHERE id = $1
RETURNING *;
