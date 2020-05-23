package postgres

import (
	"database/sql"
	"malayo/conf"
	"malayo/domain"
)

type movieRepoDB struct {
	DB *sql.DB
}

func NewMovieRepo(c conf.MediaInfo) *movieRepoDB {
	return &movieRepoDB{
		// TODO assign DB
	}
}

func (m movieRepoDB) FindAll() (*[]domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (m movieRepoDB) Get(ID string) (*domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (m movieRepoDB) Create(media *domain.Movie) (*domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (m movieRepoDB) Update(ID string) (*domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (m movieRepoDB) Delete(ID string) error {
	panic("implement me") // TODO Implement function
}
