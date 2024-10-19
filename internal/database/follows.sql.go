// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH follow AS (
  INSERT INTO feed_follows(id,created_at,updated_at,feed_id,user_id)
  VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
  )
  RETURNING id, feed_id, user_id, created_at, updated_at
)
SELECT follow.id, follow.feed_id, follow.user_id, follow.created_at, follow.updated_at, f.name as feed_name, u.name as user_name FROM follow 
JOIN feeds f ON follow.feed_id = f.id
JOIN users u ON follow.user_id = u.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	FeedID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}
