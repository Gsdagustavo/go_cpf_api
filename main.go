package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
	mux.HandleFunc("/docs", helpRequest)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
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
