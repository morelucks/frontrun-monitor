package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Server represents the HTTP server
type Server struct {
	config *Config
	mux    *http.ServeMux
}

// NewServer creates a new server instance
func NewServer(cfg *Config) *Server {
	return &Server{
		config: cfg,
		mux:    http.NewServeMux(),
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	// Register routes
	s.registerRoutes()

	addr := ":" + strconv.Itoa(s.config.Port)
	log.Printf("Server starting on %s (environment: %s)", addr, s.config.Env)

	return http.ListenAndServe(addr, s.mux)
}

// registerRoutes registers all API routes
func (s *Server) registerRoutes() {
	// Health check endpoint
	s.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy"}`)
	})

	// Root endpoint
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"Frontrun Monitor API","version":"1.0.0"}`)
	})
}
