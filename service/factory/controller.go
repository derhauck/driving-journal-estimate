package factory

import "driving-journal-estimate/service/internal/controller"

func NewRandomController() *controller.RandomController {
	return &controller.RandomController{}
}
