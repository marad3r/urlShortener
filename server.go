package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"net/http"
	"time"
)

const (
	htmlIndex = `<html><body>Welcome!</body></html>`
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, htmlIndex)
}

// NewRouter returns a new HTTP handler that implements the main server routes
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up our root handlers
	router.Get("/", handleIndex)
	router.Get("/api/create", createHandler)

	return router
}

func NewServer(c Config) *http.Server {
	handler := NewRouter()
	return &http.Server{
		Addr:         c.Port,
		Handler:      handler,
		ReadTimeout:  time.Duration(c.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(c.Server.IdleTimeout) * time.Second,
	}
}
