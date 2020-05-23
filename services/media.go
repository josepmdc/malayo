package services

import (
	"fmt"
	"malayo/conf"
	"malayo/domain"
	"malayo/json"
)

type MediaService interface {
	GetMovie(ID string) (*domain.Movie, error)
	CreateMovie(movie *domain.Movie) (*domain.Movie, error)
	GetTvShow(ID string) (*domain.TvShow, error)
	CreateTvShow(tv *domain.TvShow) (*domain.TvShow, error)
}

type mediaService struct {
	Config          *conf.Config
	MovieRepository domain.MovieRepository
	TvRepository    domain.TvRepository
}

func NewMediaService(c *conf.Config) *mediaService {
	switch c.Storage {
	case "json":
		return &mediaService{
			Config:          c,
			MovieRepository: json.NewMovieRepo(c.Media.Movies),
			TvRepository:    json.NewTvRepo(c.Media.Tv),
		}
	}
	return nil
}

func (ms *mediaService) GetMovie(ID string) (*domain.Movie, error) {
	movie, err := ms.MovieRepository.Get(ID)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (ms *mediaService) CreateMovie(movie *domain.Movie) (*domain.Movie, error) {
	movie.Path = fmt.Sprintf("%s/%s", ms.Config.Media.Movies.Path, movie.Path)
	m, err := ms.MovieRepository.Create(movie)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (ms *mediaService) GetTvShow(ID string) (*domain.TvShow, error) {
	tvShow, err := ms.TvRepository.Get(ID)
	if err != nil {
		return tvShow, err
	}
	return tvShow, nil
}

func (ms *mediaService) CreateTvShow(tv *domain.TvShow) (*domain.TvShow, error) {
	tv.Path = fmt.Sprintf("%s/%s", ms.Config.Media.Tv.Path, tv.Path)
	m, err := ms.TvRepository.Create(tv)
	if err != nil {
		return nil, err
	}
	return m, nil
}
