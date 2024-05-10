package main

import (
	"fmt"
)

// LeftRotate performs left rotation on an array 'arr' by 'd' positions
func LeftRotate(arr []int, d int) {
	n := len(arr)
	d %= n
	reverse(arr, 0, d-1)
	reverse(arr, d, n-1)
	reverse(arr, 0, n-1)
}

// RightRotate performs right rotation on an array 'arr' by 'd' positions
func RightRotate(arr []int, d int) {
	n := len(arr)
	d %= n
	reverse(arr, 0, n-1)
	reverse(arr, 0, d-1)
	reverse(arr, d, n-1)
}

// reverse reverses the elements of array 'arr' from index 'start' to index 'end'
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", arr)

	// Left rotate array by 2 positions
	LeftRotate(arr, 1)
	fmt.Println("Array after left rotation:", arr)

	// Right rotate array by 2 positions
	RightRotate(arr, 1)
	fmt.Println("Array after right rotation:", arr)
}
