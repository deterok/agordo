package loaders

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

type FileLoader struct {
	paths      []string
	loadedPath string
	data       []byte
}

func NewFileLoader(paths ...string) *FileLoader {
	return &FileLoader{paths: paths}
}

func (f *FileLoader) Load() error {
	for _, path := range f.paths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		f.loadedPath = path
		f.data = data
		return nil
	}

	// FIXME: need error
	return nil
}

func (f *FileLoader) GetMap() (map[string]interface{}, error) {
	// var result map[string]interface{}

	ext := filepath.Ext(f.loadedPath)

	switch ext {
	case ".yaml":
		return f.parseYAML()

	case ".json":
		return f.parseJSON()

	case ".ini":
		return f.parseINI()

	default:
		// TODO: add error
		return map[string]interface{}{}, nil
	}
}

func (f *FileLoader) parseYAML() (map[string]interface{}, error) {
	var parsedData map[interface{}]interface{}
	err := yaml.Unmarshal(f.data, &parsedData)

	if err != nil {
		return nil, err
	}

	convertedData, err := convertMap(parsedData)

	if err != nil {
		return nil, err
	}

	return convertedData, err
}

func (f *FileLoader) parseJSON() (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(f.data, &result)
	return result, err
}

func (f *FileLoader) parseINI() (map[string]interface{}, error) {
	var result map[string]interface{}

	ini, err := ini.Load(f.data)
	if err != nil {
		return result, err
	}

	err = ini.MapTo(result)
	return result, err
}
