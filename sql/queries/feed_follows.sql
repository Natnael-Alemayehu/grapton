-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS(
    INSERT INTO feed_follows(
        id,
        created_at,
        updated_at,
        user_id,
        feed_id
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
RETURNING *) SELECT inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;


-- name: GetFeedFollowsForUser :many
SELECT
    ff.id AS ff_id, 
    ff.created_at AS ff_created_at,
    ff.updated_at AS ff_updated_at,
    u.id AS user_id,
    u.name AS feed_follower_name,
    f.name AS feed_name,
    f.url AS feed_url,
    f.user_id AS feed_created_by
FROM feed_follows ff
JOIN users u ON ff.user_id = u.id
JOIN feeds f ON ff.feed_id = f.id
WHERE ff.user_id = $1;


-- name: UnfollowFeed :exec
DELETE FROM feed_follows ff
WHERE ff.feed_id = $1 AND ff.user_id=$2;