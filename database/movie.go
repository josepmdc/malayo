package postgres

import (
	"database/sql"
	"malayo/domain"
)

type movieRepoDB struct {
	DB *sql.DB
}

func NewMovieRepo(db *sql.DB) *movieRepoDB {
	return &movieRepoDB{
		DB: db,
	}
}

func (m movieRepoDB) FindAll() (*[]domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (m movieRepoDB) Get(ID string) (*domain.Movie, error) {
	var id, t, path, ext string
	err := m.DB.QueryRow("SELECT * FROM MOVIES WHERE id = $1", ID).Scan(&id, &t, &path, &ext)
	if err != nil {
		return nil, err
	}
	movie := domain.Movie{
		ID:   id,
		Type: t,
		Path: path,
		Ext:  ext,
	}
	return &movie, nil
}

func (m movieRepoDB) Create(movie *domain.Movie) (*domain.Movie, error) {
	stmt, err := m.DB.Prepare("INSERT INTO movies VALUES ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(movie.ID, movie.Type, movie.Path, movie.Ext)
	if err != nil {
		return nil, err
	}
	return &domain.Movie{}, nil // TODO Return created element so we can test method
}

func (m movieRepoDB) Update(ID string) (*domain.Movie, error) {
	panic("implement me") // TODO Implement function
}

func (m movieRepoDB) Delete(ID string) error {
	panic("implement me") // TODO Implement function
}
