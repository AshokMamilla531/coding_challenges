package reverse_array

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
		{input: []int{1, 2, 3}, expected: []int{3, 2, 1}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
	}
	// Iterate through test cases
	for _, tc := range testCases {
		// Call the function
		actual := ReverseArray(tc.input)
		// Compare the result with the expected output
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("ReverseArray(%v) = %v; expected %v", tc.input, actual, tc.expected)
		}
	}
}
