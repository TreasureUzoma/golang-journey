package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	IsEven bool `json:"isEven"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	numStr := r.URL.Query().Get("number")
	if numStr == "" {
		http.Error(w, "number parameter is required", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "number must be an integer", http.StatusBadRequest)
		return
	}

	isEven := number%2 == 0
	response := Response{IsEven: isEven}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Required main function
func main() {
	vercel.Start(context.Background(), Handler)
}
