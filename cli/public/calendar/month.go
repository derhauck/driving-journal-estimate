package calendar

import (
	"bufio"
	"bytes"
	"derhauck/driving-journal-estimate/public/day"
	"derhauck/driving-journal-estimate/public/logger"
	"fmt"
	"os"
)

type Day struct {
	day.Config
}

type Month struct {
	Days   []*day.Config `json:"days"`
	Total  float32       `json:"total"`
	Logger logger.Inf
}

func (m *Month) RandomDays(count uint) {
	m.Days = day.NewRandomDays(count)
}

func (m *Month) Calculate(total float32) {
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
}

func (m *Month) CalculateWithinRange(total float32, min float32, max float32) {
	m.Calculate(total)
	for _, d := range m.Days {
		if d.GetTotal() < min {
			// Todo implement min
		}
		if d.GetTotal() > max {
			// Todo implement max
		}
	}

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
		m.Logger.Logf("%s", d.String())
	}
	m.Logger.Logf("Total\tKM: %f", m.Total)
}

func (m *Month) WriteOut(fileName string) {
	path, err := os.Getwd()
	file, err := os.Create(path + "/" + fileName)
	if err != nil {
		m.Logger.Error(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(file)
	for _, d := range m.Days {
		output := fmt.Sprintf("%s\n", d.String())
		if _, err := w.Write(bytes.NewBufferString(output).Bytes()); err != nil {
			panic(err)
		}
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
}
