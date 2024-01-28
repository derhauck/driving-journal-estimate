package misc

type Parameter[T any] struct {
	key   string
	value T
}

func NewParameter[T any](key string, value T) Parameter[T] {
	return Parameter[T]{
		key:   key,
		value: value,
	}
}

func (p *Parameter[any]) GetKey() string {
	return p.key
}

func (p *Parameter[any]) GetValue() any {
	return p.value
}

func (p *Parameter[T]) SetValue(value T) {
	p.value = value
}
func (p *Parameter[T]) SetKey(key string) {
	p.key = key
}
