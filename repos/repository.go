package repos

// MediaRepository defines an interface for accesing the media stored
type MediaRepository interface {
	FindAll() (*[]Media, error)
	Get(ID string) (*Media, error)
	Create(media Media) *Media
	Update(ID string) (*Media, error)
	Delete(ID string) error
}

// Media is the central class in the domain model
type Media struct {
	ID   string
	Type string
	Path string
	Ext  string
}
