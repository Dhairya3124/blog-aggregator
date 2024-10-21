-- name: CreateFeedFollow :one
WITH follow AS (
  INSERT INTO feed_follows(id,created_at,updated_at,feed_id,user_id)
  VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
  )
  RETURNING *
)
SELECT follow.*, f.name as feed_name, u.name as user_name FROM follow 
JOIN feeds f ON follow.feed_id = f.id
JOIN users u ON follow.user_id = u.id;

-- name: GetFollowsForUser :many
select 
  ff.id,
  f.name
from feed_follows ff
join feeds f on ff.feed_id = f.id
where ff.user_id = $1;