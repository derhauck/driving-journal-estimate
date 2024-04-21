package config

import (
	"derhauck/driving-journal-estimate/public/day"
	"math/rand"
)

type FileConfigStruct struct {
	Date  string `yaml:"date" form:"date" json:"date"`
	Count int    `yaml:"count" form:"count" json:"count"`
}

func (f *FileConfigStruct) ParseToDayRandom() *day.Config {
	return f.ParseToDay(rand.Float64())
}
func (f *FileConfigStruct) ParseToDay(multiplier float64) *day.Config {
	return &day.Config{
		Date: f.Date,
		Lesson: &day.LessonType{
			Multiplier: multiplier,
			Count:      f.Count,
		},
	}
}

type File struct {
	Days     []*FileConfigStruct `yaml:"days" form:"days" json:"days,omitempty"`
	Baseline float64             `yaml:"baseline" form:"baseline" json:"baseline"`
	Total    float64             `yaml:"total" form:"total" json:"total"`
}

func (f *File) DayConfig() []*day.Config {
	var result = make([]*day.Config, 0)
	for _, config := range f.Days {
		result = append(result, config.ParseToDay(f.Baseline+rand.Float64()))
	}
	return result
}
