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

type tvRepoJSON struct {
	IndexFile string
}

func NewTvRepo(c conf.MediaInfo) *tvRepoJSON {
	return &tvRepoJSON{
		IndexFile: c.JSON,
	}
}

func (mr *tvRepoJSON) FindAll() (*[]domain.TvShow, error) {
	tvJSON, err := ioutil.ReadFile(mr.IndexFile)
	if err != nil {
		return nil, err
	}
	var tvShow []domain.TvShow
	err = json.Unmarshal(tvJSON, &tvShow)
	if err != nil {
		return nil, err
	}
	return &tvShow, nil
}

func (mr *tvRepoJSON) Get(ID string) (*domain.TvShow, error) {
	tvsJSON, err := ioutil.ReadFile(mr.IndexFile)
	if err != nil {
		return nil, err
	}
	var tvShows []domain.TvShow
	err = json.Unmarshal(tvsJSON, &tvShows)
	if err != nil {
		return nil, err
	}
	for _, tvShow := range tvShows {
		if tvShow.ID == ID {
			return &tvShow, nil
		}
	}
	return nil, fmt.Errorf("tv with ID %s not found", ID)
}

func (mr *tvRepoJSON) Create(m *domain.TvShow) (*domain.TvShow, error) {
	tv, err := unmarshalTvJSON(mr.IndexFile)
	if err != nil {
		return nil, err
	}
	tv = append(tv, *m)

	err = util.WriteToJSON(tv, mr.IndexFile, 0644)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (mr *tvRepoJSON) Update(_ string) (*domain.TvShow, error) {
	panic("implement me") // TODO Implement function
}

func (mr *tvRepoJSON) Delete(_ string) error {
	panic("implement me") // TODO Implement function
}

func unmarshalTvJSON(file string) ([]domain.TvShow, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var data []domain.TvShow

	if err := json.Unmarshal(f, &data); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return data, nil
}
