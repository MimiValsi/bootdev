package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type chirp_params struct {
	Body string `json:"body"`
}

type error_chirp struct {
	Error string `json:"error"`
	Valid bool   `json:"valid"`
}

func handlerDecodeChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	chirp := chirp_params{
		Body: "",
	}
	respError := error_chirp{
		Error: "",
		Valid: true,
	}

	err := decoder.Decode(&chirp)
	if err != nil {
		log.Printf("Error decoding parameters: %s\n", err)
		respError.Error = "Something went wrong"
		respError.Valid = false
		resp, err := json.Marshal(respError)
		if err != nil {
			log.Printf("Error encoding: %s\n", err)
		}
		w.Write(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	max_len := 140
	if len(chirp.Body) > max_len {
		respError.Error = "Chirpy is too long"
		respError.Valid = false
		resp, err := json.Marshal(respError)
		if err != nil {
			log.Printf("Error encoding: %s\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	respError.Valid = true
	resp, err := json.Marshal(respError)
	if err != nil {
		log.Printf("Error encoding: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
