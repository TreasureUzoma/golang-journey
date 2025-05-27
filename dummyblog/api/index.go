package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type Blog struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	PublishedAt string `json:"publishedAt"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Build path to JSON file
	path := filepath.Join("data", "blog.json")

	// Read the JSON file
	file, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Error loading blog data", http.StatusInternalServerError)
		return
	}

	// Parse JSON into slice of blogs
	var blogs []Blog
	if err := json.Unmarshal(file, &blogs); err != nil {
		http.Error(w, "Error parsing blog data", http.StatusInternalServerError)
		return
	}

	// Set content type and respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}
