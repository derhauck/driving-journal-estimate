package day

import (
	"errors"
	"fmt"
	"math/rand"
)

type LessonType struct {
	Multiplier float32
	Count      int
}

func (l *LessonType) GetMultiplier() float32 { return l.Multiplier }
func (l *LessonType) combined() float32 {
	return l.Multiplier * float32(l.Count)
}

func (l *LessonType) GetTotal() float32 {
	return l.Multiplier * float32(l.Count)
}

func (l *LessonType) sanityCheck() bool {
	if l.combined() > 100 || l.combined() < 0 {
		return false
	}
	return true
}

func (l *LessonType) ModMultiplier(percent float32) error {
	old := l.Multiplier
	if l.sanityCheck() == false {
		l.Multiplier = old
		return errors.New(fmt.Sprintf("sanity check went wrong, new value '%f' not valid", percent))
	}
	l.Multiplier = percent
	return nil
}

func NewLessonType() *LessonType {
	return &LessonType{
		Multiplier: 0.3,
		Count:      1,
	}
}

type LessonTypeConfigurationParameter func(l *LessonType)

type Config struct {
	Date   string
	Lesson *LessonType
	Total  float32
}

func NewRandomConfig(date string) *Config {
	config := &Config{
		Date:   date,
		Lesson: NewLessonType(),
		Total:  0,
	}
	err := config.Lesson.ModMultiplier(rand.Float32())
	config.Lesson.Count = rand.Intn(5) + 1
	if err != nil {
		return nil
	}
	return config
}

func NewRandomDays(count int) []*Config {
	days := make([]*Config, count)
	for i := 0; i < count; i++ {
		days[i] = NewRandomConfig(fmt.Sprintf("Day-%d", i+1))
	}
	return days
}
func (c *Config) GetTotal() float32 {
	return c.Total
}

func (c *Config) SetTotal(total float32) {
	c.Total = total
}

func (c *Config) GetDate() string {
	return c.Date
}

func (c *Config) SetDate(date string) {
	c.Date = date
}

func (c *Config) GetLesson() *LessonType {
	return c.Lesson
}

func (c *Config) Validate() error {
	return nil
}

func (c *Config) Calculate(total float32) error {
	old := c.GetTotal()
	c.SetTotal(total * (c.GetLesson().GetMultiplier()))
	if err := c.Validate(); err != nil {
		c.SetTotal(old)
		return err
	}
	return nil
}

func (c *Config) String() string {
	output := fmt.Sprintf("Date: %s\t KM: %.2f", c.GetDate(), c.GetTotal())
	output += fmt.Sprintf("\tMultiplier: %.0f%%", c.GetLesson().GetMultiplier()*100)
	output += fmt.Sprintf("\tTotal: %.0f%%", c.GetLesson().GetTotal()*100)
	return output
}

func (c *Config) Print() {
	fmt.Print(c)
	fmt.Println()
}
