package api

import (
	"fmt"
	v1 "malayo/api/v1"
	"malayo/conf"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter creates an http router fot the server
func NewRouter(config *conf.Config) http.Handler {
	router := chi.NewRouter()

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up root handlers
	router.Get("/", helloWorld)

	// Set up API
	router.Mount("/api/v1/", v1.NewRouter(config))

	return router
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello!")

	if err != nil {
		panic(err)
	}
}
