package indexing

import (
	"encoding/json"
	"io/ioutil"
	"malayo/conf"
	"malayo/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type File struct {
	ID   string
	Name string
	Ext  string
	Path string
}

// IndexMediaLibrary takes all of the media in the directory specified in the configuration file
// or on the command flag, and indexes it on a JSON file.
func IndexMediaLibrary(config *conf.Media) error {
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
	_, err = indexDirectory(config.Books.Path, config.Books.JSON)
	if err != nil {
		return err
	}
	return nil
}

func indexDirectory(root string, JSON string) ([]File, error) {
	root = util.GetAbsPath(root)
	var files []File
	id := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Ignore if it's not a file
		if info.IsDir() {
			return nil
		}

		_, name := filepath.Split(path)
		ext := strings.ToLower(filepath.Ext(name))

		files = append(files, File{
			ID:   strconv.Itoa(id),
			Name: name,
			Ext:  ext,
			Path: path,
		})
		id++
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = writeFilesJSON(JSON, files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func writeFilesJSON(path string, files []File) error {
	j, err := json.MarshalIndent(files, "", "")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, j, 0644)
	if err != nil {
		return err
	}
	return nil
}
