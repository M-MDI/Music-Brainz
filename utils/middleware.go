package Groupietracker

import (
    "fmt"
    "net/http"
    "strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := getArtistIDFromPath(r.URL.Path)
        fmt.Println("Artist ID:", id)
        next.ServeHTTP(w, r)
    }
}

func getArtistIDFromPath(path string) string {
    parts := strings.Split(path, "/")
    if len(parts) > 2 {
        return parts[2]
    }
    return ""
}