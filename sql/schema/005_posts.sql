-- +goose Up
CREATE TABLE posts (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    url text not null,
    description text not null,
    published_at timestamp not null,
    feed_id uuid not null references feeds(id)
);
-- +goose Down
DROP TABLE posts;