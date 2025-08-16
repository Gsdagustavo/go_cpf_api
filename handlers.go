package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only post allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "An error occurred while trying to read the r body", http.StatusInternalServerError)
		return
	}

	var body Request

	err = json.Unmarshal(bytes, &body)

	if err != nil {
		http.Error(w, "An error occurred while trying to unmarshal the r body", http.StatusInternalServerError)
		return
	}

	isValid := validateCPF(body.Cpf)

	if isValid {
		response := Response{IsValid: isValid}

		bytes, err = json.Marshal(response)

		w.Write(bytes)
	}
}
