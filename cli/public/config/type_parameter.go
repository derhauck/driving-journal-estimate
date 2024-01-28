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

func NewLessonTypeParameter(key string, value int, multiplier float32) *LessonTypeParameter {
	tmp := &LessonTypeParameter{}
	tmp.key = key
	tmp.value = value
	tmp.multiplier = multiplier
	return tmp
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

func NewLessonTypeCity(percent int) *LessonTypeParameter {
	if !validLessonParameter(percent) {
		return nil
	}
	return NewLessonTypeParameter("CITY", percent, 0.3)
}

func NewLessonTypeLand(percent int) *LessonTypeParameter {
	if !validLessonParameter(percent) {
		return nil
	}
	return NewLessonTypeParameter("LAND", percent, 0.6)
}

func NewLessonTypeHighway(percent int) *LessonTypeParameter {
	if !validLessonParameter(percent) {
		return nil
	}
	return NewLessonTypeParameter("HIGHWAY", percent, 1)
}

type LessonType struct {
	City    *LessonTypeParameter
	Land    *LessonTypeParameter
	Highway *LessonTypeParameter
}

func (l *LessonType) combined() int {
	return l.Land.value + l.Highway.value + l.City.value
}

func (l *LessonType) sanityCheck() bool {
	if l.combined() > 100 || l.combined() < 0 {
		return false
	}

	return true
}

func (l *LessonType) ModCity(percent int) {
	old := l.City.value
	l.City.modify(percent)
	if l.sanityCheck() == false {
		l.City.value = old
	}
}

func (l *LessonType) ModHighway(percent int) {
	old := l.Highway.value
	l.Highway.modify(percent)
	if l.sanityCheck() == false {
		l.City.value = old
	}
}

func (l *LessonType) ModLand(percent int) {
	old := l.Land.value
	l.Land.modify(percent)
	if l.sanityCheck() == false {
		l.Land.value = old
	}
}

func NewLessonType() *LessonType {
	return &LessonType{
		City:    NewLessonTypeCity(0),
		Land:    NewLessonTypeLand(100),
		Highway: NewLessonTypeHighway(0),
	}
}

type LessonTypeConfigurationParameter func(l *LessonType)

func NewLessonTypeCityParameter(percent int) LessonTypeConfigurationParameter {
	return func(l *LessonType) {
		l.City = NewLessonTypeCity(percent)
	}
}

type TypeParameter struct {
	Type *LessonType
}

func NewTypeParameter(lessonType *LessonType) *TypeParameter {
	return &TypeParameter{
		Type: lessonType,
	}
}
