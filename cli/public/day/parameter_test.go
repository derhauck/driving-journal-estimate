package day

import "testing"

func TestNewLessonType(t *testing.T) {
	t.Run("Can create new LessonType", func(t *testing.T) {
		l := NewLessonType()

		if l.Multiplier == 0 {
			t.Error("Multiplier should not be 0")
		}
	})
	t.Run("Can modify LessonType", func(t *testing.T) {
		l := &LessonType{
			Multiplier: 0,
		}
		var city float64 = 0.44
		err := l.ModMultiplier(city)
		if err != nil {
			t.Error(err)
		}

		if l.Multiplier != city {
			t.Errorf("Multiplier should be '%f' but is '%f'", city, l.Multiplier)
		}

	})
}
