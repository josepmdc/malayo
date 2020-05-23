package domain

type Movie struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Path string `json:"path"`
	Ext  string `json:"ext"`
}

type TvShow struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Path string `json:"path"`
	Ext  string `json:"ext"`
}

// MovieRepository defines an interface for accessing the media stored
type MovieRepository interface {
	FindAll() (*[]Movie, error)
	Get(ID string) (*Movie, error)
	Create(media *Movie) (*Movie, error)
	Update(ID string) (*Movie, error)
	Delete(ID string) error
}

// TvRepository defines an interface for accessing the media stored
type TvRepository interface {
	FindAll() (*[]TvShow, error)
	Get(ID string) (*TvShow, error)
	Create(tvShow *TvShow) (*TvShow, error)
	Update(ID string) (*TvShow, error)
	Delete(ID string) error
}
