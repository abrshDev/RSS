-- +goose Up 
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    descrtiption TEXT,
    published_at TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    feed_id UUID NOT NULL  REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE posts ;
