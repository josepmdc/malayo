package services

import (
	"malayo/conf"
	"malayo/json"
	"malayo/repos"
)

type MediaService interface {
	GetMovie(ID string) (*repos.Media, error)
	GetTvShow(ID string) (*repos.Media, error)
}

type mediaService struct {
	Movies repos.MediaRepository
	Tv     repos.MediaRepository
}

func NewMediaService(c *conf.Config) *mediaService {
	switch c.Storage {
	case "json":
		return &mediaService{
			Movies: json.NewMediaRepo(c.Media.Movies.JSON, "Movie"),
			Tv:     json.NewMediaRepo(c.Media.Tv.JSON, "TV Show"),
		}
	}
	return nil
}

func (m *mediaService) GetMovie(ID string) (*repos.Media, error) {
	media, err := m.Movies.Get(ID)
	if err != nil {
		return media, err
	}
	return media, nil
}

func (m *mediaService) GetTvShow(ID string) (*repos.Media, error) {
	media, err := m.Tv.Get(ID)
	if err != nil {
		return media, err
	}
	return media, nil
}
