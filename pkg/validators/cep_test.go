package validators

import "testing"

func TestIsValidCep(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"8 digits", "12345678", true},
		{"less than 8 digits", "1234567", false},
		{"more than 8 digits", "123456789", false},
		{"with hyphen", "12345-678", false},
		{"letters", "abcdefgh", false},
		{"digits + letters", "1a23456b", false},
		{"empty", "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidCep(test.input)
			if got != test.expected {
				t.Errorf("IsValidCEP(%q) = %v; expected: %v", test.input, got, test.expected)
			}
		})
	}
}
