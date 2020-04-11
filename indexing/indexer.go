package indexing

import (
	"encoding/json"
	"io/ioutil"
	"malayo/util"
	"os"
	"path/filepath"
)

// File specifies the metadata of a specific file
type File struct {
	Name string
	Ext  string
}

// IndexMediaLibrary takes all of the media in the directory specified in the configuration file
// or on the command flag, and indexes it on a JSON file.
func IndexMediaLibrary(root string) {
	var files []File

	root = util.GetAbsPath(root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		file := File{
			Name: path,
			Ext:  filepath.Ext(path),
		}
		files = append(files, file)
		return nil
	})

	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(files, "", "")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("files.json", json, 0644)
	if err != nil {
		panic(err)
	}

}
