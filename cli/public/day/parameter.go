package day

import (
	"driving-journal-estimate/public/misc"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type LessonTypeParameter struct {
	misc.Parameter[float32]
	multiplier float32
}

func (l *LessonTypeParameter) GetTotal() float32 {
	return l.GetValue() * l.GetMultiplier()
}

func (l *LessonTypeParameter) SetMultiplier(multiplier float32) {
	l.multiplier = multiplier
}

func (l *LessonTypeParameter) GetMultiplier() float32 {
	return l.multiplier
}
func NewLessonTypeParameter(key string, value float32, multiplier float32) *LessonTypeParameter {
	tmp := &LessonTypeParameter{}
	tmp.SetKey(key)
	tmp.SetValue(value)
	tmp.SetMultiplier(multiplier)
	return tmp
}
func validLessonParameter(percent float32) bool {
	if percent < 0 || percent > 100 {
		return false
	}
	return true
}

func (l *LessonTypeParameter) modify(percent float32) {
	if validLessonParameter(percent) {
	}
	l.SetValue(percent)
}

func NewLessonTypeCity(percent float32) *LessonTypeParameter {
	if !validLessonParameter(percent) {
		return nil
	}
	return NewLessonTypeParameter("CITY", percent, 0.3)
}

func NewLessonTypeLand(percent float32) *LessonTypeParameter {
	if !validLessonParameter(percent) {
		return nil
	}
	return NewLessonTypeParameter("LAND", percent, 0.6)
}

func NewLessonTypeHighway(percent float32) *LessonTypeParameter {
	if !validLessonParameter(percent) {
		return nil
	}
	return NewLessonTypeParameter("HIGHWAY", percent, 1)
}

type LessonType struct {
	City    *LessonTypeParameter
	Land    *LessonTypeParameter
	Highway *LessonTypeParameter
}

func (l *LessonType) GetCity() *LessonTypeParameter    { return l.City }
func (l *LessonType) GetLand() *LessonTypeParameter    { return l.Land }
func (l *LessonType) GetHighway() *LessonTypeParameter { return l.Highway }
func (l *LessonType) combined() float32 {
	return l.Land.GetValue() + l.Highway.GetValue() + l.City.GetValue()
}

func (l *LessonType) GetTotal() float32 {
	return l.Land.GetTotal() + l.Highway.GetTotal() + l.City.GetTotal()
}

func (l *LessonType) sanityCheck() bool {
	if l.combined() > 100 || l.combined() < 0 {
		return false
	}
	return true
}

func (l *LessonType) ModCity(percent float32) error {
	old := l.City.GetValue()
	l.City.modify(percent)
	if l.sanityCheck() == false {
		l.City.SetValue(old)
		return errors.New(fmt.Sprintf("sanity check went wrong, new value '%f' not valid", percent))
	}
	return nil
}

func (l *LessonType) ModHighway(percent float32) error {
	old := l.Highway.GetValue()
	l.Highway.modify(percent)
	if l.sanityCheck() == false {
		l.Highway.SetValue(old)
		return errors.New(fmt.Sprintf("sanity check went wrong, new value '%f' not valid", percent))
	}
	return nil
}

func (l *LessonType) ModLand(percent float32) error {
	old := l.Land.GetValue()
	l.Land.modify(percent)
	if l.sanityCheck() == false {
		l.Land.SetValue(old)
		return errors.New(fmt.Sprintf("sanity check went wrong, new value '%f' not valid", percent))
	}
	return nil
}

func NewLessonType() *LessonType {
	return &LessonType{
		City:    NewLessonTypeCity(0.3),
		Land:    NewLessonTypeLand(0.6),
		Highway: NewLessonTypeHighway(0.1),
	}
}

type LessonTypeConfigurationParameter func(l *LessonType)

func NewLessonTypeCityParameter(percent float32) LessonTypeConfigurationParameter {
	return func(l *LessonType) {
		l.City = NewLessonTypeCity(percent)
	}
}

type Config struct {
	Date   string
	Lesson *LessonType
	Count  int
	Total  float32
}

func NewConfig() *Config {
	return &Config{
		Date:   NowDateString(),
		Lesson: NewLessonType(),
		Total:  0,
	}
}

func NowDateString() string {
	date := time.Now()
	return fmt.Sprintf("%d/%d/%d", date.Day(), date.Month(), date.Year())
}

func NewRandomConfig() *Config {
	config := &Config{
		Date:   NowDateString(),
		Lesson: NewLessonType(),
		Total:  0,
	}
	err := config.Lesson.ModLand(rand.Float32())
	if err != nil {
		return nil
	}
	err = config.Lesson.ModCity(rand.Float32())
	if err != nil {
		return nil
	}

	if ran := rand.Float32(); ran > 0.9 {
		err := config.Lesson.ModHighway(rand.Float32() * 43)
		if err != nil {
			return nil
		}
	} else {
		err := config.Lesson.ModHighway(rand.Float32())
		if err != nil {
			return nil
		}
	}
	return config
}

func NewRandomDays(count int) []*Config {
	days := make([]*Config, count)
	for i := 0; i < count; i++ {
		days[i] = NewRandomConfig()
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

func (c *Config) GetCount() int {
	return c.Count
}

func (c *Config) Validate() error {
	return nil
}

func (c *Config) Calculate(total float32) error {
	old := c.GetTotal()
	c.SetTotal(total * (c.GetLesson().GetLand().GetTotal() + c.GetLesson().GetCity().GetTotal() + c.GetLesson().GetHighway().GetTotal()))
	if err := c.Validate(); err != nil {
		c.SetTotal(old)
		return err
	}
	return nil
}

func (c *Config) String() string {
	output := fmt.Sprintf("Date: %s\t KM: %.2f", c.GetDate(), c.GetTotal())
	output += fmt.Sprintf("\tCity: %.0f%%", c.GetLesson().GetCity().GetTotal()*100)
	output += fmt.Sprintf("\tLand: %.0f%%", c.GetLesson().GetLand().GetTotal()*100)
	output += fmt.Sprintf("\tHighway: %.0f%%", c.GetLesson().GetHighway().GetTotal()*100)
	output += fmt.Sprintf("\tTotal: %.0f%%", c.GetLesson().GetTotal()*100)
	return output
}

func (c *Config) Print() {
	fmt.Print(c)
	fmt.Println()
}
