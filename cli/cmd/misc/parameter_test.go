package misc

import "testing"

func TestParameter(t *testing.T) {
	t.Run("Can create integer parameter", func(t *testing.T) {
		key := "test"
		value := 10
		p := NewParameter[int](key, value)
		if p.GetKey() != key {
			t.Errorf("Expected parameter key to be '%s', got '%s'", key, p.GetKey())
		}
		if p.GetValue() != value {
			t.Errorf("Expected parameter value to be '%d', got '%d'", value, p.GetValue())
		}
	})
	t.Run("Can modify integer parameter", func(t *testing.T) {
		key := "test"
		value := 10
		p := Parameter[int]{}
		p.SetKey(key)
		p.SetValue(value)
		if p.GetKey() != key {
			t.Errorf("Expected parameter key to be '%s', got '%s'", key, p.GetKey())
		}
		if p.GetValue() != value {
			t.Errorf("Expected parameter value to be '%d', got '%d'", value, p.GetValue())
		}
	})
}
