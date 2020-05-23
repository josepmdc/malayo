package v1

import (
	"encoding/json"
	"fmt"
	"malayo/domain"
	"malayo/services"
	"malayo/util"
	"net/http"

	"github.com/go-chi/chi"
)

type mediaHandler struct {
	service services.MediaService
}

func (handler *mediaHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/movie", func(r chi.Router) {
		r.Route("/{movieID}", func(r chi.Router) {
			r.Get("/", handler.getMovie("movieID", handler.service))
		})
		r.Route("/new", func(r chi.Router) {
			r.Post("/", handler.createMovie(handler.service))
		})
	})

	r.Route("/tv", func(r chi.Router) {
		r.Route("/{tvShowID}", func(r chi.Router) {
			r.Get("/", handler.getTvShow("tvShowID", handler.service))
		})
		r.Route("/new", func(r chi.Router) {
			r.Post("/", handler.createTvShow(handler.service))
		})
	})

	return r
}

func (handler *mediaHandler) getMovie(param string, service services.MediaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieID := chi.URLParam(r, param)
		movie, err := service.GetMovie(movieID)
		response(movie, err, w)
	}
}

func (handler *mediaHandler) createMovie(service services.MediaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var movie domain.Movie
		err := decoder.Decode(&movie)
		if err != nil {
			fmt.Println(err.Error())
			util.ResponseJSON(w, nil, http.StatusBadRequest)
			return
		}
		_, err = service.CreateMovie(&movie)
		if err != nil {
			fmt.Println(err.Error())
			util.ResponseJSON(w, nil, http.StatusBadRequest)
			return
		}
		util.ResponseJSON(w, nil, http.StatusOK)
	}
}

func (handler *mediaHandler) getTvShow(param string, service services.MediaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tvShowID := chi.URLParam(r, param)
		tvShow, err := service.GetTvShow(tvShowID)
		response(tvShow, err, w)
	}
}

func (handler *mediaHandler) createTvShow(service services.MediaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var tvShow domain.TvShow

		if err := decoder.Decode(&tvShow); err != nil {
			fmt.Println(err.Error())
			util.ResponseJSON(w, nil, http.StatusBadRequest)
			return
		}

		if _, err := service.CreateTvShow(&tvShow); err != nil {
			fmt.Println(err.Error())
			util.ResponseJSON(w, nil, http.StatusBadRequest)
			return
		}
		util.ResponseJSON(w, nil, http.StatusOK)
	}
}

func response(media interface{}, err error, w http.ResponseWriter) {
	if err != nil {
		util.ResponseJSON(w, "File not found", http.StatusNotFound)
		return
	}
	util.ResponseJSON(w, media, http.StatusOK)
}
