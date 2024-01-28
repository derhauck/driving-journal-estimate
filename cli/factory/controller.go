package factory

import "driving-journal-estimate/internal/controller"

func NewRandomController() *controller.RandomController {
	return &controller.RandomController{}
}
