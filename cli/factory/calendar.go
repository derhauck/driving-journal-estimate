package factory

import (
	"driving-journal-estimate/public/calendar"
	"driving-journal-estimate/public/day"
)

func NewRandomMonth(count int) *calendar.Month {
	return &calendar.Month{
		Days:   day.NewRandomDays(count),
		Logger: GetLogger(),
	}
}
func NewMonth() *calendar.Month {
	return &calendar.Month{
		Days:   make([]*day.Config, 0),
		Logger: GetLogger(),
	}
}
