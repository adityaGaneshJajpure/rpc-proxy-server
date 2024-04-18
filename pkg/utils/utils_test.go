package utils

import "testing"

func TestStringContains(t *testing.T) {
	// Test cases
	tests := []struct {
		slice []string
		val   string
		want  bool
	}{
		{[]string{"test", "test2", "test3"}, "test2", true},
		{[]string{"test", "test2", "test3"}, "test4", false},
		{[]string{}, "test", false},
	}

	// Iterate over test cases
	for _, tc := range tests {
		// Call the function
		got := StringContains(tc.slice, tc.val)

		// Check the result
		if got != tc.want {
			t.Errorf("StringContains(%v, %q) = %t; want %t", tc.slice, tc.val, got, tc.want)
		}
	}
}
