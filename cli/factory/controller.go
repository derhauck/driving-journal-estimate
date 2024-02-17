package factory

import "driving-journal-estimate/internal/controller"

func NewRandomController() *controller.CalendarController {
	return &controller.CalendarController{
		Month: NewMonth(),
	}
}
