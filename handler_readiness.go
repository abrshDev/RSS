package main

import "net/http"

func Handlereadiness(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, 200, struct{}{})
}
