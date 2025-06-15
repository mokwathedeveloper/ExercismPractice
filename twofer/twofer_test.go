package twofer

import "testing"

func TestShareWith(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty name",
			input:    "",
			expected: "One for you, one for me.",
		},
		{
			name:     "Alice",
			input:    "Alice",
			expected: "One for Alice, one for me.",
		},
		{
			name:     "Bob",
			input:    "Bob",
			expected: "One for Bob, one for me.",
		},
		{
			name:     "with spaces",
			input:    "John Smith",
			expected: "One for John Smith, one for me.",
		},
		{
			name:     "with symbols",
			input:    "Mary-Jane",
			expected: "One for Mary-Jane, one for me.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShareWith(tt.input); got != tt.expected {
				t.Errorf("ShareWith(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
