package json

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"malayo/conf"
	"malayo/domain"
	"malayo/util"
)

type movieRepoJSON struct {
	IndexFile string
}

func NewMovieRepo(c conf.MediaInfo) *movieRepoJSON {
	return &movieRepoJSON{
		IndexFile: c.JSON,
	}
}

func (mr *movieRepoJSON) FindAll() (*[]domain.Movie, error) {
	mediaJSON, err := ioutil.ReadFile(mr.IndexFile)
	if err != nil {
		return nil, err
	}
	var media []domain.Movie
	err = json.Unmarshal(mediaJSON, &media)
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (mr *movieRepoJSON) Get(ID string) (*domain.Movie, error) {
	moviesJSON, err := ioutil.ReadFile(mr.IndexFile)
	if err != nil {
		return nil, err
	}
	var movies []domain.Movie
	err = json.Unmarshal(moviesJSON, &movies)
	if err != nil {
		return nil, err
	}
	for _, m := range movies {
		if m.ID == ID {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("movie with ID %s not found", ID)
}

func (mr *movieRepoJSON) Create(m *domain.Movie) (*domain.Movie, error) {
	movie, err := unmarshalMovieJSON(mr.IndexFile)
	if err != nil {
		return nil, err
	}
	movie = append(movie, *m)

	err = util.WriteToJSON(movie, mr.IndexFile, 0644)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (mr *movieRepoJSON) Update(_ string) (*domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (mr *movieRepoJSON) Delete(_ string) error {
	panic("implement me") // TODO Implement function
}

func unmarshalMovieJSON(file string) ([]domain.Movie, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var data []domain.Movie

	if err := json.Unmarshal(f, &data); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return data, nil
}
