package upstream

import (
	"io"
	"net/http"
)

func ProxyRequest(w http.ResponseWriter, url string) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch from Farcaster API"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, `{"error": "Failed to read response"}`, http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
