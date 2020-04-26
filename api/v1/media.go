package v1

import (
	"malayo/repos"
	"malayo/services"
	"malayo/util"
	"net/http"

	"github.com/go-chi/chi"
)

type mediaHandler struct {
	s services.MediaService
}

func (h *mediaHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/movie", func(r chi.Router) {
		r.Route("/{movieID}", func(r chi.Router) {
			r.Get("/", h.getMovie)
		})
	})

	r.Route("/tv", func(r chi.Router) {
		r.Route("/{tvShowID}", func(r chi.Router) {
			r.Get("/", h.getTvShow)
		})
	})

	return r
}

func (h *mediaHandler) getMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	media, err := h.s.GetMovie(movieID)
	response(media, err, w)
}

func (h *mediaHandler) getTvShow(w http.ResponseWriter, r *http.Request) {
	showID := chi.URLParam(r, "tvShowID")
	media, err := h.s.GetTvShow(showID)
	response(media, err, w)
}

func response(media *repos.Media, err error, w http.ResponseWriter) {
	if err != nil {
		util.ResponseJSON(w, nil, http.StatusNotFound)
		return
	}
	util.ResponseJSON(w, media, http.StatusOK)
}
