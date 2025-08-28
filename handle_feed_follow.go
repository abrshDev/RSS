package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/abrshDev/RSS/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	fmt.Println("user id", user.ID)
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	feedfollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("failed when creating a user: %v", err))
		return
	}

	RespondWithJson(w, 200, feedfollowtofeedfollow(feedfollow))

}
func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	feedfollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("failed getting feeds follows: %v", err))
		return
	}

	RespondWithJson(w, 200, dbfeedfollowtofeedfollow(feedfollows))

}

func (apicfg *apiConfig) HandleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedfollowidstring := chi.URLParam(r, "feedid")

	feedfollowid, err := uuid.Parse(feedfollowidstring)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("canot parse feedfollowid:", err))
		return
	}
	err = apicfg.DB.DeleteFeedFollows(r.Context(), db.DeleteFeedFollowsParams{
		FeedID: feedfollowid,
		UserID: user.ID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("canpot delete data:", err))
		return
	}
	RespondWithJson(w, 200, struct{}{})

}
