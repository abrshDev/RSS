package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	db "github.com/abrshDev/RSS/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *db.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %d goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("canot fetch next feed to follow :", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}
func scrapeFeed(dbs *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	defer wg.Done()

	_, err := dbs.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("failed to mark feed that is fetched", err)
		return
	}
	rssfeed, err := urlToBeFetched(feed.Url)
	if err != nil {
		fmt.Println("failed to fetch url", err)
		return
	}
	for _, item := range rssfeed.Channel.Item {
		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		_, err := dbs.CreatePost(context.Background(), db.CreatePostParams{
			ID:           uuid.New(),
			CreatedAt:    time.Now().UTC(),
			UpdatedAt:    time.Now().UTC(),
			Title:        item.Title,
			Descrtiption: description,
			PublishedAt:  item.PubDate,
			Url:          item.Link,
			FeedID:       feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Failed to create post:", err)
			fmt.Println("just tring smt")
		}
	}
	log.Printf("log %s collected, %v posts found", feed.Name, len(rssfeed.Channel.Item))

}
