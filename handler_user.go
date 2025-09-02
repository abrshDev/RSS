package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/abrshDev/RSS/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	user, err := apiCfg.DB.CreatedUser(r.Context(), db.CreatedUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("failed when creating a user: %v", err))
		return
	}

	RespondWithJson(w, 200, dbtodb(user))

}
func (apiCfg *apiConfig) handleGetUserByApi(w http.ResponseWriter, r *http.Request, user db.User) {

	RespondWithJson(w, 200, dbtodb(user))
}

func (apiCFG *apiConfig) handleGetPostUser(w http.ResponseWriter, r *http.Request, user db.User) {

	posts, err := apiCFG.DB.GetPostsForUser(r.Context(), db.GetPostsForUserParams{
		UserID: user.ID, Limit: 2})
	if err != nil {
		log.Fatal("couldnot get post for user:", posts)
	}
	RespondWithJson(w, 200, dbslicetodbslice(posts))

}
