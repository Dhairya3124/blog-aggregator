-- name: CreateFeed :one
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES(
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeed :one
select * from feeds where name = $1;

-- name: GetFeeds :many
select f.name,f.url,u.name from feeds f join users u on f.user_id = u.id;

-- name: GetFeedByURL :one
select * from feeds where url = $1;