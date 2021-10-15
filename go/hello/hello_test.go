package hello

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"now",
			"hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
