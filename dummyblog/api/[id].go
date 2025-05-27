package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Blog struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	PublishedAt string `json:"publishedAt"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Extract the blog id from the path
	id := strings.TrimPrefix(r.URL.Path, "/api/blog/")

	// Build path to JSON file
	path := filepath.Join("data", "blog.json")

	// Read JSON file
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

	// Search for blog by id
	for _, blog := range blogs {
		if blog.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blog)
			return
		}
	}

	// If not found, respond with 404
	http.Error(w, "Blog not found", http.StatusNotFound)
}
