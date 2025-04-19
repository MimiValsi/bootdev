package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidation(w http.ResponseWriter, r *http.Request) {
	bad_words := []string{"kerfuffle", "sharbert", "fornax"}

	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
		Valid       bool   `json:"valid"`
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
				// body_splited[i] = "****"
				params.Body = strings.ReplaceAll(params.Body, body_splited[i], "****")
			}
		}
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: params.Body,
	})
}
