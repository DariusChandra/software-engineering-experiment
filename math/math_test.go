package math

import "testing"

func TestSum(t *testing.T) {
	// Test cases
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},      // Test case 1
		{10, 20, 30},   // Test case 2
		{0, 0, 0},      // Test case 3
		{-1, 1, 0},     // Test case 4
		{100, -50, 50}, // Test case 5
	}

	// Iterate through test cases
	for _, tc := range testCases {
		// Call sum function with test case inputs
		result := Sum(tc.a, tc.b)
		// Check if result matches expected
		if result != tc.expected {
			t.Errorf("sum(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
