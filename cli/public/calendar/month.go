package calendar

import (
	"driving-journal-estimate/public/day"
	"fmt"
)

type Day struct {
	day.Config
}

type Month struct {
	Days  []*day.Config
	Total float32
}

func NewMonth(days []*day.Config) *Month {
	return &Month{
		Days: days,
	}
}

func NewRandomMonth(count int) *Month {
	return &Month{Days: day.NewRandomDays(count)}
}

func (m *Month) Calculate(total float32) error {
	var newTotal float32 = 0
	var totalDailyMultiplier float32 = 0
	for _, d := range m.Days {
		totalDailyMultiplier += d.GetLesson().GetTotal()
	}

	dailyDiff := total / totalDailyMultiplier
	for _, d := range m.Days {
		d.SetTotal(dailyDiff * d.GetLesson().GetTotal())
		newTotal += dailyDiff * d.GetLesson().GetTotal()
	}

	m.Total = newTotal
	return nil
}

func (m *Month) String() string {
	var output string
	for _, d := range m.Days {
		output += fmt.Sprintf("%s\n", d.String())
	}
	output += fmt.Sprintf("Total\tKM: %f\n", m.Total)
	return output
}

func (m *Month) Print() {
	fmt.Println(m)
}
