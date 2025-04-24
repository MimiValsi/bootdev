package main

import (
	"database/sql"
	"mimivalsi/chirpy/internal/auth"
	"mimivalsi/chirpy/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get bearer token", err)
		return
	}

	tokenString := uuid.MustParse(bearerToken)

	refreshToken, err := cfg.db.CheckToken(r.Context(), tokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't check for refresh token", err)
		return
	}

	// refreshToken.ExpiresAt
	if time.Now().After(refreshToken.ExpiresAt) {
		respondWithError(w, http.StatusUnauthorized, "Refresh token expired", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: bearerToken,
	})
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get bearer token", err)
		return
	}

	var tn sql.NullTime
	if tn.Valid {
		tn.Time.UTC()
	}
	err = cfg.db.RevokeToken(r.Context(), database.RevokeTokenParams{
		UserID:    uuid.MustParse(bearerToken),
		RevokedAt: tn,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke token", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
