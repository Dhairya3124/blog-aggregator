-- +goose Up
CREATE TABLE feeds (
    id uuid primary key,
    name text not null,
    url text unique not null,
    user_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null,
    updated_at timestamp not null
);
-- +goose Down
DROP TABLE feeds;