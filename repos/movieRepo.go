package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"malayo/models"
)

// GetMovieByID returns the path of the movie that matches the specified ID
func GetMovieByID(id int) (string, error) {
	movies, err := ioutil.ReadFile("movies.json") // TODO: Read file path from config
	if err != nil {
		return "", err
	}
	files := []models.File{}
	err = json.Unmarshal([]byte(movies), &files)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.ID == id {
			return file.Path, nil
		}
	}
	return "", fmt.Errorf("Movie with ID %d not found", id)
}
