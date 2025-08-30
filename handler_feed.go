package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/abrshDev/RSS/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	fmt.Println("user id", user.ID)
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	feed, err := apiCfg.DB.CreatedFeed(r.Context(), db.CreatedFeedParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Name:          params.Name,
		Url:           params.Url,
		UserID:        user.ID,
		Lastfetchedat: sql.NullTime{},
	})
	fmt.Println("feed name:", params.Name)
	fmt.Println("feed url:", params.Url)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("failed when creating a user: %v", err))
		return
	}

	RespondWithJson(w, 200, feedtofeed(feed))

}

func (apicfg *apiConfig) HandleGetFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintf("couldnot get feeds :%v", err))
	}
	RespondWithJson(w, 200, feedstofeeds(feeds))

}
