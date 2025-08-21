package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Fatal("responsing with 5xx error :", msg)
	}
	type errResponse struct {
		Error string `json:"error`
	}
	RespondWithJson(w, code, errResponse{
		Error: "something went wrong",
	})
}
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal:", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
