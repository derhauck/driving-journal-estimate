package factory

import (
	"driving-journal-estimate/public/config"
	"gopkg.in/yaml.v3"
)

func LoadConfigFromFile(path string) *config.File {
	var file config.File
	err := yaml.Unmarshal([]byte(path), &file)
	if err != nil {
		return nil
	}

	return &file

}
