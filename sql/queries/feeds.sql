-- name: CreatedFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id,lastfetchedat)
VALUES ($1, $2, $3, $4, $5,$6,$7)
RETURNING id, created_at, updated_at, name, url,user_id,lastfetchedat;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
ORDER BY lastfetchedat ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET lastfetchedat = NOW()
WHERE id = $1
RETURNING id, name, url, lastfetchedat;


