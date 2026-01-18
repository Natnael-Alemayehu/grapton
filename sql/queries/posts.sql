-- name: CreatePost :one
INSERT INTO posts(
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostForUser :many
SELECT 
    p.id,
    p.created_at,
    p.updated_at,
    p.title AS post_title,
    p.url AS post_url,
    p.description,
    p.published_at,
    p.feed_id,
    f.name AS feeds_name,
    f.url AS feed_url,
    f.user_id AS user_id,
    u.name AS user_name
FROM posts p
JOIN feeds f ON f.id = p.feed_id
JOIN users u ON u.id = f.user_id
WHERE f.user_id = $1
ORDER BY p.updated_at DESC
LIMIT $2;