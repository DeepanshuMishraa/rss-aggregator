package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	fmt.Println("Port: ", port)

	dbString := os.Getenv("DATABASE_URL")

	if dbString == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	fmt.Println("Database URL: ", dbString)
	conn, err := sql.Open("postgres", dbString)

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	apiConfig := apiConfig{
		DB:database.New
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()

	v1Router.Get("/health", HandlerReadiness)
	v1Router.Get("/error", handleErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %s\n", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	fmt.Println("Hello, World!")
}
