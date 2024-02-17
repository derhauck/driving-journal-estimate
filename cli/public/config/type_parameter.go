package config

type LessonTypeParameter struct {
	Parameter[int]
	multiplier float32
}

func (l *LessonTypeParameter) GetMultiplier() float32 {
	return l.multiplier
}

func (l *LessonTypeParameter) Calculate(total int, seed float32) float32 {
	result := (float32(l.GetValue()) / 100) * float32(total) * (l.GetMultiplier() + seed)
	return result
}

func validLessonParameter(percent int) bool {
	if percent < 0 || percent > 100 {
		return false
	}
	return true
}

func (l *LessonTypeParameter) modify(percent int) {
	if validLessonParameter(percent) {
		l.value = percent
	}
}
