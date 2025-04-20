package main

import (
	"database/sql"
	"fmt"
	"log"
	"mimivalsi/chirpy/internal/database"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	fpRoot = "."
	port   = "8080"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func main() {

	// little package that reads from .env file inside project root folder
	// and sets it to ENV
	godotenv.Load()

	// with .env file set to ENV, it is "safe" to read it with os.Getenv
	// function.
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
	}

	mux := http.NewServeMux()

	handler := http.StripPrefix("/app", http.FileServer(http.Dir(fpRoot)))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidation)
	mux.Handle("POST /api/users", middlewareDB(&apiCfg))

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(s.ListenAndServe())

}
