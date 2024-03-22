package factory

import (
	"derhauck/driving-journal-estimate/public/config"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromFile(path string) *config.File {
	var file config.File

	reader, err := os.ReadFile(path)
	if err != nil {
		GetLogger().Error(err)
		return nil
	}
	err = yaml.Unmarshal(reader, &file)
	if err != nil {
		GetLogger().Error(err)
		return nil
	}

	if file.Baseline == 0.0 {
		file.Baseline = .5
	}

	if file.Total == 0.0 {
		file.Total = 10000
	}

	GetLogger().Logf("Baseline: %f", file.Baseline)
	return &file

}
