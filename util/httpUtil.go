package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ResponseJSON returns an HTTP response in JSON format
func ResponseJSON(w http.ResponseWriter, data interface{}, c int) {
	dj, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	_, err = fmt.Fprintf(w, "%s", dj)

	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
