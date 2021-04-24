package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	Erro string `json:"erro"`
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

// Handle error that have status code more than 400
func HandleStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var errorAPI ErrorAPI

	json.NewDecoder(r.Body).Decode(&errorAPI)
	ResponseJSON(w, r.StatusCode, errorAPI)
}
