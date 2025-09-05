package http

import (
	"farcaster-api-proxy-go/internal/handlers"
	"farcaster-api-proxy-go/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	
	r.Use(middleware.CORS)
	
	r.HandleFunc("/health", handlers.Health).Methods("GET")
	r.HandleFunc("/{username}/{hash}", handlers.UserThreadCasts).Methods("GET")
	r.HandleFunc("/{username}", handlers.UserByUsername).Methods("GET")
	r.HandleFunc("/fids/{fid}", handlers.UserByFid).Methods("GET")
	r.HandleFunc("/search/casts", handlers.SearchCasts).Methods("GET")
	r.HandleFunc("/search/users", handlers.SearchUsers).Methods("GET")
	r.HandleFunc("/search/summary", handlers.SearchSummary).Methods("GET")
	
	return r
}
