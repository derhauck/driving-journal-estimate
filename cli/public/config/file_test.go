package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestFileConfigStruct_NewDay(t *testing.T) {
	t.Run("Can create new day", func(t *testing.T) {
		expectation := FileConfigStruct{
			Date:  "2021-01-01",
			Count: 10,
		}
		result := expectation.ParseToDayRandom()
		if result.Date != expectation.Date {
			t.Errorf("Expected date '%s', got '%s'", expectation.Date, result.Date)
		}

		if result.Lesson.Count != expectation.Count {
			t.Errorf("Expected count '%d', got '%d'", expectation.Count, result.Lesson.Count)
		}
		if result.Lesson.GetMultiplier() > 1.0 || result.Lesson.GetMultiplier() < 0.0 {
			t.Errorf("Expected multiplier to be within range of 0-1, got '%f'", result.Lesson.GetMultiplier())
		}
	})
}

func TestFile_DayConfig(t *testing.T) {
	t.Run("Can create new day", func(t *testing.T) {
		path, err := os.Getwd()
		if err != nil {
			t.Error(err)
		}
		path = path + "/test.config.yaml"
		reader, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
		}
		var file File
		err = yaml.Unmarshal(reader, &file)
		if err != nil {
			t.Error(err)
		}
		result := file.DayConfig()
		if len(result) == 0 {
			t.Error("Result is empty")
		}
		if result[0].Date != "10-02-2023" {
			t.Errorf("Expected date '10-02-2023', got '%s'", result[0].Date)
		}
		if result[0].Lesson.Count != 10 {
			t.Errorf("Expected count '10', got '%d'", result[0].Lesson.Count)
		}

	})
}
