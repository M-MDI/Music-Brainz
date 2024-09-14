package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	if strings.Contains(r.URL.Path, "/templates/") {
		ErrorHandler(w, "This Page Not Found !!", http.StatusNotFound)
		return
	}
	path := r.URL.Path[len("/src"):]
	fullPath := filepath.Join("./static", path)
	filepath, err := os.Stat(fullPath)
	if err != nil {
		ErrorHandler(w, "This Page Not Found !!", http.StatusNotFound)
		return
	}
	if !filepath.IsDir() {
		http.StripPrefix("/src", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	} else {
		ErrorHandler(w, "This Page Not Found !!", http.StatusNotFound)
		return
	}
}
