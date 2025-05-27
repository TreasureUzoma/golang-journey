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
	id := strings.TrimPrefix(r.URL.Path, "/api/blog/")

	path := filepath.Join("data", "blog.json")
	file, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Error loading blog data", http.StatusInternalServerError)
		return
	}

	var blogs []Blog
	if err := json.Unmarshal(file, &blogs); err != nil {
		http.Error(w, "Error parsing blog data", http.StatusInternalServerError)
		return
	}

	for _, blog := range blogs {
		if blog.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blog)
			return
		}
	}

	http.Error(w, "Blog not found", http.StatusNotFound)
}
