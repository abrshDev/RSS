package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("port not found in the environment")
	}
	Router := chi.NewRouter()
	Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Allow all with https or http
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any browser
	}))
	v1router := chi.NewRouter()
	v1router.Get("/ready", Handlereadiness)
	v1router.Get("/err", HandleErr)
	Router.Mount("/v1", v1router)
	srv := &http.Server{
		Handler: Router,
		Addr:    ":" + portstring,
	}
	log.Printf("server starting on port :%v", portstring)
	fmt.Println("port string:", portstring)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
