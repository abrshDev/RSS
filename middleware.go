package main

import (
	"fmt"
	"net/http"

	"github.com/abrshDev/RSS/internal/auth"
	"github.com/google/uuid"

	db "github.com/abrshDev/RSS/internal/database"
)

type authhandler func(w http.ResponseWriter, r *http.Request, user db.User)

func (apiCfg *apiConfig) Middleware(handler authhandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r.Header)
		fmt.Println("api_key:", api_key)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error :%v", err))
			return
		}
		parsedkey, _ := uuid.Parse(api_key)

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), parsedkey)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("couldnot get user  :%v", err))
			return
		}

		fmt.Println("user:", user)
		handler(w, r, user)
	}
}
