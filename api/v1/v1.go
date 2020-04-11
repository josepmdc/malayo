package v1

import (
	"malayo/conf"
	"malayo/util"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

var validBearer string

type userResponse struct {
	Name string
	Age  int
}

// NewRouter creates an http router for the API
func NewRouter(config *conf.Config) http.Handler {
	r := chi.NewRouter()

	validBearer = config.Token

	r.With(requireAuthentication).Post("/user", getUser)
	r.Get("/user", getUser)

	return r
}

func getUser(w http.ResponseWriter, r *http.Request) {

	response := userResponse{
		Name: "John Doe",
		Age:  56,
	}

	util.ResponseJSON(w, response, http.StatusOK)
}

// TODO This method is made for testing purpouses and must be removed once the project is set up
func requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		if token != validBearer {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
