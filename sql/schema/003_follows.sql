-- +goose Up
CREATE TABLE feed_follows (
    id uuid primary key,
    feed_id uuid not null references feeds(id) on delete cascade,
    user_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null,
    updated_at timestamp not null,
    unique (feed_id, user_id)
);
-- +goose Down
DROP TABLE feed_follows;