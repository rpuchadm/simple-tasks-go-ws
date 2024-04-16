// utils.go

package main

import (
	"encoding/json"
	"net/http"
)

// respondWithError envía una respuesta de error al cliente con el código de estado y el mensaje proporcionados.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON envía una respuesta JSON al cliente con el código de estado y los datos proporcionados.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		// Si hay un error al serializar los datos, enviamos un error interno del servidor.
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
