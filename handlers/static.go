package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, "/templates/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	path := r.URL.Path[len("/src"):]

	fullPath := filepath.Join("./static", path)

	filepath, err := os.Stat(fullPath)

	if err != nil {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}

	if !filepath.IsDir() {
		http.StripPrefix("/src", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
