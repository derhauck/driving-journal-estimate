package calendar

import (
	"driving-journal-estimate/cmd/day"
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
	avg := (total * 0.7) / float32(len(m.Days))
	var newTotal float32 = 0
	for _, d := range m.Days {
		err := d.Calculate(avg)
		if err != nil {
			return err
		}
		newTotal = newTotal + d.GetTotal()

	}

	totalDiff := total - newTotal
	if totalDiff > 0 {
		dailyDiff := totalDiff / float32(len(m.Days))
		for _, d := range m.Days {
			dailyTotal := d.GetTotal()
			d.SetTotal(dailyTotal + dailyDiff)
			newTotal = newTotal + dailyDiff
		}
	}
	m.Total = newTotal
	return nil
}

func (m *Month) Print() {
	for _, d := range m.Days {
		d.Print()
	}
	fmt.Println(fmt.Sprintf("Total\tKM: %f", m.Total))
}
