package main

import (
	"encoding/json"
	"mimivalsi/chirpy/internal/database"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	bad_words := []string{"kerfuffle", "sharbert", "fornax"}

	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	type returnVals struct {
		Chirp
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const max_len = 140
	if len(params.Body) > max_len {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	// var cleaned_body string
	body_splited := strings.Split(params.Body, " ")
	for i := range body_splited {
		for _, bw := range bad_words {
			if strings.ToLower(body_splited[i]) == bw {
				params.Body = strings.ReplaceAll(params.Body, body_splited[i], "****")
			}
		}
	}

	chirpParams := database.CreateChirpParams{
		Body:   params.Body,
		UserID: params.UserID,
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, returnVals{
		Chirp: Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		},
	})
}
