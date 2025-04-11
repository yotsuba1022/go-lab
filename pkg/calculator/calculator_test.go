package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -2, -3},
		{"zero", 0, 5, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Add(test.a, test.b); got != test.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, got, test.expected)
			}
		})
	}
}
