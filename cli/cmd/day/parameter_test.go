package day

import "testing"

func TestNewLessonType(t *testing.T) {
	t.Run("Can create new LessonType", func(t *testing.T) {
		l := NewLessonType()
		if l.Highway == nil {
			t.Error("Highway should not be nil")
		}
		if l.Land == nil {
			t.Error("Land should not be nil")
		}
		if l.City == nil {
			t.Error("City should not be nil")
		}
	})
	t.Run("Can modify LessonType", func(t *testing.T) {
		l := &LessonType{
			City:    NewLessonTypeCity(0),
			Land:    NewLessonTypeLand(0),
			Highway: NewLessonTypeHighway(0),
		}
		var highway float32 = 0.33
		var city float32 = 0.44
		var land float32 = 0.75
		err := l.ModHighway(highway)
		if err != nil {
			t.Error(err)
		}
		err = l.ModCity(city)
		if err != nil {
			t.Error(err)
		}
		err = l.ModLand(land)
		if err != nil {
			t.Error(err)
		}

		if l.Highway.GetValue() != highway {
			t.Errorf("Highway should be '%f' but is '%f'", highway, l.Highway.GetValue())
		}
		if l.City.GetValue() != city {
			t.Errorf("City should be '%f' but is '%f'", city, l.City.GetValue())
		}
		if l.Land.GetValue() != land {
			t.Errorf("Land should be '%f' but is '%f'", land, l.Land.GetValue())
		}
	})
}
