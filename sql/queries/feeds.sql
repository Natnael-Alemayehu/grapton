-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: ListFeed :many
SELECT * from feeds;

-- name: FeedDetail :many
SELECT 
    f.id AS feed_id,
    f.created_at AS feed_created_at, 
    f.updated_at AS feed_updated_at, 
    f.name AS feed_name, 
    f.url AS feed_url,
    u.name AS user_name,
    u.id AS user_id

FROM feeds f
JOIN users u ON u.id = f.user_id;

-- name: GetFeedByName :one
SELECT * FROM feeds WHERE name ILIKE $1 LIMIT 1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url LIKE $1 LIMIT 1;

-- name: MarkFeedFetched :one 
UPDATE feeds SET 
updated_at = now(),
last_fetched_at = now()
WHERE id= $1
RETURNING *;


-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY 
    last_fetched_at NULLS FIRST,
    last_fetched_at ASC
LIMIT 1;