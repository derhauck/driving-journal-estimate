package main

import (
	"driving-journal-estimate/cmd/calendar"
	"fmt"
)

func main() {
	month := calendar.NewRandomMonth(30)
	err := month.Calculate(10000)
	month.Print()
	if err != nil {
		return
	}
}

func init() {
	fmt.Println("init")
}
