-- +goose Up
CREATE TABLE users (
    id uuid primary key,
    name text unique not null,
    created_at timestamp not null,
    updated_at timestamp not null
);
-- +goose Down
DROP TABLE users;