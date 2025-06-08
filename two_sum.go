This is a classic problem known as “Two Sum”.

Problem Statement (clarified):

Given an array of integers nums and an integer target, return the indices of the two numbers such that they add up to the target.

⸻

✅ Approach:

Use a hash map (in Go, it’s a map) to store each number and its index as you iterate.
For each number num, calculate the complement = target - num.
If the complement exists in the map, return the indices.

⸻

✅ Go Code:

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Map to store number and its index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// Check if complement exists in the map
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		// Store current number and index
		numMap[num] = i
	}

	// Return empty if no solution found
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Indices:", result) // Output: [0, 1]
}


⸻

💡 Example:

nums := []int{2, 7, 11, 15}
target := 9

	•	2 + 7 = 9 → so indices 0 and 1 are returned.

⸻

🧠 Time Complexity:
	•	O(n) time, where n is the length of nums
	•	O(n) space for the map

Let me know if you want the version without using a map (brute force, O(n²) approach) or want to write test cases.