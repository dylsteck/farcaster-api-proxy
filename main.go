package main

import (
	"log"
	"net/http"
	"os"

	httpRouter "farcaster-api-proxy-go/internal/http"
)

func main() {
	r := httpRouter.NewRouter()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
