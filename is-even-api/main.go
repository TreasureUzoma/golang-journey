package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	IsEven bool `json:"isEven"`
}

func isEvenHandler(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("number")
	if numStr == "" {
		http.Error(w, "number parameter is required", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "number parameter is required", http.StatusBadRequest)
		return
	}

	isEven := number % 2 == 0
	response := Response{IsEven: isEven}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/is-even", isEvenHandler)
	fmt.Println("Server is listening on port 8080");
	log.Fatal(http.ListenAndServe(":8080", nil))
}