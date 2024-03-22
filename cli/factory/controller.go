package factory

import "derhauck/driving-journal-estimate/internal/controller"

func NewRandomController() *controller.CalendarController {
	return &controller.CalendarController{
		Month:  NewMonth(),
		Logger: GetLogger(),
	}
}
