package calendar

import (
	"driving-journal-estimate/public/day"
	"driving-journal-estimate/public/logger"
	"fmt"
)

type Day struct {
	day.Config
}

type Month struct {
	Days   []*day.Config
	Total  float32
	Logger logger.Inf
}

func (m *Month) RandomDays(count int) {
	m.Days = day.NewRandomDays(count)
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

func (m *Month) CalculateWithinRange(total float32, min float32, max float32) error {
	err := m.Calculate(total)
	if err != nil {
		return err
	}
	for _, d := range m.Days {
		if d.GetTotal() < min {
			// Todo implement min
		}
		if d.GetTotal() > max {
			// Todo implement max
		}
	}

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

	for _, d := range m.Days {
		m.Logger.Infof("%s", d.String())
	}
	m.Logger.Infof("Total\tKM: %f", m.Total)
}
