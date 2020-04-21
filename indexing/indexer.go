package indexing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"malayo/conf"
	"malayo/models"
	"malayo/util"
	"os"
	"path/filepath"
	"strings"
)

// IndexMediaLibrary takes all of the media in the directory specified in the configuration file
// or on the command flag, and indexes it on a JSON file.
func IndexMediaLibrary(config *conf.Config) error {
	_, err := indexDirectory(config.Movies.Path, config.Movies.JSON)
	if err != nil {
		return err
	}
	_, err = indexDirectory(config.Tv.Path, config.Tv.JSON)
	if err != nil {
		return err
	}
	_, err = indexDirectory(config.Music.Path, config.Music.JSON)
	if err != nil {
		return err
	}
	return nil
}

func indexDirectory(root string, JSON string) ([]models.File, error) {
	fmt.Println(root)
	root = util.GetAbsPath(root)
	var files []models.File
	id := 0
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		_, name := filepath.Split(path)
		ext := strings.ToLower(filepath.Ext(name))

		files = append(files, models.File{
			ID:   id,
			Name: name,
			Ext:  ext,
			Path: path,
		})
		id++
		return nil
	})
	err := writeFilesJSON(JSON, files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func writeFilesJSON(path string, files []models.File) error {
	json, err := json.MarshalIndent(files, "", "")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, json, 0644)
	if err != nil {
		return err
	}
	return nil
}
