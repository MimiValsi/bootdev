package main

import (
	"encoding/json"
	"errors"
	"mimivalsi/chirpy/internal/database"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	chirpParams := database.CreateChirpParams{
		Body:   cleaned,
		UserID: params.UserID,
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	},
	)
}

func validateChirp(body string) (string, error) {
	const max_len = 140
	if len(body) > max_len {
		return "", errors.New("Chirp is too long")
	}

	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	cleaned := getCleanedBody(body, badWords)

	return cleaned, nil
}

func getCleanedBody(body string, badWords []string) string {
	body_splited := strings.Split(body, " ")
	for i := range body_splited {
		for _, bw := range badWords {
			if strings.ToLower(body_splited[i]) == bw {
				body = strings.ReplaceAll(body, body_splited[i], "****")
			}
		}
	}

	return body
}

func (cfg *apiConfig) handlerFetchAllChirps(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.FetchAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)

}

func (cfg *apiConfig) handlerFetchChirp(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("ChirpID")
	if id == "" {
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't parse UUID", err)
		return
	}

	dbChirp, err := cfg.db.FetchSingleChirp(r.Context(), uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't fetch uuid", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}
