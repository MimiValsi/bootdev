package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)

	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Must be dev to drop from users", nil)
		return
	}

	cfg.db.DeleteUsers(r.Context())
	w.WriteHeader(http.StatusOK)
}
