-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, name
FROM users
WHERE name ILIKE $1
LIMIT 1;

-- name: QueryUsers :many
SELECT * FROM users;


-- name: DeleteUsers :exec
DELETE FROM users;