-- name: CreatePost :one
INSERT INTO posts(id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES(
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
-- name: GetPostsForUser :many
SELECT p.id,
       p.title,
       p.url,
       p.description
FROM posts p
JOIN feed_follows f ON p.feed_id = f.feed_id
WHERE f.user_id = $1
order by p.published_at desc
limit $2;