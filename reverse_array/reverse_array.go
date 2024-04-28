package reverse_array

func ReverseArray(input []int) []int {
	// Initialize pointers
	left := 0
	right := len(input) - 1

	// Iterate until pointers meet
	for left < right {
		// Swap elements at left and right indices
		input[left], input[right] = input[right], input[left]
		// Move pointers inward
		left++
		right--
	}

	// Return the reversed array
	return input
}
