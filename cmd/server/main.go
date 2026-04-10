package main

import (
	"log"
	"net/http"
	"os"

	"greenapitest/internal/api"
	"greenapitest/internal/handlers"
)

func main() {
	client := api.NewClient()
	h := handlers.NewHandler(client)
	mux := http.NewServeMux()

	mux.HandleFunc("/api/getSettings", h.GetSettings)
	mux.HandleFunc("/api/getStateInstance", h.GetStateInstance)
	mux.HandleFunc("/api/sendMessage", h.SendMessage)
	mux.HandleFunc("/api/sendFileByUrl", h.SendFileByURL)

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "web/templates/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
