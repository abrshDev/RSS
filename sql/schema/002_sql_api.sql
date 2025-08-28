-- +goose Up

CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER TABLE users
ADD COLUMN api_key uuid UNIQUE NOT NULL DEFAULT gen_random_uuid();

-- +goose Down

ALTER TABLE users DROP COLUMN api_key;
