package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Cpf string `json:"cpf"`
}

type Response struct {
	IsValid bool `json:"is_valid"`
}

func main() {

	http.HandleFunc("/cpf", func(w http.ResponseWriter, r *http.Request) {

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
	})

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func validateCPF(cpf string) (isValid bool) {
	if len(cpf) != 11 {
		return
	}

	allSame := true

	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}

	if allSame {
		return
	}

	var numbers [11]int

	for i, ch := range cpf {
		numbers[i] = int(ch - '0')
	}

	// first check
	sum1 := 0
	for i := 0; i < 9; i++ {
		sum1 += numbers[i] * (10 - i)
	}

	rem1 := sum1 % 11
	check1 := 0
	if rem1 >= 2 {
		check1 = 11 - rem1
	}

	// second check
	sum2 := 0
	for i := 0; i < 10; i++ {
		sum2 += numbers[i] * (11 - i)
	}

	rem2 := sum2 % 11
	check2 := 0
	if rem2 >= 2 {
		check2 = 11 - rem2
	}

	return numbers[9] == check1 && numbers[10] == check2
}
