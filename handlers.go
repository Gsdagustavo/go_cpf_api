package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	IsValid bool `json:"is_valid"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only get allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	values := r.URL.Query()

	if !values.Has("cpf") {
		http.Error(w, "CPF is required", http.StatusBadRequest)
		return
	}

	cpf := values.Get("cpf")

	if cpf == "" || len(cpf) != 11 {
		http.Error(w, "Invalid CPF", http.StatusBadRequest)
		return
	}

	isValid := validateCPF(cpf)
	resp := Response{IsValid: isValid}

	_ = json.NewEncoder(w).Encode(resp)
}

func helpRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only get allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	http.ServeFile(w, r, "help.html")
}
