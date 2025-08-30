-- +goose Up
ALTER TABLE feeds ADD COLUMN lastfetchedat TIMESTAMP ;

-- +goose Down
ALTER TABLE feeds DROP COLUMN lastfetchedat;