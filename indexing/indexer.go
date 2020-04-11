package indexing

import (
	"encoding/json"
	"io/ioutil"
	"malayo/models"
	"malayo/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// IndexMediaLibrary takes all of the media in the directory specified in the configuration file
// or on the command flag, and indexes it on a JSON file.
func IndexMediaLibrary(root string) {
	root = util.GetAbsPath(root)
	indexMovies(root)
	indexMusic(root)
}

func indexMovies(root string) {
	files := indexDirectory(root, viper.GetStringSlice("videoFormats"), viper.GetStringSlice("videoFolders"))
	writeJSON("videos.json", files) // TODO Specify in config the output file
}

func indexMusic(root string) {
	files := indexDirectory(root, viper.GetStringSlice("audioFormats"), viper.GetStringSlice("musicFolders"))
	writeJSON("music.json", files) // TODO Specify in config the output file
}

func indexDirectory(root string, extensions []string, dirs []string) []models.File {
	var files []models.File
	id := 0
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		_, name := filepath.Split(path)
		ext := strings.ToLower(filepath.Ext(name))
		_, foundExt := util.Find(extensions, ext)

		subDirs := strings.Split(path, "/")
		foundDir := false
		for _, subDir := range subDirs {
			_, foundDir = util.Find(dirs, subDir)
			if foundDir {
				break
			}
		}

		if !foundExt || !foundDir {
			return nil
		}

		files = append(files, models.File{
			ID:   id,
			Name: name,
			Ext:  ext,
			Path: path,
		})
		id++
		return nil
	})
	return files
}

func writeJSON(path string, files []models.File) {
	json, err := json.MarshalIndent(files, "", "")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, json, 0644)
	if err != nil {
		panic(err)
	}
}
