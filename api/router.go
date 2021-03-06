package api

import (
	"fmt"
	v1 "malayo/api/v1"
	"malayo/conf"
	"malayo/indexing"
	"malayo/services"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter creates an http router fot the server
func NewRouter(config *conf.Config, ms services.MediaService) http.Handler {
	router := chi.NewRouter()

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up root handlers
	router.Get("/", helloWorld)
	router.Get("/generate", generateIndex(&config.Media))

	// Set up API
	router.Mount("/api/v1/", v1.NewRouter(config, ms))

	return router
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello!")

	if err != nil {
		panic(err)
	}
}

func generateIndex(config *conf.Media) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := indexing.IndexMediaLibrary(config)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}
