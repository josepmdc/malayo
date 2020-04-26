package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"malayo/repos"
)

type mediaRepo struct {
	IndexFile string
	MediaType string
}

func NewMediaRepo(indexFile string, mediaType string) repos.MediaRepository {
	return &mediaRepo{
		IndexFile: indexFile,
		MediaType: mediaType,
	}
}

func (m *mediaRepo) FindAll() (*[]repos.Media, error) {
	mediaeJSON, err := ioutil.ReadFile(m.IndexFile)
	if err != nil {
		return nil, err
	}
	mediae := &[]repos.Media{}
	err = json.Unmarshal(mediaeJSON, mediae)
	if err != nil {
		return nil, err
	}
	return mediae, nil
}

func (m *mediaRepo) Get(ID string) (*repos.Media, error) {
	mediaeJSON, err := ioutil.ReadFile(m.IndexFile)
	if err != nil {
		return nil, err
	}
	mediae := []repos.Media{}
	err = json.Unmarshal(mediaeJSON, &mediae)
	if err != nil {
		return nil, err
	}
	for _, media := range mediae {
		if media.ID == ID {
			return &media, nil
		}
	}
	return nil, fmt.Errorf("%s with ID %s not found", m.MediaType, ID)
}

func (m *mediaRepo) Create(_ repos.Media) *repos.Media {
	return &repos.Media{} // TODO Implement function
}

func (m *mediaRepo) Update(_ string) (*repos.Media, error) {
	return &repos.Media{}, nil // TODO Implement function
}

func (m *mediaRepo) Delete(_ string) error {
	return nil // TODO Implement function
}
