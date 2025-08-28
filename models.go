package main

import (
	"time"

	db "github.com/abrshDev/RSS/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    uuid.UUID `json:"apikey"`
}

func dbtodb(dbuser db.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKey:    dbuser.ApiKey,
	}
}

type Feed struct {
	ID uuid.UUID `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userid"`
}

func feedtofeed(feed db.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}
func feedstofeeds(feed []db.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range feed {
		feeds = append(feeds, feedtofeed(dbfeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func feedfollowtofeedfollow(feedfollo db.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedfollo.ID,
		CreatedAt: feedfollo.CreatedAt,
		UpdatedAt: feedfollo.UpdatedAt,
		UserID:    feedfollo.UserID,
		FeedID:    feedfollo.FeedID,
	}
}
func dbfeedfollowtofeedfollow(feedfollow []db.FeedFollow) []FeedFollow {
	feedfollows := []FeedFollow{}
	for _, dbfeedfollow := range feedfollow {
		feedfollows = append(feedfollows, feedfollowtofeedfollow(dbfeedfollow))
	}
	return feedfollows
}
