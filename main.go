package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	db "github.com/abrshDev/RSS/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	/* 	feed, err := urlToBeFetched("https://www.wagslane.dev/index.xml")

	   	if err != nil {
	   		fmt.Println("error:", err)
	   	}
	   	fmt.Println(feed) */
	godotenv.Load()
	portstring := os.Getenv("PORT")
	dburl := os.Getenv("DB_URL")
	if dburl == "" {
		log.Fatal("Db Url not found in the environment")
	}
	if portstring == "" {
		log.Fatal("port not found in the environment")
	}
	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("canot connect to database", err)
	}
	queries := db.New(conn)

	apicfg := apiConfig{
		DB: queries,
	}
	go startScraping(queries, 10, time.Minute)
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
	v1router.Post("/create", apicfg.handlerCreateUser)
	v1router.Get("/users", apicfg.Middleware(apicfg.handleGetUserByApi))

	v1router.Post("/feeds", apicfg.Middleware(apicfg.handlerCreateFeed))
	v1router.Get("/feeds", apicfg.HandleGetFeed)

	v1router.Post("/feed_follows", apicfg.Middleware(apicfg.handlerCreateFeedFollow))
	v1router.Get("/feed_follows", apicfg.Middleware(apicfg.handlerGetFeedFollow))
	v1router.Delete("/feed_follows/{feedid}", apicfg.Middleware(apicfg.HandleDeleteFeedFollow))

	v1router.Get("/getposts", apicfg.Middleware(apicfg.handleGetPostUser))
	Router.Mount("/v1", v1router)
	srv := &http.Server{
		Handler: Router,
		Addr:    ":" + portstring,
	}
	log.Printf("server starting on port :%v", portstring)
	fmt.Println("port string:", portstring)

	srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
