package config

import (
	"testing"
)

func TestNewParameter(t *testing.T) {
	param := NewParameter("key", 10)

	if param.GetKey() != "key" {
		t.Errorf("Expected key 'key', got '%s'", param.GetKey())
	}

	if param.GetValue() != 10 {
		t.Errorf("Expected value 10, got %v", param.GetValue())
	}
}
